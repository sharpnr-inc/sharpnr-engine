/* ══════════════════════════════════════════════════
   sidebar.js — template list + selection + delete
══════════════════════════════════════════════════ */

function renderSidebar() {
  const list = document.getElementById("tplList");

  // Hide loading spinner
  const spinner = document.getElementById("sidebarLoading");
  if (spinner) spinner.remove();

  if (state.templates.length === 0) {
    list.innerHTML = `
      <div class="empty-state" style="padding:20px 14px;">
        No templates yet.<br/>Click "+ new template" to create one.
      </div>`;
    return;
  }

  list.innerHTML = state.templates
    .map(
      (t) => `
    <div class="tpl-item ${t.id === state.activeId ? "active" : ""}" data-id="${t.id}">
      <div class="tpl-dot"></div>
      <span class="tpl-name">${t.name}</span>
      ${t.dirty ? `<span class="tpl-dirty" title="Unsaved changes">●</span>` : ""}
      <button class="tpl-del" data-del="${t.name}" title="Delete ${t.name}">×</button>
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
  renderVarsPanel(tpl);
  loadEditor(tpl);
  refreshPreview();
}

/* ── FIX: delete was rendered but never handled ──
   Uses event delegation on #tplList so dynamically
   rendered items are always covered.
── */
function bindSidebarEvents() {
  const list = document.getElementById("tplList");

  list.addEventListener("click", async (e) => {
    // Delete button
    const delBtn = e.target.closest("[data-del]");
    if (delBtn) {
      e.stopPropagation();
      const filename = delBtn.dataset.del;
      const id = filename.replace(".html", "");
      if (!confirm(`Delete "${filename}"? This cannot be undone.`)) return;

      setStatus("deleting…", "warn");
      try {
        await apiDeleteTemplate(filename);
        state.templates = state.templates.filter((t) => t.id !== id);
        if (state.activeId === id) {
          state.activeId = state.templates[0]?.id ?? null;
        }
        renderSidebar();
        if (state.activeId) {
          selectTemplate(state.activeId);
        } else {
          document.getElementById("varsForm").innerHTML =
            `<div class="empty-state">Select a template to preview it.</div>`;
          document.getElementById("previewFrame").srcdoc = "";
          document.getElementById("editorFrame").srcdoc = "";
          loadEditor(null);
        }
        setStatus("deleted", "ok");
      } catch (err) {
        setStatus(`Delete failed: ${err.message}`, "error");
      }
      return;
    }

    // Template row click → select
    const item = e.target.closest(".tpl-item");
    if (item) selectTemplate(item.dataset.id);
  });
}
