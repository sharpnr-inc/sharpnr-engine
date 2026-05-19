/* ══════════════════════════════════════════════════
   api.js — HTTP layer for Mail Studio

   Port resolution strategy
   ────────────────────────
   The Makefile runs two servers in parallel:
     • Flask  :5000  (or $FLASK_PORT)  → API
     • Live   :5500  (or $LIVE_PORT)   → static HTML

   When the page loads from Live Server (port ≠ Flask port),
   we need to hit Flask's absolute URL for API calls.

   1. Try GET /api/config from same origin first
      (works when page IS served by Flask).
   2. If that 404s, fall back to the well-known default port 5000.
══════════════════════════════════════════════════ */

const DEFAULT_FLASK_PORT = 5000;
let API = "/api"; // will be rewritten by resolveAPI() if needed

async function resolveAPI() {
  // If already on Flask, relative /api works fine
  try {
    const res = await fetch("/api/config", { method: "GET" });
    if (res.ok) {
      // Page is served by Flask — keep relative URL
      return;
    }
  } catch (_) {
    /* not on Flask */
  }

  // Page is on Live Server — point at Flask directly
  const port = 5000;
  API = `http://127.0.0.1:${port}/api`;
}

async function apiFetch(path, options = {}) {
  const res = await fetch(`${API}${path}`, {
    headers: { "Content-Type": "application/json" },
    ...options,
  });
  if (!res.ok) {
    let msg = `${res.status} ${res.statusText}`;
    try {
      const body = await res.json();
      if (body.error) msg = body.error;
    } catch (_) {
      /* ignore */
    }
    throw new Error(msg);
  }
  return res.json();
}

/* ── Public API ── */

async function loadFields() {
  return apiFetch("/fields");
}

async function loadBaseLayout() {
  const data = await apiFetch("/base_layout");
  return data.content;
}

async function loadTemplates() {
  const list = await apiFetch("/templates");
  const editable = list.filter((t) => !t.builtin);

  const results = await Promise.allSettled(
    editable.map((t) => apiFetch(`/templates/${t.name}`)),
  );

  return results
    .filter((r) => r.status === "fulfilled")
    .map((r) => {
      const f = r.value;
      return {
        id: f.name.replace(".html", ""),
        name: f.name,
        content: extractContent(f.content),
        raw: f.content,
      };
    });
}

async function apiSaveTemplate(name, goTemplateSource) {
  return apiFetch(`/templates/${name}`, {
    method: "PUT",
    body: JSON.stringify({ content: goTemplateSource }),
  });
}

async function apiDeleteTemplate(name) {
  return apiFetch(`/templates/${name}`, { method: "DELETE" });
}
