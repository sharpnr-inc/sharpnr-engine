function renderSidebar() {
  const list = document.getElementById("tplList");

  if (state.templates.length === 0) {
    list.innerHTML = `
      <div style="padding:20px 14px;text-align:center;
                  color:var(--text-tertiary);font-size:11px;
                  font-family:var(--font-mono);line-height:1.8;">
        No templates yet.
      </div>`;
    return;
  }

  list.innerHTML = state.templates
    .map(
      (t) => `
    <div class="tpl-item ${t.id === state.activeId ? "active" : ""}" data-id="${t.id}">
      <div class="tpl-dot"></div>
      <span class="tpl-name">${t.name}</span>
      <button class="tpl-del" data-del="${t.id}" title="Delete">×</button>
    </div>
  `,
    )
    .join("");
}

function selectTemplate(id) {
  state.activeId = id;
  const tpl = state.templates.find((t) => t.id === id);
  if (!tpl) return;
  renderSidebar();
  renderVarsPanel(tpl); // ← add this
  refreshPreview();
}
