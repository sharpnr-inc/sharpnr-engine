/* ═══════════════════════════════════════════════════
   SHARPNR MAIL STUDIO — mail_studio.js
═══════════════════════════════════════════════════ */

/* ── BASE LAYOUT ── */
const BASE_LAYOUT = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Sharpnr</title>
</head>
<body style="margin:0;padding:0;background-color:#f8fafc;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',system-ui,sans-serif;">
  <table cellpadding="0" cellspacing="0" border="0" width="100%">
    <tr>
      <td align="center" style="padding:40px 20px">
        <table cellpadding="0" cellspacing="0" border="0" style="max-width:480px;width:100%;background:#ffffff;border-radius:16px;box-shadow:0 4px 20px rgba(0,0,0,0.08);">
          <tr>
            <td style="background:linear-gradient(135deg,#6366f1 0%,#8b5cf6 100%);padding:40px 30px;text-align:center;border-radius:16px 16px 0 0;">
              <h1 style="color:#ffffff;font-size:24px;font-weight:700;margin:0;letter-spacing:-0.5px;">Sharpnr</h1>
            </td>
          </tr>
          <tr>
            <td style="padding:50px 40px;">
              __CONTENT__
            </td>
          </tr>
          <tr>
            <td style="background:#f9fafb;padding:25px 40px;text-align:center;border-radius:0 0 16px 16px;border-top:1px solid #e5e7eb;">
              <p style="color:#9ca3af;font-size:12px;margin:0;">© 2026 Sharpnr. All rights reserved.</p>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>`;

/* ── BUILT-IN TEMPLATES ── */
const BUILTIN_TEMPLATES = [
  {
    id: "otp",
    name: "otp.html",
    builtin: true,
    fields: [
      {
        key: "Code",
        label: ".Code",
        type: "text",
        optional: false,
        default: "483920",
      },
      {
        key: "Expiry",
        label: ".Expiry",
        type: "text",
        optional: false,
        default: "10 minutes",
      },
    ],
    content: `<div style="text-align:center;margin-bottom:40px;">
  <h2 style="color:#1f2937;font-size:22px;font-weight:600;margin:0 0 12px;line-height:1.3;">
    Verify Your Organization
  </h2>
  <p style="color:#6b7280;font-size:16px;margin:0;line-height:1.5;">
    Enter this verification code to complete your registration:
  </p>
</div>

<div style="background:#f8fafc;border:2px solid #e2e8f0;border-radius:12px;padding:30px;text-align:center;margin:30px 0;">
  <div style="background:#ffffff;border:2px solid #6366f1;border-radius:8px;padding:16px;display:inline-block;">
    <span style="font-size:28px;font-weight:700;letter-spacing:6px;color:#6366f1;font-family:'Courier New',monospace;">
      {{.Code}}
    </span>
  </div>
  <p style="color:#94a3b8;font-size:12px;margin:15px 0 0;font-style:italic;">
    Expires in {{.Expiry}}
  </p>
</div>`,
  },
  {
    id: "notification",
    name: "notification.html",
    builtin: true,
    fields: [
      {
        key: "Title",
        label: ".Title",
        type: "text",
        optional: false,
        default: "Action Required",
      },
      {
        key: "Subtitle",
        label: ".Subtitle",
        type: "text",
        optional: true,
        default: "Your account needs attention",
      },
      {
        key: "Message",
        label: ".Message",
        type: "textarea",
        optional: false,
        default:
          "We noticed some activity on your account that requires your review. Please take a moment to verify your recent actions.",
      },
      {
        key: "ActionURL",
        label: ".ActionURL",
        type: "text",
        optional: true,
        default: "https://sharpnr.com/review",
      },
      {
        key: "ActionText",
        label: ".ActionText",
        type: "text",
        optional: true,
        default: "Review Activity",
      },
      {
        key: "FooterNote",
        label: ".FooterNote",
        type: "text",
        optional: true,
        default:
          "If you didn't request this, you can safely ignore this email.",
      },
    ],
    content: `<div style="text-align:center;margin-bottom:36px;">
  <h2 style="color:#1f2937;font-size:22px;font-weight:600;margin:0 0 10px;line-height:1.3;">
    {{.Title}}
  </h2>
  {{if .Subtitle}}
  <p style="color:#6b7280;font-size:16px;margin:0;line-height:1.5;">
    {{.Subtitle}}
  </p>
  {{end}}
