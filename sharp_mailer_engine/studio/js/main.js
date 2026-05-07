/* ══════════════════════════════════════════════════
   SHARPNR MAIL STUDIO — main.js
   Shared state + init + events
══════════════════════════════════════════════════ */

const state = {
  templates: [],
  activeId: null,
  activeTab: "preview",
  baseLayoutRaw: null,
};

/* ── INIT ── */
async function init() {
  const [baseLayoutRaw, templates] = await Promise.all([
    loadBaseLayout(),
    loadTemplates(),
  ]);

  state.baseLayoutRaw = baseLayoutRaw;
  state.templates = templates;

  renderSidebar();
  if (state.templates.length > 0) selectTemplate(state.templates[0].id);
}

/* ── EVENTS ── */
document.addEventListener("DOMContentLoaded", () => {
  // Sidebar clicks
  document.getElementById("tplList").addEventListener("click", (e) => {
    const delBtn = e.target.closest("[data-del]");
    if (delBtn) {
      deleteTemplate(delBtn.dataset.del);
      return;
    }
    const item = e.target.closest(".tpl-item");
    if (item) selectTemplate(item.dataset.id);
  });

  // Tab switches
  document.querySelectorAll(".ptab").forEach((btn) => {
    btn.addEventListener("click", () => switchTab(btn.dataset.tab));
  });

  init().catch((err) => {
    console.error("[studio] init failed:", err);
    document.getElementById("tplList").innerHTML = `
      <div style="padding:16px;color:#dc2626;font-size:11px;font-family:monospace;">
        ⚠ Could not reach API.<br/>Is file_server.py running?
      </div>`;
  });
});
