/* ══════════════════════════════════════════════════
   vars.js — variables panel
══════════════════════════════════════════════════ */

const FIELD_DEFAULTS = {
  Name: "John Doe",
  Email: "john@example.com",
  Title: "Hello from Sharpnr",
  Subtitle: "Here is a quick update",
  Message: "This is the body of your email.",
  Code: "847 291",
  ExpiryMinutes: "10",
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
  EmployeeName: "Jane Smith",
  EmployeeEmail: "jane@example.com",
  Role: "Engineer",
  Initials: "JS",
};

// Fields that should render as <textarea> instead of <input>
const TEXTAREA_FIELDS = new Set([
  "Message",
  "FooterNote",
  "Subtitle",
  "Body",
  "Content",
]);

// Fields that are optional (shown with "(optional)" label)
const OPTIONAL_FIELDS = new Set([
  "Subtitle",
  "FooterNote",
  "ActionURL",
  "ActionText",
  "LinkText",
  "BillingPeriod",
  "OrganizationName",
  "Link",
]);

/* ── FIX: detect fields from BOTH {{.Key}} and {{if .Key}} patterns ──
   Previously only matched {{.Key}}, so optional fields wrapped in
   {{if .Key}}…{{end}} never appeared in the vars panel.
── */
function getTemplateFields(tpl) {
  const content = tpl.content;
  const seen = new Set();
  const ordered = [];

  // Match {{.Key}} and {{if .Key}} in document order so the form
  // mirrors the visual top-to-bottom order of the template.
  const re = /\{\{(?:if\s+)?\.([A-Za-z][A-Za-z0-9_]*)\}\}/g;
  let m;
  while ((m = re.exec(content)) !== null) {
    const key = m[1];
    if (!seen.has(key)) {
      seen.add(key);
      ordered.push(key);
    }
  }

  return ordered;
}

function renderVarsPanel(tpl) {
  const form = document.getElementById("varsForm");
  if (!form) return;

  const keys = getTemplateFields(tpl);

  if (keys.length === 0) {
    form.innerHTML = `
      <div class="empty-state">
        No <code>{{.Variables}}</code> found in this template.
      </div>`;
    return;
  }

  // Initialise defaults for this template if first visit
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
      const val = String(data[key] ?? "");
      const optional = OPTIONAL_FIELDS.has(key);
      const optLabel = optional ? ` <span class="opt">(optional)</span>` : "";

      if (TEXTAREA_FIELDS.has(key)) {
        return `
      <div class="field-group">
        <label class="field-label">.${key}${optLabel}</label>
        <textarea class="field-textarea" data-key="${key}">${escapeAttr(val)}</textarea>
      </div>`;
      }

      return `
    <div class="field-group">
      <label class="field-label">.${key}${optLabel}</label>
      <input class="field-input" type="text" data-key="${key}" value="${escapeAttr(val)}" />
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
