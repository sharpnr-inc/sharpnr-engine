function extractContent(raw) {
  const m = raw.match(/\{\{define\s+"content"\}\}([\s\S]*?)\{\{end\}\}/);
  return m ? m[1].trim() : raw.trim();
}

function renderGoTemplate(content, data) {
  let out = content;

  // {{if .Key}}...{{end}}
  out = out.replace(
    /\{\{if \.(\w+)\}\}([\s\S]*?)\{\{end\}\}/g,
    (_, key, inner) => {
      const val = data[key];
      return val && String(val).trim() ? inner : "";
    },
  );

  // {{.Key}}
  out = out.replace(/\{\{\.(\w+)\}\}/g, (_, key) => {
    return data[key] !== undefined ? String(data[key]) : "";
  });

  return out;
}

function buildPreviewHTML(tpl) {
  if (!BASE_LAYOUT_RAW)
    return `<!DOCTYPE html><html><body>${tpl.content}</body></html>`;

  const data =
    state.varData && state.varData[tpl.id] ? state.varData[tpl.id] : {};
  const content = renderGoTemplate(tpl.content, data);

  return BASE_LAYOUT_RAW.replace(/\{\{define\s+"[^"]+"\}\}\s*/g, "")
    .replace(/\s*\{\{end\}\}\s*$/g, "")
    .replace(/\{\{block\s+"content"\s+\.\}\}\{\{end\}\}/g, content)
    .trim();
}

function buildGoTemplate(content) {
  return `{{template "base_layout.html" .}}\n\n{{define "content"}}\n${content}\n{{end}}`;
}