</div>

<div style="background:#f8fafc;border:1px solid #e5e7eb;border-radius:12px;padding:28px;margin:28px 0;">
  <p style="color:#374151;font-size:15px;line-height:1.6;margin:0;">
    {{.Message}}
  </p>
</div>

{{if .ActionURL}}
<div style="text-align:center;margin-top:30px;">
  <a href="{{.ActionURL}}" style="display:inline-block;background:#6366f1;color:#ffffff;font-size:14px;font-weight:600;text-decoration:none;padding:12px 22px;border-radius:8px;">
    {{.ActionText}}
  </a>
</div>
{{end}}

{{if .FooterNote}}
<p style="margin-top:30px;color:#9ca3af;font-size:12px;text-align:center;line-height:1.4;">
  {{.FooterNote}}
</p>
{{end}}`,
  },
  {
    id: "billing",
    name: "billing.html",
    builtin: true,
    fields: [
      {
        key: "Name",
        label: ".Name",
        type: "text",
        optional: false,
        default: "John Doe",
      },
      {
        key: "Amount",
        label: ".Amount",
        type: "text",
        optional: false,
        default: "$49.00",
      },
      {
        key: "Date",
        label: ".Date",
        type: "text",
        optional: false,
        default: new Date().toLocaleDateString("en-US", {
          year: "numeric",
          month: "long",
          day: "numeric",
        }),
      },
      {
        key: "Plan",
        label: ".Plan",
        type: "text",
        optional: false,
        default: "Pro Plan",
      },
      {
        key: "Link",
        label: ".Link",
        type: "text",
        optional: true,
        default: "https://sharpnr.com/billing",
      },
    ],
    content: `<div style="text-align:center;margin-bottom:32px;">
  <h2 style="color:#1f2937;font-size:22px;font-weight:600;margin:0 0 8px;">
    Payment Confirmed
  </h2>
  <p style="color:#6b7280;font-size:15px;margin:0;">
    Hi {{.Name}}, your payment was successful.
  </p>
</div>

<div style="background:#f8fafc;border:1px solid #e5e7eb;border-radius:12px;padding:24px;margin:24px 0;">
  <table width="100%" cellpadding="0" cellspacing="0" border="0">
    <tr>
      <td style="color:#6b7280;font-size:13px;padding:8px 0;border-bottom:1px solid #e5e7eb;">Plan</td>
      <td style="color:#1f2937;font-size:13px;font-weight:600;text-align:right;padding:8px 0;border-bottom:1px solid #e5e7eb;">{{.Plan}}</td>
    </tr>
    <tr>
      <td style="color:#6b7280;font-size:13px;padding:8px 0;border-bottom:1px solid #e5e7eb;">Date</td>
      <td style="color:#1f2937;font-size:13px;font-weight:600;text-align:right;padding:8px 0;border-bottom:1px solid #e5e7eb;">{{.Date}}</td>
    </tr>
    <tr>
      <td style="color:#6b7280;font-size:14px;font-weight:600;padding:12px 0 0;">Total</td>
      <td style="color:#6366f1;font-size:20px;font-weight:700;text-align:right;padding:12px 0 0;">{{.Amount}}</td>
    </tr>
  </table>
</div>

{{if .Link}}
<div style="text-align:center;margin-top:28px;">
  <a href="{{.Link}}" style="display:inline-block;background:#6366f1;color:#ffffff;font-size:14px;font-weight:600;text-decoration:none;padding:12px 22px;border-radius:8px;">
    View Receipt
  </a>
