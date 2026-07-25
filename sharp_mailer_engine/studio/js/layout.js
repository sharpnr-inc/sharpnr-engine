/* ══════════════════════════════════════════════════
   layout.js — template rendering helpers
══════════════════════════════════════════════════ */

/* ── Extract the {{define "content"}}…{{end}} block from a raw Go template file ──
   The closing {{end}} must be matched by depth, not by the first occurrence:
   content blocks routinely contain their own {{if}}…{{end}} pairs, and a
   non-greedy match would truncate everything after the first inner {{end}}
   (which then gets written back to disk on save — data loss). ── */
function extractContent(raw) {
  const def = /\{\{define\s+"content"\}\}/.exec(raw);
  if (!def) return raw.trim();

  const start = def.index + def[0].length;
  const tokenRe = /\{\{-?\s*(if|range|with|block|define|end)\b[^{}]*\}\}/g;
  tokenRe.lastIndex = start;

  let depth = 0;
  let tok;
  while ((tok = tokenRe.exec(raw)) !== null) {
    if (tok[1] === "end") {
      if (depth === 0) return raw.slice(start, tok.index).trim();
      depth--;
    } else {
      depth++;
    }
  }
  return raw.slice(start).trim();
}

/* ── Client-side Go template renderer (preview only, not for production) ──
   Handles:
     {{.Key}}
     {{if .Key}}…{{end}}
     {{if .Key}}…{{else}}…{{end}}
── */
function renderGoTemplate(content, data) {
  let out = content;

  // {{if .Key}}...{{else}}...{{end}}  (else block optional)
  out = out.replace(
    /\{\{if \.(\w+)\}\}([\s\S]*?)(?:\{\{else\}\}([\s\S]*?))?\{\{end\}\}/g,
    (_, key, ifBlock, elseBlock = "") => {
      const val = data[key];
      return val && String(val).trim() ? ifBlock : elseBlock;
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
