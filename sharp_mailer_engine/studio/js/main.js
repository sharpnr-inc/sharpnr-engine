/* ══════════════════════════════════════════════════
   main.js — app state, init, global event bindings
══════════════════════════════════════════════════ */

const state = {
  templates: [], // [{id, name, content, raw}]
  activeId: null,
  activeTab: "preview",
  baseLayoutRaw: null, // raw base_layout.html source (from /api/base_layout)
  fields: {}, // Go struct map from /api/fields
  varData: {}, // {templateId: {FieldKey: value}}
};

let darkMode = window.matchMedia("(prefers-color-scheme: dark)").matches;

/* ══════════════════════════════════
   STATUS PILL
══════════════════════════════════ */
let statusTimer = null;

function setStatus(msg, type = "ok") {
  const pill = document.getElementById("statusPill");
  pill.textContent = msg;
  pill.className = `status-pill visible ${type}`;
  clearTimeout(statusTimer);
  statusTimer = setTimeout(() => {
    pill.className = "status-pill";
  }, 2000);
}

/* ══════════════════════════════════
   THEME
══════════════════════════════════ */
function applyTheme() {
  document.documentElement.setAttribute("data-theme", darkMode ? "dark" : "");
  const icon = document.getElementById("themeIcon");
  icon.innerHTML = darkMode
    ? `<circle cx="12" cy="12" r="5"/>
       <line x1="12" y1="1" x2="12" y2="3"/>
       <line x1="12" y1="21" x2="12" y2="23"/>
       <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/>
       <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/>
       <line x1="1" y1="12" x2="3" y2="12"/>
       <line x1="21" y1="12" x2="23" y2="12"/>
       <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/>
       <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/>`
    : `<path d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z"/>`;
}

/* ══════════════════════════════════
   SAVE (active template → disk)
══════════════════════════════════ */
async function saveActiveTemplate() {
  const tpl = state.templates.find((t) => t.id === state.activeId);
  if (!tpl) return;

  const btn = document.getElementById("saveBtn");
  btn.textContent = "saving…";
  btn.disabled = true;

  try {
    const goSource = buildGoTemplate(tpl.content);
    await apiSaveTemplate(tpl.name, goSource);
    // Update raw so the "go template" tab reflects the saved state
    tpl.raw = goSource;
    setStatus("saved!", "ok");
  } catch (err) {
    setStatus(`Save failed: ${err.message}`, "error");
  } finally {
    btn.textContent = "save";
    btn.disabled = false;
  }
}

/* ══════════════════════════════════
   COPY (copies current tab content)
══════════════════════════════════ */
function copyCurrentView() {
  const tpl = state.templates.find((t) => t.id === state.activeId);
  if (!tpl) return;

  const text =
    state.activeTab === "gotpl"
      ? (tpl.raw ?? buildGoTemplate(tpl.content))
      : buildPreviewHTML(tpl);

  navigator.clipboard
    .writeText(text)
    .then(() => {
      setStatus("copied!", "ok");
    })
    .catch(() => {
      setStatus("Copy failed", "error");
    });
}

/* ══════════════════════════════════
   RELOAD templates from disk
══════════════════════════════════ */
async function reloadTemplates() {
  const btn = document.getElementById("reloadBtn");
  btn.style.opacity = "0.4";
  btn.disabled = true;

  try {
    const templates = await loadTemplates();
    state.templates = templates;

    // Keep selection if it still exists, else pick first
    const still = state.templates.find((t) => t.id === state.activeId);
    state.activeId = still ? state.activeId : (state.templates[0]?.id ?? null);

    renderSidebar();
    if (state.activeId) selectTemplate(state.activeId);
    setStatus("reloaded", "ok");
  } catch (err) {
    setStatus(`Reload failed: ${err.message}`, "error");
  } finally {
    btn.style.opacity = "";
    btn.disabled = false;
  }
}

/* ══════════════════════════════════
   INIT
══════════════════════════════════ */
async function init() {
  applyTheme();

  // Resolve Flask URL before any fetch (handles Live Server vs Flask port)
  await resolveAPI();

  // Load in parallel — fields failure is non-fatal
  const [baseLayoutRaw, templates, fieldsResult] = await Promise.allSettled([
    loadBaseLayout(),
    loadTemplates(),
    loadFields(),
  ]);

  if (baseLayoutRaw.status === "fulfilled") {
    state.baseLayoutRaw = baseLayoutRaw.value;
  } else {
    console.warn("[studio] base_layout load failed:", baseLayoutRaw.reason);
  }

  if (templates.status === "fulfilled") {
    state.templates = templates.value;
  } else {
    throw templates.reason; // propagate to show API error in sidebar
  }

  if (fieldsResult.status === "fulfilled") {
    state.fields = fieldsResult.value;
  } else {
    console.warn("[studio] /api/fields failed:", fieldsResult.reason);
    state.fields = {};
  }

  renderSidebar();
  if (state.templates.length > 0) {
    selectTemplate(state.templates[0].id);
  }
}

/* ══════════════════════════════════
   GLOBAL EVENT BINDINGS
══════════════════════════════════ */
document.addEventListener("DOMContentLoaded", () => {
  // Preview tabs
  document.querySelectorAll(".ptab").forEach((btn) => {
    btn.addEventListener("click", () => switchTab(btn.dataset.tab));
  });

  // Topbar buttons
  document
    .getElementById("saveBtn")
    .addEventListener("click", saveActiveTemplate);
  document.getElementById("copyBtn").addEventListener("click", copyCurrentView);
  document
    .getElementById("reloadBtn")
    .addEventListener("click", reloadTemplates);
  document.getElementById("themeToggle").addEventListener("click", () => {
    darkMode = !darkMode;
    applyTheme();
  });

  // Keyboard shortcut: Cmd/Ctrl+S to save
  document.addEventListener("keydown", (e) => {
    if ((e.metaKey || e.ctrlKey) && e.key === "s") {
      e.preventDefault();
      saveActiveTemplate();
    }
  });

  // Module-level events
  bindSidebarEvents();
  bindVarsEvents();
  bindModalEvents();

  // Boot
  init().catch((err) => {
    console.error("[studio] init failed:", err);
    document.getElementById("tplList").innerHTML = `
      <div style="padding:16px;color:#dc2626;font-size:11px;font-family:monospace;line-height:1.6;">
        ⚠ Could not reach the API.<br/><br/>
        Make sure file_server.py is running:<br/>
        <code>cd studio &amp;&amp; python file_server.py</code>
      </div>`;
  });
});
