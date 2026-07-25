"""
Sharpnr Mail Studio — file_server.py

Directory layout:
  sharpnr-engine/
  ├── email_templates/         ← TEMPLATES_DIR  (editable .html files)
  │   ├── assets/shicon.svg
  │   ├── base_layout.html
  │   ├── otp.html
  │   └── welcome_org.html
  └── sharp_mailer_engine/
      ├── templates_data.go    ← Go struct definitions
      └── studio/              ← STUDIO_DIR  (this file lives here)
          ├── file_server.py
          ├── mail_studio.html
          ├── mail_studio.css
          └── js/*.js

Run:
  cd sharp_mailer_engine/studio/
  pip install flask flask-cors
  python file_server.py
  → http://127.0.0.1:5000
"""

import re
from pathlib import Path
from flask import Flask, request, jsonify, send_from_directory
from flask_cors import CORS

# ── PATHS ──────────────────────────────────────────────────────────────
STUDIO_DIR       = Path(__file__).parent.resolve()
TEMPLATES_DIR    = (STUDIO_DIR / "../../email_templates").resolve()   # sharpnr-engine/email_templates/
TEMPLATE_DATA_GO = (STUDIO_DIR / "../templates_data.go").resolve()    # sharp_mailer_engine/templates_data.go
GO_SRC_DIR       = STUDIO_DIR.parent.resolve()                        # sharp_mailer_engine/ — scan all *.go here

EXCLUDED_FILES = {"base_layout.html"}

app = Flask(__name__, static_folder=str(STUDIO_DIR), static_url_path="")
CORS(app)


# ── HELPERS ────────────────────────────────────────────────────────────

def safe_path(filename: str) -> Path:
    """Resolve filename inside TEMPLATES_DIR; block path traversal."""
    target = (TEMPLATES_DIR / filename).resolve()
    if not str(target).startswith(str(TEMPLATES_DIR)):
        raise ValueError("Path traversal blocked.")
    return target


def scan_go_structs() -> dict:
    """
    Scan Go source files for struct definitions.
    Strategy:
      1. Try templates_data.go first (the dedicated template data file).
      2. Fall back to scanning all *.go files in sharp_mailer_engine/
         if templates_data.go doesn't exist or has no structs.
    Returns:
      {
        "OTPData": [
          {"key": "Name",  "optional": false},
          {"key": "Code",  "optional": false},
          ...
        ],
        ...
      }
    """
    structs = {}
    # Prefer the dedicated file; fall back to full scan
    go_files = (
        [TEMPLATE_DATA_GO] if TEMPLATE_DATA_GO.exists()
        else list(GO_SRC_DIR.glob("*.go"))
    )

    for go_file in go_files:
        try:
            source = go_file.read_text(encoding="utf-8")
        except OSError:
            continue

        for m in re.finditer(r'type\s+(\w+)\s+struct\s*\{([^}]+)\}', source):
            struct_name = m.group(1)
            body        = m.group(2)
            fields      = []

            for line in body.splitlines():
                line = line.strip()
                if not line or line.startswith("//"):
                    continue
                # Match:  FieldName   string   `json:"..."` // optional
                fm = re.match(r'^([A-Z]\w*)\s+\S+', line)
                if fm:
                    key      = fm.group(1)
                    optional = "optional" in line.lower()
                    fields.append({"key": key, "optional": optional})

            if fields:
                structs[struct_name] = fields

    return structs


# ── STATIC: studio UI ─────────────────────────────────────────────────

@app.route("/")
def index():
    return send_from_directory(STUDIO_DIR, "mail_studio.html")


@app.route("/js/<path:filename>")
def js_files(filename):
    js_dir = STUDIO_DIR / "js"
    return send_from_directory(js_dir, filename)


@app.route("/<path:filename>")
def static_files(filename):
    # Whitelist: only serve known studio assets
    allowed_exts = {".css", ".js", ".ico", ".svg", ".png"}
    p = Path(filename)
    if p.suffix in allowed_exts:
        return send_from_directory(STUDIO_DIR, filename)
    return "", 404


# ── API ────────────────────────────────────────────────────────────────

