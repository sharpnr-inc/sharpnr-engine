/* ══════════════════════════════════════════════════
   editor.js — Go template content editor (edit tab)

   Edits the {{define "content"}} block of the active
   template. The base_layout wrapper is re-applied by
   buildGoTemplate() on save, matching how templates
   are structured on disk.
══════════════════════════════════════════════════ */

let editorDebounce = null;

/* Load the active template's content block into the textarea. */
function loadEditor(tpl) {
  const ta = document.getElementById("editorArea");
  ta.value = tpl ? tpl.content : "";
}

function bindEditorEvents() {
  const ta = document.getElementById("editorArea");

  ta.addEventListener("input", () => {
    const tpl = state.templates.find((t) => t.id === state.activeId);
    if (!tpl) return;

    tpl.content = ta.value;
    tpl.dirty = tpl.content !== extractContent(tpl.raw);

    clearTimeout(editorDebounce);
    editorDebounce = setTimeout(() => {
      refreshPreview();
      renderVarsPanel(tpl); // template vars may have been added/removed
      renderSidebar(); // update the unsaved-changes dot
    }, 250);
  });

  // Tab key inserts two spaces instead of moving focus
  ta.addEventListener("keydown", (e) => {
    if (e.key === "Tab") {
      e.preventDefault();
      ta.setRangeText("  ", ta.selectionStart, ta.selectionEnd, "end");
      ta.dispatchEvent(new Event("input"));
    }
  });
}
