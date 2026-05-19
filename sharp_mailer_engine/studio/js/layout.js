/* ══════════════════════════════════════════════════
   layout.js — template rendering helpers
══════════════════════════════════════════════════ */

/* ── Extract {{define "content"}}…{{end}} block from a raw Go template file ── */
function extractContent(raw) {
  const m = raw.match(/\{\{define\s+"content"\}\}([\s\S]*?)\{\{end\}\}/);
  return m ? m[1].trim() : raw.trim();
}

/* ── Client-side Go template renderer (preview only, not for production) ──
   Handles:
     {{.Key}}
     {{if .Key}}…{{end}}
── */
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

/* ── Build the full preview HTML by injecting rendered content into base_layout ──
   FIX: was referencing undefined global `BASE_LAYOUT_RAW`
        now correctly reads from state.baseLayoutRaw
── */
function buildPreviewHTML(tpl) {
  const raw = state.baseLayoutRaw;
  const data =
    state.varData && state.varData[tpl.id] ? state.varData[tpl.id] : {};
  const content = renderGoTemplate(tpl.content, data);

  if (!raw) {
    // Fallback if base layout hasn't loaded yet
    return `<!DOCTYPE html><html><head><meta charset="UTF-8"/></head><body style="font-family:sans-serif;padding:40px;">${content}</body></html>`;
  }

  return (
    raw
      // Strip Go define/end wrappers (they're not valid HTML)
      .replace(/\{\{define\s+"[^"]+"\}\}\s*/g, "")
      .replace(/\s*\{\{end\}\}\s*$/, "")
      // Replace the content slot — handles both {{block}} and {{template}} forms
      .replace(/\{\{block\s+"content"\s+\.\}\}\s*\{\{end\}\}/g, content)
      .replace(/\{\{template\s+"content"\s+\.\}\}/g, content)
      .trim()
  );
}

/* ── Build the Go template file source that gets saved to disk ── */
function buildGoTemplate(contentBlock) {
  return `{{template "base_layout.html" .}}\n\n{{define "content"}}\n${contentBlock}\n{{end}}\n`;
}