</div>
{{end}}`,
  },
];

/* ── AVAILABLE FIELDS FOR NEW TEMPLATE WIZARD ── */
const ALL_FIELDS = [
  "Title",
  "Subtitle",
  "Message",
  "Code",
  "Expiry",
  "ActionURL",
  "ActionText",
  "FooterNote",
  "Name",
  "Email",
  "OrganizationName",
  "Link",
  "LinkText",
  "Amount",
  "Date",
  "Reason",
  "Status",
  "Plan",
];

const OPTIONAL_FIELDS = new Set([
  "Subtitle",
  "FooterNote",
  "ActionURL",
  "ActionText",
  "LinkText",
]);

const FIELD_DEFAULTS = {
  Code: "483920",
  Expiry: "10 minutes",
  Name: "John Doe",
  Email: "john@example.com",
  Title: "Hello from Sharpnr",
  Subtitle: "Here is a quick update for you",
  Message:
    "This is the body of your email. Replace this with your actual content.",
  ActionURL: "https://sharpnr.com",
  ActionText: "Take Action",
  FooterNote: "If you didn't request this, you can safely ignore this email.",
  OrganizationName: "Acme Corp",
  Link: "https://sharpnr.com",
  LinkText: "Click here",
  Amount: "$49.00",
  Date: new Date().toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
  }),
  Reason: "Suspicious activity detected",
  Status: "Active",
  Plan: "Pro Plan",
};

/* ── STATE ── */
let templates = BUILTIN_TEMPLATES.map((t) => deepClone(t));
let activeId = "otp";
let fieldData = {};
let activeTab = "preview";
let newChips = new Set(["Title", "Message"]);
let darkMode = window.matchMedia("(prefers-color-scheme: dark)").matches;

function deepClone(obj) {
  return JSON.parse(JSON.stringify(obj));
}

/* ══════════════════════════════════
   THEME
══════════════════════════════════ */
function applyTheme() {
  document.documentElement.setAttribute("data-theme", darkMode ? "dark" : "");
  const icon = document.getElementById("themeIcon");
  if (darkMode) {
    icon.innerHTML = `
      <circle cx="12" cy="12" r="5"/>
      <line x1="12" y1="1" x2="12" y2="3"/>
      <line x1="12" y1="21" x2="12" y2="23"/>
      <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"/>
      <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"/>
      <line x1="1" y1="12" x2="3" y2="12"/>
      <line x1="21" y1="12" x2="23" y2="12"/>
      <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"/>
      <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"/>`;
  } else {
    icon.innerHTML = `<path d="M21 12.79A9 9 0 1111.21 3 7 7 0 0021 12.79z"/>`;
  }
}

/* ══════════════════════════════════
   GO TEMPLATE RENDERER
══════════════════════════════════ */
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

function getRenderedHTML(tpl) {
  const data = getFieldData(tpl);
  const rendered = renderGoTemplate(tpl.content, data);
  return BASE_LAYOUT.replace("__CONTENT__", rendered);
}

function getGoTemplate(tpl) {
  return `{{define "${tpl.name}"}}
{{template "base_layout.html" .}}
{{end}}

{{define "content"}}
${tpl.content}
{{end}}`;
}

function getFieldData(tpl) {
  if (!fieldData[tpl.id]) {
    const d = {};
    tpl.fields.forEach((f) => {
      d[f.key] = f.default;
    });
    fieldData[tpl.id] = d;
  }
  return fieldData[tpl.id];
}

/* ══════════════════════════════════
   SIDEBAR
══════════════════════════════════ */
function renderSidebar() {
  const list = document.getElementById("tplList");
  list.innerHTML = templates
    .map(
      (t) => `
    <div class="tpl-item ${t.id === activeId ? "active" : ""}" data-id="${t.id}">
      <div class="tpl-dot"></div>
      <span class="tpl-name">${t.name}</span>
      ${
        !t.builtin
          ? `<button class="tpl-del" data-del="${t.id}" title="Delete">×</button>`
          : ""
      }
    </div>
  `,
    )
    .join("");
}

