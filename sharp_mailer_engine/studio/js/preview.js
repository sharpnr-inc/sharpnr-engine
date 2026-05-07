/* ══════════════════════════════════════════════════
   preview.js — iframe preview + source views
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
  } else if (state.activeTab === "source") {
    document.getElementById("sourceView").textContent = buildPreviewHTML(tpl);
  } else if (state.activeTab === "gotpl") {
    // Show the raw Go template source (what actually lives on disk)
    document.getElementById("gotplView").textContent =
      tpl.raw ?? buildGoTemplate(tpl.content);
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
  document
    .getElementById("sourceWrap")
    .classList.toggle("hidden", tab !== "source");
  document
    .getElementById("gotplWrap")
    .classList.toggle("hidden", tab !== "gotpl");

  refreshPreview();
}
