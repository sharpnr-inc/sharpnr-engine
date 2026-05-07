const state = {
  templates: [],
  activeId: null,
  activeTab: "preview",
  baseLayoutRaw: null,
  fields: {},
  varData: {}, // ← add this
};

async function init() {
  const [baseLayoutRaw, templates, fields] = await Promise.all([
    loadBaseLayout(),
    loadTemplates(),
    loadFields(),
  ]);

  state.baseLayoutRaw = baseLayoutRaw;
  state.templates = templates;
  state.fields = fields;

  renderSidebar();
  if (state.templates.length > 0) selectTemplate(state.templates[0].id);
}

document.addEventListener("DOMContentLoaded", () => {
  document.getElementById("tplList").addEventListener("click", (e) => {
    const item = e.target.closest(".tpl-item");
    if (item) selectTemplate(item.dataset.id);
  });

  document.querySelectorAll(".ptab").forEach((btn) => {
    btn.addEventListener("click", () => switchTab(btn.dataset.tab));
  });

  bindModalEvents(); // ← must be after DOMContentLoaded, elements exist now
  bindVarsEvents(); // ← add this

  init().catch((err) => {
    console.error("[studio] init failed:", err);
    document.getElementById("tplList").innerHTML = `
      <div style="padding:16px;color:#dc2626;font-size:11px;font-family:monospace;">
        ⚠ Could not reach API.<br/>Is file_server.py running?
      </div>`;
  });
});