@app.route("/api/config", methods=["GET"])
def api_config():
    """
    GET /api/config
    Probed by api.js at startup: a 200 here means the page is being
    served by Flask, so relative /api URLs work.
    """
    return jsonify({
        "ok":            True,
        "templates_dir": str(TEMPLATES_DIR),
        "go_src_dir":    str(GO_SRC_DIR),
    })


@app.route("/api/templates", methods=["GET"])
def api_list_templates():
    """
    GET /api/templates
    Returns all *.html files in templates/ (excluding base_layout.html).
    [
      {"name": "otp.html",     "builtin": false},
      {"name": "billing.html", "builtin": false},
      ...
    ]
    """
    TEMPLATES_DIR.mkdir(parents=True, exist_ok=True)
    result = []
    for p in sorted(TEMPLATES_DIR.glob("*.html")):
        result.append({
            "name":    p.name,
            "builtin": p.name in EXCLUDED_FILES,
        })
    return jsonify(result)


@app.route("/api/templates/<filename>", methods=["GET"])
def api_get_template(filename: str):
    """
    GET /api/templates/<filename>
    Returns the raw source of one template file.
    {"name": "otp.html", "content": "{{template ...}}..."}
    """
    try:
        path = safe_path(filename)
    except ValueError as e:
        return jsonify({"error": str(e)}), 403

    if not path.exists():
        return jsonify({"error": f"{filename} not found"}), 404

    return jsonify({
        "name":    filename,
        "content": path.read_text(encoding="utf-8"),
    })


@app.route("/api/templates/<filename>", methods=["PUT"])
def api_save_template(filename: str):
    """
    PUT /api/templates/<filename>
    Body: {"content": "<full go template source>"}
    Writes content directly to disk.
    """
    if not filename.endswith(".html"):
        return jsonify({"error": "Only .html files allowed."}), 400
    if filename in EXCLUDED_FILES:
        return jsonify({"error": f"{filename} is protected."}), 403

    try:
        path = safe_path(filename)
    except ValueError as e:
        return jsonify({"error": str(e)}), 403

    body    = request.get_json(force=True, silent=True) or {}
    content = body.get("content")
    if not content:
        return jsonify({"error": "Missing 'content' field."}), 400

    TEMPLATES_DIR.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")
    return jsonify({"ok": True, "name": filename})


@app.route("/api/templates/<filename>", methods=["DELETE"])
def api_delete_template(filename: str):
    """DELETE /api/templates/<filename>"""
    if filename in EXCLUDED_FILES:
        return jsonify({"error": f"{filename} is protected."}), 403
    try:
        path = safe_path(filename)
    except ValueError as e:
        return jsonify({"error": str(e)}), 403

    if not path.exists():
        return jsonify({"error": f"{filename} not found"}), 404

    path.unlink()
    return jsonify({"ok": True, "name": filename})


@app.route("/api/base_layout", methods=["GET"])
def api_base_layout():
    """
    GET /api/base_layout
    Returns base_layout.html so the JS can use it as the email wrapper.
    {"content": "<!DOCTYPE html>..."}
    """
    path = TEMPLATES_DIR / "base_layout.html"
    if not path.exists():
        return jsonify({"error": "base_layout.html not found"}), 404
    return jsonify({"content": path.read_text(encoding="utf-8")})


@app.route("/api/fields", methods=["GET"])
def api_fields():
    """
    GET /api/fields
    Scans all *.go files in the mailer engine directory for struct
    definitions and returns their fields.
    {"OTPData": [{"key": "Name", "optional": false}, ...], ...}
    """
    structs = scan_go_structs()
    if not structs:
        return jsonify({"error": "No Go structs found. Check GO_SRC_DIR path."}), 404
    return jsonify(structs)


# ── ENTRY POINT ────────────────────────────────────────────────────────
if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument("--port", type=int, default=5000)
    parser.add_argument("--host", default="127.0.0.1")
    args = parser.parse_args()

    print(f"\n  Sharpnr Mail Studio")
    print(f"  Studio dir    : {STUDIO_DIR}")
    print(f"  Templates dir : {TEMPLATES_DIR}")
    print(f"  Go source dir : {GO_SRC_DIR}")
    print(f"  Open          → http://{args.host}:{args.port}\n")

    app.run(host=args.host, port=args.port, debug=True)