/* ══════════════════════════════════
   VARS FORM
══════════════════════════════════ */
function renderVarsForm(tpl) {
  const data = getFieldData(tpl);
  const form = document.getElementById("varsForm");

  form.innerHTML = tpl.fields
    .map((f) => {
      const val = data[f.key] ?? f.default;
      const optBadge = f.optional ? ` <span class="opt">(optional)</span>` : "";

      if (f.type === "textarea") {
        return `
        <div class="field-group">
          <label class="field-label">${f.label}${optBadge}</label>
          <textarea class="field-textarea" data-key="${f.key}">${escapeAttr(val)}</textarea>
        </div>`;
      }

      return `
      <div class="field-group">
        <label class="field-label">${f.label}${optBadge}</label>
        <input class="field-input" type="text" data-key="${f.key}" value="${escapeAttr(val)}" />
      </div>`;
    })
    .join("");
}

function escapeAttr(str) {
  return String(str)
    .replace(/&/g, "&amp;")
    .replace(/"/g, "&quot;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;");
}

/* ══════════════════════════════════
   PREVIEW / TABS
══════════════════════════════════ */
function refreshPreview() {
  const tpl = templates.find((t) => t.id === activeId);
  if (!tpl) return;

  if (activeTab === "preview") {
    const frame = document.getElementById("previewFrame");
    frame.srcdoc = getRenderedHTML(tpl);
    frame.onload = () => {
      try {
        const h = frame.contentDocument.body.scrollHeight;
        frame.style.height = Math.max(h + 40, 400) + "px";
      } catch (e) {
        /* cross-origin safety */
      }
    };
  } else if (activeTab === "code") {
    document.getElementById("codeView").textContent = getRenderedHTML(tpl);
  } else if (activeTab === "gotpl") {
    document.getElementById("gotplView").textContent = getGoTemplate(tpl);
  }
}

function switchTab(tab) {
  activeTab = tab;

  document.querySelectorAll(".ptab").forEach((btn) => {
    btn.classList.toggle("active", btn.dataset.tab === tab);
  });

  document
    .getElementById("previewWrap")
    .classList.toggle("hidden", tab !== "preview");
  document
    .getElementById("codeWrap")
    .classList.toggle("hidden", tab !== "code");
  document
    .getElementById("gotplWrap")
    .classList.toggle("hidden", tab !== "gotpl");

  refreshPreview();
}

/* ══════════════════════════════════
   SELECT TEMPLATE
══════════════════════════════════ */
function selectTemplate(id) {
  activeId = id;
  const tpl = templates.find((t) => t.id === id);
  if (!tpl) return;
  renderSidebar();
  renderVarsForm(tpl);
  refreshPreview();
}

/* ══════════════════════════════════
   SAVE TO TEMPLATES FOLDER
══════════════════════════════════ */
async function saveToTemplates() {
  const tpl = templates.find((t) => t.id === activeId);
  if (!tpl) return;

  // We offer both the Go template source (for Go projects) and the rendered HTML
  const isGoTab = activeTab === "gotpl";
  const content = isGoTab ? getGoTemplate(tpl) : getRenderedHTML(tpl);
  const suggestedName = tpl.name;

  // Use File System Access API if available (Chrome / Edge with live server)
  if (window.showSaveFilePicker) {
    try {
      const handle = await window.showSaveFilePicker({
        suggestedName,
        types: [
          {
            description: "HTML Template",
            accept: { "text/html": [".html"] },
          },
        ],
        startIn: "documents",
      });
      const writable = await handle.createWritable();
      await writable.write(content);
      await writable.close();
      showSaved("saved!");
    } catch (err) {
      if (err.name !== "AbortError") {
        console.error("Save error:", err);
        fallbackDownload(content, suggestedName);
      }
    }
  } else {
    // Fallback: trigger browser download
    fallbackDownload(content, suggestedName);
  }
}

function fallbackDownload(content, filename) {
  const blob = new Blob([content], { type: "text/html" });
  const a = document.createElement("a");
  a.href = URL.createObjectURL(blob);
  a.download = filename;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(a.href);
  showSaved("downloaded!");
}

/* ══════════════════════════════════
   COPY
══════════════════════════════════ */
function copyCurrentTab() {
  const tpl = templates.find((t) => t.id === activeId);
  if (!tpl) return;

  const text =
    activeTab === "gotpl" ? getGoTemplate(tpl) : getRenderedHTML(tpl);

  navigator.clipboard.writeText(text).then(() => {
    const btn = document.getElementById("copyBtn");
    const orig = btn.textContent;
    btn.textContent = "copied!";
    setTimeout(() => {
      btn.textContent = orig;
    }, 1500);
  });
}

/* ══════════════════════════════════
   SAVED PILL
══════════════════════════════════ */
let savedTimer = null;
function showSaved(msg = "saved") {
  const pill = document.getElementById("savedPill");
  pill.textContent = msg;
  pill.classList.add("visible");
  clearTimeout(savedTimer);
  savedTimer = setTimeout(() => pill.classList.remove("visible"), 1800);
}

/* ══════════════════════════════════
   DELETE TEMPLATE
══════════════════════════════════ */
function deleteTemplate(id) {
  if (!confirm(`Delete "${id}.html"?`)) return;
  templates = templates.filter((t) => t.id !== id);
  if (activeId === id) {
    activeId = templates[0]?.id ?? null;
  }
  renderSidebar();
  if (activeId) selectTemplate(activeId);
}

/* ══════════════════════════════════
   NEW TEMPLATE MODAL
══════════════════════════════════ */
function openModal() {
  newChips = new Set(["Title", "Message"]);
  renderChips();
  const nameInput = document.getElementById("newTplName");
  nameInput.value = "";
  nameInput.classList.remove("error");
  document.getElementById("nameError").classList.remove("visible");
  document.getElementById("modalOverlay").classList.add("open");
  setTimeout(() => nameInput.focus(), 150);
}

function closeModal() {
  document.getElementById("modalOverlay").classList.remove("open");
}

function renderChips() {
  const container = document.getElementById("fieldChips");
  container.innerHTML = ALL_FIELDS.map(
    (f) => `
    <div class="chip ${newChips.has(f) ? "on" : ""}" data-field="${f}">.${f}</div>
  `,
  ).join("");
}

function toggleChip(field) {
  newChips.has(field) ? newChips.delete(field) : newChips.add(field);
  renderChips();
}

function scaffoldContent(fields) {
  const parts = [];

  // Render in a sensible order
  const order = [
    "Title",
    "Subtitle",
    "Code",
    "Expiry",
    "Name",
    "Message",
    "Amount",
    "Date",
    "Plan",
    "Status",
    "Reason",
    "Email",
    "OrganizationName",
    "ActionURL",
    "FooterNote",
    "Link",
    "LinkText",
  ];

  const sorted = [...fields].sort((a, b) => {
    const ai = order.indexOf(a);
    const bi = order.indexOf(b);
    return (ai === -1 ? 999 : ai) - (bi === -1 ? 999 : bi);
  });

  for (const key of sorted) {
    switch (key) {
      case "Title":
        parts.push(
          `<h2 style="color:#1f2937;font-size:22px;font-weight:600;margin:0 0 12px;text-align:center;line-height:1.3;">{{.Title}}</h2>`,
        );
        break;
      case "Subtitle":
        parts.push(`{{if .Subtitle}}
<p style="color:#6b7280;font-size:16px;margin:0 0 24px;text-align:center;line-height:1.5;">{{.Subtitle}}</p>
{{end}}`);
        break;
      case "Code":
        parts.push(`<div style="text-align:center;padding:24px 0;">
  <div style="display:inline-block;border:2px solid #6366f1;border-radius:8px;padding:16px;">
    <span style="font-size:28px;font-weight:700;letter-spacing:6px;color:#6366f1;font-family:'Courier New',monospace;">{{.Code}}</span>
  </div>
</div>`);
        break;
      case "Expiry":
        parts.push(
          `<p style="color:#94a3b8;font-size:12px;text-align:center;font-style:italic;margin:0;">Expires in {{.Expiry}}</p>`,
        );
        break;
      case "Name":
        parts.push(
          `<p style="color:#374151;font-size:15px;margin:0 0 8px;">Hi <strong>{{.Name}}</strong>,</p>`,
        );
        break;
      case "Message":
        parts.push(`<div style="background:#f8fafc;border:1px solid #e5e7eb;border-radius:12px;padding:28px;margin:20px 0;">
  <p style="color:#374151;font-size:15px;line-height:1.6;margin:0;">{{.Message}}</p>
</div>`);
        break;
      case "Amount":
        parts.push(`<div style="text-align:center;padding:12px 0;">
  <span style="font-size:32px;font-weight:700;color:#6366f1;">{{.Amount}}</span>
</div>`);
        break;
      case "Date":
        parts.push(
          `<p style="color:#6b7280;font-size:13px;text-align:center;margin:4px 0;">{{.Date}}</p>`,
        );
        break;
      case "Plan":
        parts.push(
          `<p style="color:#374151;font-size:14px;margin:6px 0;"><strong>Plan:</strong> {{.Plan}}</p>`,
        );
        break;
      case "Status":
        parts.push(
          `<p style="color:#374151;font-size:14px;margin:6px 0;"><strong>Status:</strong> {{.Status}}</p>`,
        );
        break;
      case "Reason":
        parts.push(
          `<p style="color:#374151;font-size:14px;margin:6px 0;"><strong>Reason:</strong> {{.Reason}}</p>`,
        );
        break;
      case "Email":
        parts.push(
          `<p style="color:#374151;font-size:14px;margin:6px 0;"><strong>Email:</strong> {{.Email}}</p>`,
        );
        break;
      case "OrganizationName":
        parts.push(
          `<p style="color:#374151;font-size:14px;margin:6px 0;"><strong>Organization:</strong> {{.OrganizationName}}</p>`,
        );
        break;
      case "ActionURL":
        parts.push(`{{if .ActionURL}}
<div style="text-align:center;margin-top:28px;">
  <a href="{{.ActionURL}}" style="display:inline-block;background:#6366f1;color:#ffffff;font-size:14px;font-weight:600;text-decoration:none;padding:12px 24px;border-radius:8px;">
    {{.ActionText}}
  </a>
</div>
{{end}}`);
        break;
      case "FooterNote":
        parts.push(`{{if .FooterNote}}
<p style="margin-top:28px;color:#9ca3af;font-size:12px;text-align:center;line-height:1.4;">{{.FooterNote}}</p>
{{end}}`);
        break;
      case "Link":
        parts.push(`{{if .Link}}
<div style="text-align:center;margin-top:20px;">
  <a href="{{.Link}}" style="color:#6366f1;font-size:14px;text-decoration:underline;">{{.LinkText}}</a>
</div>
{{end}}`);
        break;
      default:
        parts.push(
          `<p style="color:#374151;font-size:15px;margin:8px 0;">{{.${key}}}</p>`,
        );
    }
  }

  return parts.join("\n\n");
}

function createTemplate() {
  const nameInput = document.getElementById("newTplName");
  const errorEl = document.getElementById("nameError");

  let rawName = nameInput.value
    .trim()
    .toLowerCase()
    .replace(/\s+/g, "_")
    .replace(/[^a-z0-9_]/g, "");

  if (!rawName) {
    nameInput.classList.add("error");
    errorEl.textContent = "Template name is required.";
    errorEl.classList.add("visible");
    nameInput.focus();
    return;
  }

  if (!rawName.endsWith(".html")) rawName += ".html";
  const id = rawName.replace(".html", "");

  if (templates.find((t) => t.id === id)) {
    nameInput.classList.add("error");
    errorEl.textContent = `"${rawName}" already exists.`;
    errorEl.classList.add("visible");
    nameInput.focus();
    return;
  }

  nameInput.classList.remove("error");
  errorEl.classList.remove("visible");

  const selectedFields = [...newChips];

  const fields = selectedFields.map((key) => ({
    key,
    label: `.${key}`,
    type: key === "Message" ? "textarea" : "text",
    optional: OPTIONAL_FIELDS.has(key),
    default: FIELD_DEFAULTS[key] ?? `Sample ${key}`,
  }));

  const content =
    selectedFields.length > 0
      ? scaffoldContent(selectedFields)
      : `<p style="color:#374151;font-size:15px;text-align:center;">Your email content goes here.</p>`;

  templates.push({ id, name: rawName, builtin: false, fields, content });

  closeModal();
  selectTemplate(id);
}

/* ══════════════════════════════════
   EVENT DELEGATION
══════════════════════════════════ */
function bindEvents() {
  // Sidebar: select / delete
  document.getElementById("tplList").addEventListener("click", (e) => {
    const delBtn = e.target.closest("[data-del]");
    if (delBtn) {
      e.stopPropagation();
      deleteTemplate(delBtn.dataset.del);
      return;
    }
    const item = e.target.closest(".tpl-item");
    if (item) selectTemplate(item.dataset.id);
  });

  // Add template button
  document.getElementById("addTplBtn").addEventListener("click", openModal);

  // Vars form: live update
  document.getElementById("varsForm").addEventListener("input", (e) => {
    const el = e.target;
    if (!el.dataset.key) return;
    const tpl = templates.find((t) => t.id === activeId);
    if (!tpl) return;
    getFieldData(tpl)[el.dataset.key] = el.value;
    refreshPreview();
    showSaved();
  });

  // Preview tabs
  document.querySelectorAll(".ptab").forEach((btn) => {
    btn.addEventListener("click", () => switchTab(btn.dataset.tab));
  });

  // Copy button
  document.getElementById("copyBtn").addEventListener("click", copyCurrentTab);

  // Save button
  document.getElementById("saveBtn").addEventListener("click", saveToTemplates);

  // Theme toggle
  document.getElementById("themeToggle").addEventListener("click", () => {
    darkMode = !darkMode;
    applyTheme();
  });

  // Modal: close on overlay click
  document.getElementById("modalOverlay").addEventListener("click", (e) => {
    if (e.target === document.getElementById("modalOverlay")) closeModal();
  });

  // Modal: close on Escape
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") closeModal();
    if (
      e.key === "Enter" &&
      document.getElementById("modalOverlay").classList.contains("open")
    ) {
      createTemplate();
    }
  });

  // Modal: cancel / create
  document.getElementById("modalCancel").addEventListener("click", closeModal);
  document
    .getElementById("modalCreate")
    .addEventListener("click", createTemplate);

  // Modal: chip toggle (event delegation)
  document.getElementById("fieldChips").addEventListener("click", (e) => {
    const chip = e.target.closest(".chip");
    if (chip) toggleChip(chip.dataset.field);
  });

  // Modal: clear name error on input
  document.getElementById("newTplName").addEventListener("input", () => {
    document.getElementById("newTplName").classList.remove("error");
    document.getElementById("nameError").classList.remove("visible");
  });
}

/* ══════════════════════════════════
   INIT
══════════════════════════════════ */
document.addEventListener("DOMContentLoaded", () => {
  applyTheme();
  bindEvents();
  selectTemplate("otp");
});
