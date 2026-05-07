/* ═══════════════════════════════════════════════════
   SHARPNR MAIL STUDIO — mail_studio.js  (phase 1: load + preview)
═══════════════════════════════════════════════════ */

const API = "http://127.0.0.1:5000/api";

/* ── API ── */
async function apiFetch(path) {
  const res = await fetch(`${API}${path}`);
  if (!res.ok) throw new Error(`API error: ${res.status} ${path}`);
  return res.json();
}

/* ── STATE ── */
let templates = []; // [{ name, content }]
let activeId = null; // active template name (without .html)
let BASE_LAYOUT_RAW = null;

/* ══════════════════════════════════
   LOAD
══════════════════════════════════ */
async function init() {
  // Load base layout + template list in parallel
  const [layoutData, list] = await Promise.all([
    apiFetch("/base_layout"),
    apiFetch("/templates"),
  ]);

  BASE_LAYOUT_RAW = layoutData.content;

  // Fetch content of each template (skip base_layout.html)
  const fetches = list
    .filter((t) => !t.builtin)
    .map((t) => apiFetch(`/templates/${t.name}`));

  const files = await Promise.all(fetches);

  templates = files.map((f) => ({
    id: f.name.replace(".html", ""),
    name: f.name,
    content: extractContent(f.content), // just the content block
    raw: f.content, // full go template on disk
  }));

  renderSidebar();

  if (templates.length > 0) selectTemplate(templates[0].id);
}

/* ── Pull {{define "content"}}...{{end}} out of a saved file ── */
function extractContent(raw) {
  const m = raw.match(/\{\{define\s+"content"\}\}([\s\S]*?)\{\{end\}\}/);
  return m ? m[1].trim() : raw.trim();
}

/* ══════════════════════════════════
   BASE LAYOUT — strip Go tags, inject content slot
══════════════════════════════════ */
function buildPreviewHTML(contentSlot) {
  if (!BASE_LAYOUT_RAW)
    return `<!DOCTYPE html><html><body>${contentSlot}</body></html>`;

  return BASE_LAYOUT_RAW.replace(/\{\{define\s+"[^"]+"\}\}\s*/g, "") // strip {{define "base_layout.html"}}
    .replace(/\s*\{\{end\}\}\s*$/g, "") // strip trailing {{end}}
    .replace(/\{\{block\s+"content"\s+\.\}\}\{\{end\}\}/g, contentSlot) // inject content
    .trim();
}

/* ══════════════════════════════════
   SIDEBAR
══════════════════════════════════ */
function renderSidebar() {
  const list = document.getElementById("tplList");

  if (templates.length === 0) {
    list.innerHTML = `
      <div style="padding:20px 14px;text-align:center;
                  color:var(--text-tertiary);font-size:11px;
                  font-family:var(--font-mono);line-height:1.8;">
        No templates yet.
      </div>`;
    return;
  }

  list.innerHTML = templates
    .map(
      (t) => `
    <div class="tpl-item ${t.id === activeId ? "active" : ""}" data-id="${t.id}">
      <div class="tpl-dot"></div>
      <span class="tpl-name">${t.name}</span>
    </div>
  `,
    )
    .join("");
}

/* ══════════════════════════════════
   SELECT + PREVIEW
══════════════════════════════════ */
function selectTemplate(id) {
  activeId = id;
  const tpl = templates.find((t) => t.id === id);
  if (!tpl) return;

  renderSidebar(); // re-render to update active state

  // Show in iframe
  const frame = document.getElementById("previewFrame");
  frame.srcdoc = buildPreviewHTML(tpl.content);
  frame.onload = () => {
    try {
      const h = frame.contentDocument.body.scrollHeight;
      frame.style.height = Math.max(h + 40, 400) + "px";
    } catch (e) {}
  };
}

/* ══════════════════════════════════
   EVENTS
══════════════════════════════════ */
document.addEventListener("DOMContentLoaded", () => {
  // Sidebar click
  document.getElementById("tplList").addEventListener("click", (e) => {
    const item = e.target.closest(".tpl-item");
    if (item) selectTemplate(item.dataset.id);
  });

  init().catch((err) => {
    console.error("[studio] init failed:", err);
    document.getElementById("tplList").innerHTML = `
      <div style="padding:16px;color:#dc2626;font-size:11px;font-family:monospace;">
        ⚠ Could not reach API.<br/>Is file_server.py running?
      </div>`;
  });
});
