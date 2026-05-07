function extractContent(raw) {
  const m = raw.match(/\{\{define\s+"content"\}\}([\s\S]*?)\{\{end\}\}/);
  return m ? m[1].trim() : raw.trim();
}

function buildPreviewHTML(contentSlot) {
  if (!state.baseLayoutRaw)
    return `<!DOCTYPE html><html><body>${contentSlot}</body></html>`;

  return state.baseLayoutRaw
    .replace(/\{\{define\s+"[^"]+"\}\}\s*/g, "")
    .replace(/\s*\{\{end\}\}\s*$/g, "")
    .replace(/\{\{block\s+"content"\s+\.\}\}\{\{end\}\}/g, contentSlot)
    .trim();
}

function buildGoTemplate(content) {
  return `{{template "base_layout.html" .}}\n\n{{define "content"}}\n${content}\n{{end}}`;
}
