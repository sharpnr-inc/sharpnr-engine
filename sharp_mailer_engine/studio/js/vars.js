/* ══════════════════════════════════════════════════
   vars.js — variables panel for live preview
══════════════════════════════════════════════════ */

// Default sample values for known field keys
const FIELD_DEFAULTS = {
  Name: "John Doe",
  Email: "john@example.com",
  Title: "Hello from Sharpnr",
  Subtitle: "Here is a quick update",
  Message: "This is the body of your email.",
  Code: "847 291",
  ExpiryMinutes: 10,
  ActionURL: "https://sharpnr.com",
  ActionText: "Take Action",
  FooterNote: "If you didn't request this, ignore this email.",
  OrganizationName: "Acme Corp",
  Link: "https://sharpnr.com",
  LinkText: "Click here",
  Amount: "$49.00",
  Date: new Date().toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
  }),
  BillingPeriod: "May 1 – May 31, 2026",
  Reason: "Suspicious activity detected",
  Status: "Active",
  Plan: "Pro Plan",
};

function getTemplateFields(tpl) {
  // Scan content for {{.FieldName}} occurrences and return unique keys
  const matches = [...tpl.content.matchAll(/\{\{\.(\w+)\}\}/g)];
  const keys = [...new Set(matches.map((m) => m[1]))];
  return keys;
}

function renderVarsPanel(tpl) {
  const form = document.getElementById("varsForm");
  if (!form) return;

  const keys = getTemplateFields(tpl);

  if (keys.length === 0) {
    form.innerHTML = `
      <div style="padding:20px 0;text-align:center;
                  color:var(--text-tertiary);font-size:11px;
                  font-family:var(--font-mono);">
        No variables in this template.
      </div>`;
    return;
  }

  // Init defaults if not set
  if (!state.varData) state.varData = {};
  if (!state.varData[tpl.id]) {
    state.varData[tpl.id] = {};
    keys.forEach((key) => {
      state.varData[tpl.id][key] = FIELD_DEFAULTS[key] ?? `Sample ${key}`;
    });
  }

  const data = state.varData[tpl.id];

  form.innerHTML = keys
    .map((key) => {
      const val = data[key] ?? "";
      const isLong =
        key === "Message" || key === "FooterNote" || key === "Subtitle";
      const optional = [
        "Subtitle",
        "FooterNote",
        "ActionURL",
        "ActionText",
        "LinkText",
        "BillingPeriod",
        "OrganizationName",
      ].includes(key);

      if (isLong) {
        return `
        <div class="field-group">
          <label class="field-label">.${key}${optional ? ' <span class="opt">(optional)</span>' : ""}</label>
          <textarea class="field-textarea" data-key="${key}">${escapeAttr(String(val))}</textarea>
        </div>`;
      }

      return `
      <div class="field-group">
        <label class="field-label">.${key}${optional ? ' <span class="opt">(optional)</span>' : ""}</label>
        <input class="field-input" type="text" data-key="${key}" value="${escapeAttr(String(val))}" />
      </div>`;
    })
    .join("");
}

function escapeAttr(str) {
  return str
    .replace(/&/g, "&amp;")
    .replace(/"/g, "&quot;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;");
}

function bindVarsEvents() {
  document.getElementById("varsForm").addEventListener("input", (e) => {
    const el = e.target;
    if (!el.dataset.key) return;

    const tpl = state.templates.find((t) => t.id === state.activeId);
    if (!tpl) return;

    if (!state.varData[tpl.id]) state.varData[tpl.id] = {};
    state.varData[tpl.id][el.dataset.key] = el.value;

    refreshPreview();
  });
}
