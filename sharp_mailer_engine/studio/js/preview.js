/* ══════════════════════════════════════════════════
   preview.js — iframe preview, edit split, source view
══════════════════════════════════════════════════ */

function refreshPreview() {
  const tpl = state.templates.find((t) => t.id === state.activeId);
  if (!tpl) return;

  if (state.activeTab === "preview") {
    const frame = document.getElementById("previewFrame");
    frame.srcdoc = buildPreviewHTML(tpl);
    frame.onload = () => {
      try {
        const h = frame.contentDocument.body.scrollHeight;
        frame.style.height = Math.max(h + 40, 400) + "px";
      } catch (_) {
        /* cross-origin */
      }
    };
  } else if (state.activeTab === "edit") {
    // Live preview beside the editor — fills its pane, scrolls internally
    document.getElementById("editorFrame").srcdoc = buildPreviewHTML(tpl);
  } else if (state.activeTab === "source") {
    document.getElementById("sourceView").textContent = buildPreviewHTML(tpl);
  }
}

function switchTab(tab) {
  state.activeTab = tab;

  document.querySelectorAll(".ptab").forEach((btn) => {
    btn.classList.toggle("active", btn.dataset.tab === tab);
  });

  document
    .getElementById("previewWrap")
    .classList.toggle("hidden", tab !== "preview");
  document.getElementById("editWrap").classList.toggle("hidden", tab !== "edit");
  document
    .getElementById("sourceWrap")
    .classList.toggle("hidden", tab !== "source");

  // Viewport toggle only makes sense on the preview tab
  document
    .getElementById("viewportToggle")
    .classList.toggle("hidden", tab !== "preview");

  if (tab === "edit") {
    const tpl = state.templates.find((t) => t.id === state.activeId);
    loadEditor(tpl);
  }

  refreshPreview();
}

/* ── Desktop / mobile preview width ── */
function setViewport(vp) {
  state.viewport = vp;
  document.querySelectorAll(".vp-btn").forEach((btn) => {
    btn.classList.toggle("active", btn.dataset.vp === vp);
  });
  document
    .getElementById("previewWrap")
    .classList.toggle("mobile", vp === "mobile");
}
