"""
Sharpnr Mail Studio — file_server.py
Flask dev server: serves the studio UI and exposes a small
REST API so mail_studio.js can read/write real files on disk.

Project layout assumed:
  sharpnr-engine/
  ├── email_templates/          ← templates live here
  └── sharp_mailer_engine/
      └── studio/
          ├── file_server.py    ← this file
          ├── mail_studio.html
          ├── mail_studio.css
          └── mail_studio.js
"""

import os
import json
from pathlib import Path
from flask import Flask, request, jsonify, send_from_directory
from flask_cors import CORS
import re

# ── PATHS ──────────────────────────────────────────────────────────────────────
STUDIO_DIR    = Path(__file__).parent.resolve()
TEMPLATES_DIR = (STUDIO_DIR / "../../email_templates").resolve()
TEMPLATE_DATA_GO = (STUDIO_DIR / "../templates_data.go").resolve()

app = Flask(__name__, static_folder=str(STUDIO_DIR), static_url_path="")
CORS(app)   # allow the browser to call the API from the same origin / live-reload


# ── HELPERS ────────────────────────────────────────────────────────────────────
def safe_path(filename: str) -> Path:
    """Resolve a filename inside TEMPLATES_DIR; reject path traversal."""
    target = (TEMPLATES_DIR / filename).resolve()
    if not str(target).startswith(str(TEMPLATES_DIR)):
        raise ValueError("Path traversal attempt blocked.")
    return target


# ══════════════════════════════════════════════════════════════════════════════
#  STATIC — serve the studio UI
# ══════════════════════════════════════════════════════════════════════════════
@app.route("/")
def index():
    return send_from_directory(STUDIO_DIR, "mail_studio.html")


# ══════════════════════════════════════════════════════════════════════════════
#  API — /api/templates
# ══════════════════════════════════════════════════════════════════════════════

@app.route("/api/templates", methods=["GET"])
def list_templates():
    """
    GET /api/templates
    Returns a JSON array of every *.html file in email_templates/.

    Response:
      [
        { "name": "base_layout.html", "builtin": true  },
        { "name": "otp.html",         "builtin": false },
        ...
      ]

    "builtin" = true only for base_layout.html (the wrapper, not editable
    as a standalone template in the studio).
    """
    TEMPLATES_DIR.mkdir(parents=True, exist_ok=True)
    files = sorted(TEMPLATES_DIR.glob("*.html"))
    result = []
    for f in files:
        result.append({
            "name":    f.name,
            "builtin": f.name == "base_layout.html",
        })
    return jsonify(result)


@app.route("/api/templates/<filename>", methods=["GET"])
def get_template(filename: str):
    """
    GET /api/templates/<filename>
    Returns the raw HTML content of one template file.

    Response:
      { "name": "otp.html", "content": "<!DOCTYPE html>..." }
    """
    try:
        path = safe_path(filename)
    except ValueError as e:
        return jsonify({"error": str(e)}), 400

    if not path.exists():
        return jsonify({"error": f"{filename} not found"}), 404

    return jsonify({
        "name":    filename,
        "content": path.read_text(encoding="utf-8"),
    })


@app.route("/api/templates/<filename>", methods=["PUT"])
def save_template(filename: str):
    """
    PUT /api/templates/<filename>
    Creates or overwrites a template file.

    Body (JSON):
      { "content": "<html>...</html>" }

    Response:
      { "ok": true, "name": "otp.html" }
    """
    if not filename.endswith(".html"):
        return jsonify({"error": "Only .html files are allowed."}), 400

    try:
        path = safe_path(filename)
    except ValueError as e:
        return jsonify({"error": str(e)}), 400

    body = request.get_json(force=True, silent=True) or {}
    content = body.get("content")
    if content is None:
        return jsonify({"error": "Missing 'content' field."}), 400

    TEMPLATES_DIR.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")
    return jsonify({"ok": True, "name": filename})


@app.route("/api/templates/<filename>", methods=["DELETE"])
def delete_template(filename: str):
    """
    DELETE /api/templates/<filename>
    Deletes a template file. Refuses to delete base_layout.html.

    Response:
      { "ok": true, "name": "otp.html" }
    """
    if filename == "base_layout.html":
        return jsonify({"error": "base_layout.html cannot be deleted."}), 403

    try:
        path = safe_path(filename)
    except ValueError as e:
        return jsonify({"error": str(e)}), 400

    if not path.exists():
        return jsonify({"error": f"{filename} not found"}), 404

    path.unlink()
    return jsonify({"ok": True, "name": filename})


# ══════════════════════════════════════════════════════════════════════════════
#  API — /api/base_layout
# ══════════════════════════════════════════════════════════════════════════════

@app.route("/api/base_layout", methods=["GET"])
def get_base_layout():
    """
    GET /api/base_layout
    Convenience endpoint — returns base_layout.html content so the JS
    can use it as the email wrapper instead of the hard-coded string.

    Response:
      { "content": "<!DOCTYPE html>..." }
    """
    path = safe_path("base_layout.html")
    if not path.exists():
        return jsonify({"error": "base_layout.html not found in email_templates/"}), 404

    return jsonify({"content": path.read_text(encoding="utf-8")})


@app.route("/api/fields", methods=["GET"])
def get_fields():
    """
    Parses template_data.go and returns all structs + their fields.

    Response:
    {
      "OTPData": [
        { "key": "Name",          "optional": false },
        { "key": "Code",          "optional": false },
        { "key": "ExpiryMinutes", "optional": false }
      ],
      "WelcomeData": [ ... ]
    }
    """
    if not TEMPLATE_DATA_GO.exists():
        return jsonify({"error": "template_data.go not found"}), 404

    source = TEMPLATE_DATA_GO.read_text(encoding="utf-8")

    # Match each struct block
    structs = {}
    for struct_match in re.finditer(r'type (\w+) struct \{([^}]+)\}', source):
        struct_name = struct_match.group(1)
        body        = struct_match.group(2)
        fields      = []

        for line in body.splitlines():
            line = line.strip()
            if not line or line.startswith("//"):
                continue
            # e.g. "Name   string" or "OrganizationName string // optional"
            field_match = re.match(r'^(\w+)\s+\w+', line)
            if field_match:
                key      = field_match.group(1)
                optional = "// optional" in line
                fields.append({ "key": key, "optional": optional })

        if fields:
            structs[struct_name] = fields

    return jsonify(structs)


# ══════════════════════════════════════════════════════════════════════════════
#  ENTRY POINT
# ══════════════════════════════════════════════════════════════════════════════
if __name__ == "__main__":
    print(f"[studio] Templates dir : {TEMPLATES_DIR}")
    print(f"[studio] Studio dir    : {STUDIO_DIR}")
    print(f"[studio] Serving at    : http://127.0.0.1:5000")
    app.run(debug=True, port=5000)