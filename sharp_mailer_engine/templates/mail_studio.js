/* ═══════════════════════════════════════════════════
   SHARPNR MAIL STUDIO — mail_studio.js
═══════════════════════════════════════════════════ */

/* ── LOGO (base64 inline SVG — matches base_layout.html exactly) ── */
const LOGO_B64 =
  "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+Cjxzdmcgd2lkdGg9IjEwMjQiIGhlaWdodD0iMTAyNCIgdmlld0JveD0iMCAwIDEwMjQgMTAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayI+CiAgPHRpdGxlPkFwcCBJY29uPC90aXRsZT4KICA8ZGVmcz4KICAgIDxzdHlsZT4KICAgICAgLmJhY2tncm91bmQgeyBmaWxsOiAjMEEwQTJFOyB9CiAgICAgIC5mb3JlZ3JvdW5kIHsgZmlsbDogI0ZGRkZGRjsgfQogICAgPC9zdHlsZT4KICA8L2RlZnM+CiAgPHJlY3QgY2xhc3M9ImJhY2tncm91bmQiIHdpZHRoPSIxMDI0IiBoZWlnaHQ9IjEwMjQiLz4KICA8cGF0aCBjbGFzcz0iZm9yZWdyb3VuZCIgZD0iTSA1MTIgMTgwIEwgNzQwIDcyMCBMIDY0MCA3MjAgTCA2MDAgNjUzIEwgNDI0IDY1MyBMIDM4NCA3MjAgTCAyODQgNzIwIFoiLz4KPC9zdmc+";

/* ────────────────────────────────────────────────────────
   BASE_LAYOUT
   This is the single source of truth for the email wrapper.
   It is kept 1:1 in sync with base_layout.html on disk.
   __CONTENT__ is the slot replaced at preview/export time.
──────────────────────────────────────────────────────── */
function buildBaseLayout(contentSlot) {
  return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width,initial-scale=1.0"/>
  <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
  <title>Sharpnr</title>
</head>
<body style="margin:0;padding:0;background-color:#EAEDF5;font-family:-apple-system,BlinkMacSystemFont,'Helvetica Neue',Arial,sans-serif;-webkit-font-smoothing:antialiased;">
  <table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color:#EAEDF5;">
    <tr>
      <td align="center" style="padding:44px 16px 52px;">
        <table cellpadding="0" cellspacing="0" border="0" style="max-width:500px;width:100%;background:#ffffff;border-radius:14px;overflow:hidden;box-shadow:0 4px 28px rgba(10,10,46,0.14),0 1px 4px rgba(10,10,46,0.06);">

          <!-- HEADER -->
          <tr>
            <td style="background-color:#0A0A2E;padding:26px 36px;border-radius:14px 14px 0 0;">
              <table cellpadding="0" cellspacing="0" border="0">
                <tr>
                  <td style="width:38px;height:38px;border:1px solid rgba(255,255,255,0.18);border-radius:9px;text-align:center;vertical-align:middle;background-color:#0A0A2E;">
                    <img src="${LOGO_B64}" width="22" height="22" alt="Sharpnr" style="display:block;margin:8px auto;border:0;"/>
                  </td>
                  <td style="padding-left:11px;vertical-align:middle;">
                    <span style="color:#ffffff;font-size:19px;font-weight:700;letter-spacing:-0.5px;line-height:1;">Sharpnr</span>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- CONTENT -->
          <tr>
            <td style="padding:44px 36px 36px;">
              ${contentSlot}
            </td>
          </tr>

          <!-- FOOTER -->
          <tr>
            <td style="background-color:#F7F8FC;border-top:1px solid #EEF0F8;padding:20px 36px 22px;border-radius:0 0 14px 14px;">
              <table cellpadding="0" cellspacing="0" border="0" width="100%">
                <tr>
                  <td align="center" style="padding-bottom:9px;">
                    <a href="#" style="color:#9CA3C0;font-size:11px;text-decoration:none;margin:0 8px;">Unsubscribe</a>
                    <span style="color:#D1D5E8;font-size:11px;">·</span>
                    <a href="#" style="color:#9CA3C0;font-size:11px;text-decoration:none;margin:0 8px;">Privacy Policy</a>
                    <span style="color:#D1D5E8;font-size:11px;">·</span>
                    <a href="#" style="color:#9CA3C0;font-size:11px;text-decoration:none;margin:0 8px;">Help Centre</a>
                  </td>
                </tr>
                <tr>
                  <td align="center">
                    <p style="color:#C4C9DC;font-size:11px;line-height:1.6;margin:0;">
                      © 2026 Sharpnr · All rights reserved<br/>
                      You received this because you have an account with us.
                    </p>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>
</body>
</html>`;
}

/* ── BUILT-IN TEMPLATES ── */
const BUILTIN_TEMPLATES = [
  {
    id: "otp",
    name: "otp.html",
    builtin: true,
    fields: [
      {
        key: "Name",
        label: ".Name",
        type: "text",
        optional: false,
        default: "Rahul",
      },
      {
        key: "Code",
        label: ".Code",
        type: "text",
        optional: false,
        default: "847 291",
      },
      {
        key: "ExpiryMinutes",
        label: ".ExpiryMinutes",
        type: "text",
        optional: false,
        default: "10",
      },
    ],
    content: `<!-- Badge -->
<p style="display:inline-block;background:#EEF0FA;color:#0A0A2E;font-size:10px;font-weight:700;letter-spacing:1px;text-transform:uppercase;padding:4px 10px;border-radius:5px;margin:0 0 16px;">Security</p>

<!-- Heading -->
<h1 style="color:#0A0A2E;font-size:22px;font-weight:700;letter-spacing:-0.5px;line-height:1.3;margin:0 0 12px;">Your verification code</h1>

<!-- Intro -->
<p style="color:#5A607A;font-size:15px;line-height:1.7;margin:0 0 28px;">Hi <strong style="color:#0A0A2E;">{{.Name}}</strong>, use the code below to verify your identity. It is valid for <strong style="color:#0A0A2E;">{{.ExpiryMinutes}} minutes</strong> and can only be used once.</p>

<!-- OTP Box -->
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-bottom:28px;">
  <tr>
    <td style="background-color:#F7F8FC;border:1.5px solid #E2E5F0;border-radius:10px;padding:24px 16px;text-align:center;">
      <p style="color:#9CA3C0;font-size:10px;font-weight:700;letter-spacing:1.5px;text-transform:uppercase;margin:0 0 14px;">One-Time Password</p>
      <p style="color:#0A0A2E;font-size:38px;font-weight:800;letter-spacing:10px;font-family:'Courier New',Courier,monospace;margin:0;line-height:1;">{{.Code}}</p>
      <p style="color:#9CA3C0;font-size:12px;margin:14px 0 0;">Expires in {{.ExpiryMinutes}} minutes</p>
    </td>
  </tr>
</table>

<!-- Body note -->
<p style="color:#5A607A;font-size:14px;line-height:1.7;margin:0 0 22px;">Didn't request this? You can safely ignore this email — your account remains secure.</p>

<!-- Divider -->
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin:4px 0 20px;"><tr><td style="border-top:1px solid #F0F2F8;font-size:0;line-height:0;">&nbsp;</td></tr></table>

<!-- Security note -->
<p style="color:#9CA3C0;font-size:12px;line-height:1.65;margin:0;">🔒 Never share this code with anyone, including Sharpnr support.</p>`,
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
    content: `<!-- Badge -->
<p style="display:inline-block;background:#EEF0FA;color:#0A0A2E;font-size:10px;font-weight:700;letter-spacing:1px;text-transform:uppercase;padding:4px 10px;border-radius:5px;margin:0 0 16px;">Notification</p>

<!-- Heading -->
<h1 style="color:#0A0A2E;font-size:22px;font-weight:700;letter-spacing:-0.5px;line-height:1.3;margin:0 0 10px;">{{.Title}}</h1>

{{if .Subtitle}}
<p style="color:#5A607A;font-size:15px;line-height:1.7;margin:0 0 24px;">{{.Subtitle}}</p>
{{end}}

<!-- Message card -->
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin:4px 0 24px;">
  <tr>
    <td style="background-color:#F7F8FC;border:1px solid #EEF0F8;border-radius:10px;padding:22px 20px;">
      <p style="color:#3D4560;font-size:15px;line-height:1.7;margin:0;">{{.Message}}</p>
    </td>
  </tr>
</table>

{{if .ActionURL}}
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-bottom:24px;">
  <tr>
    <td align="center">
      <a href="{{.ActionURL}}" style="display:inline-block;background-color:#0A0A2E;color:#ffffff;font-size:15px;font-weight:600;text-decoration:none;letter-spacing:-0.2px;padding:14px 36px;border-radius:8px;line-height:1;">{{.ActionText}}</a>
    </td>
  </tr>
</table>
{{end}}

{{if .FooterNote}}
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-top:8px;"><tr><td style="border-top:1px solid #F0F2F8;font-size:0;line-height:0;">&nbsp;</td></tr></table>
<p style="color:#9CA3C0;font-size:12px;line-height:1.65;margin:16px 0 0;">{{.FooterNote}}</p>
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
        key: "BillingPeriod",
        label: ".BillingPeriod",
        type: "text",
        optional: true,
        default: "May 1 – May 31, 2026",
      },
      {
        key: "Link",
        label: ".Link",
        type: "text",
        optional: true,
        default: "https://sharpnr.com/billing",
      },
    ],
    content: `<!-- Badge -->
<p style="display:inline-block;background:#EEF0FA;color:#0A0A2E;font-size:10px;font-weight:700;letter-spacing:1px;text-transform:uppercase;padding:4px 10px;border-radius:5px;margin:0 0 16px;">Billing</p>

<!-- Heading -->
<h1 style="color:#0A0A2E;font-size:22px;font-weight:700;letter-spacing:-0.5px;line-height:1.3;margin:0 0 12px;">Payment confirmed</h1>
<p style="color:#5A607A;font-size:15px;line-height:1.7;margin:0 0 28px;">Hi <strong style="color:#0A0A2E;">{{.Name}}</strong>, your payment was processed successfully.</p>

<!-- Billing table -->
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color:#F7F8FC;border:1px solid #EEF0F8;border-radius:10px;margin-bottom:24px;">
  <tr>
    <td style="padding:13px 20px;border-bottom:1px solid #EEF0F8;">
      <table width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
          <td style="color:#9CA3C0;font-size:13px;">Plan</td>
          <td style="color:#0A0A2E;font-size:13px;font-weight:600;text-align:right;">{{.Plan}}</td>
        </tr>
      </table>
    </td>
  </tr>
  {{if .BillingPeriod}}
  <tr>
    <td style="padding:13px 20px;border-bottom:1px solid #EEF0F8;">
      <table width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
          <td style="color:#9CA3C0;font-size:13px;">Billing Period</td>
          <td style="color:#0A0A2E;font-size:13px;font-weight:600;text-align:right;">{{.BillingPeriod}}</td>
        </tr>
      </table>
    </td>
  </tr>
  {{end}}
  <tr>
    <td style="padding:13px 20px;border-bottom:1px solid #EEF0F8;">
      <table width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
          <td style="color:#9CA3C0;font-size:13px;">Date</td>
          <td style="color:#0A0A2E;font-size:13px;font-weight:600;text-align:right;">{{.Date}}</td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td style="padding:14px 20px;">
      <table width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
          <td style="color:#0A0A2E;font-size:14px;font-weight:700;">Total</td>
          <td style="color:#0A0A2E;font-size:18px;font-weight:800;text-align:right;letter-spacing:-0.5px;">{{.Amount}}</td>
        </tr>
      </table>
    </td>
  </tr>
</table>

{{if .Link}}
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-bottom:8px;">
  <tr>
    <td align="center">
      <a href="{{.Link}}" style="display:inline-block;background-color:#0A0A2E;color:#ffffff;font-size:15px;font-weight:600;text-decoration:none;letter-spacing:-0.2px;padding:14px 36px;border-radius:8px;line-height:1;">View Receipt</a>
    </td>
  </tr>
</table>
{{end}}`,
  },
];

/* ── AVAILABLE FIELDS FOR NEW TEMPLATE WIZARD ── */
const ALL_FIELDS = [
  "Title",
  "Subtitle",
  "Message",
  "Code",
  "ExpiryMinutes",
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
  "BillingPeriod",
];

const OPTIONAL_FIELDS = new Set([
  "Subtitle",
  "FooterNote",
  "ActionURL",
  "ActionText",
  "LinkText",
  "BillingPeriod",
]);

const FIELD_DEFAULTS = {
  Code: "847 291",
  ExpiryMinutes: "10",
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
  BillingPeriod: "May 1 – May 31, 2026",
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
   GO TEMPLATE RENDERER (client-side preview only)
   Handles: {{.Key}}  {{if .Key}}...{{end}}
══════════════════════════════════ */
function renderGoTemplate(content, data) {
  let out = content;

  // Strip Go-only block tags used in base_layout (not needed in preview)
  out = out.replace(/\{\{define "[^"]+"\}\}/g, "");
  out = out.replace(/\{\{end\}\}\s*$/g, "");

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
  const renderedContent = renderGoTemplate(tpl.content, data);
  return buildBaseLayout(renderedContent);
}

/* ── Produces the final .html file saved to disk ──
   This is what your Go mailer engine actually reads.
   It uses {{template "base_layout.html" .}} so the
   Go runtime provides the real header/footer — not
   the JS-rendered preview version.
── */
function getGoTemplate(tpl) {
  return `{{template "base_layout.html" .}}

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
   SCAFFOLD (used when creating a new template)
   All generated HTML uses the new #0A0A2E brand tokens.
══════════════════════════════════ */
function scaffoldContent(fields) {
  const parts = [];

  const order = [
    "Title",
    "Subtitle",
    "Code",
    "ExpiryMinutes",
    "Name",
    "Message",
    "Amount",
    "Date",
    "Plan",
    "BillingPeriod",
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

  // Auto-detect a badge category from the field set
  let badge = "Notification";
  if (fields.includes("Code")) badge = "Security";
  if (fields.includes("Amount")) badge = "Billing";

  parts.push(`<!-- Badge -->
<p style="display:inline-block;background:#EEF0FA;color:#0A0A2E;font-size:10px;font-weight:700;letter-spacing:1px;text-transform:uppercase;padding:4px 10px;border-radius:5px;margin:0 0 16px;">${badge}</p>`);

  for (const key of sorted) {
    switch (key) {
      case "Title":
        parts.push(
          `<h1 style="color:#0A0A2E;font-size:22px;font-weight:700;letter-spacing:-0.5px;line-height:1.3;margin:0 0 12px;">{{.Title}}</h1>`,
        );
        break;

      case "Subtitle":
        parts.push(`{{if .Subtitle}}
<p style="color:#5A607A;font-size:15px;line-height:1.7;margin:0 0 24px;">{{.Subtitle}}</p>
{{end}}`);
        break;

      case "Name":
        parts.push(
          `<p style="color:#5A607A;font-size:15px;line-height:1.7;margin:0 0 20px;">Hi <strong style="color:#0A0A2E;">{{.Name}}</strong>,</p>`,
        );
        break;

      case "Code":
        parts.push(`<!-- OTP Box -->
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-bottom:28px;">
  <tr>
    <td style="background-color:#F7F8FC;border:1.5px solid #E2E5F0;border-radius:10px;padding:24px 16px;text-align:center;">
      <p style="color:#9CA3C0;font-size:10px;font-weight:700;letter-spacing:1.5px;text-transform:uppercase;margin:0 0 14px;">One-Time Password</p>
      <p style="color:#0A0A2E;font-size:38px;font-weight:800;letter-spacing:10px;font-family:'Courier New',Courier,monospace;margin:0;line-height:1;">{{.Code}}</p>
    </td>
  </tr>
</table>`);
        break;

      case "ExpiryMinutes":
        parts.push(
          `<p style="color:#9CA3C0;font-size:12px;text-align:center;margin:-16px 0 24px;">Expires in {{.ExpiryMinutes}} minutes</p>`,
        );
        break;

      case "Message":
        parts.push(`<!-- Message card -->
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-bottom:24px;">
  <tr>
    <td style="background-color:#F7F8FC;border:1px solid #EEF0F8;border-radius:10px;padding:22px 20px;">
      <p style="color:#3D4560;font-size:15px;line-height:1.7;margin:0;">{{.Message}}</p>
    </td>
  </tr>
</table>`);
        break;

      case "Amount":
        parts.push(
          `<p style="color:#0A0A2E;font-size:32px;font-weight:800;letter-spacing:-1px;text-align:center;margin:0 0 4px;">{{.Amount}}</p>`,
        );
        break;

      case "Date":
        parts.push(
          `<p style="color:#9CA3C0;font-size:12px;text-align:center;margin:0 0 20px;">{{.Date}}</p>`,
        );
        break;

      case "Plan":
        parts.push(`<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color:#F7F8FC;border:1px solid #EEF0F8;border-radius:10px;margin-bottom:24px;">
  <tr><td style="padding:13px 20px;">
    <table width="100%" cellpadding="0" cellspacing="0" border="0">
      <tr>
        <td style="color:#9CA3C0;font-size:13px;">Plan</td>
        <td style="color:#0A0A2E;font-size:13px;font-weight:600;text-align:right;">{{.Plan}}</td>
      </tr>
    </table>
  </td></tr>
</table>`);
        break;

      case "BillingPeriod":
        parts.push(`{{if .BillingPeriod}}
<p style="color:#9CA3C0;font-size:12px;text-align:center;margin:-12px 0 20px;">{{.BillingPeriod}}</p>
{{end}}`);
        break;

      case "Status":
        parts.push(
          `<p style="color:#3D4560;font-size:14px;margin:6px 0;"><strong style="color:#0A0A2E;">Status:</strong> {{.Status}}</p>`,
        );
        break;

      case "Reason":
        parts.push(
          `<p style="color:#3D4560;font-size:14px;margin:6px 0;"><strong style="color:#0A0A2E;">Reason:</strong> {{.Reason}}</p>`,
        );
        break;

      case "Email":
        parts.push(
          `<p style="color:#3D4560;font-size:14px;margin:6px 0;"><strong style="color:#0A0A2E;">Email:</strong> {{.Email}}</p>`,
        );
        break;

      case "OrganizationName":
        parts.push(
          `<p style="color:#3D4560;font-size:14px;margin:6px 0;"><strong style="color:#0A0A2E;">Organization:</strong> {{.OrganizationName}}</p>`,
        );
        break;

      case "ActionURL":
        parts.push(`{{if .ActionURL}}
<!-- CTA Button -->
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin:24px 0 8px;">
  <tr>
    <td align="center">
      <a href="{{.ActionURL}}" style="display:inline-block;background-color:#0A0A2E;color:#ffffff;font-size:15px;font-weight:600;text-decoration:none;letter-spacing:-0.2px;padding:14px 36px;border-radius:8px;line-height:1;">{{.ActionText}}</a>
    </td>
  </tr>
</table>
{{end}}`);
        break;

      case "FooterNote":
        parts.push(`{{if .FooterNote}}
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin:16px 0 12px;"><tr><td style="border-top:1px solid #F0F2F8;font-size:0;line-height:0;">&nbsp;</td></tr></table>
<p style="color:#9CA3C0;font-size:12px;line-height:1.65;margin:0;">{{.FooterNote}}</p>
{{end}}`);
        break;

      case "Link":
        parts.push(`{{if .Link}}
<div style="text-align:center;margin-top:16px;">
  <a href="{{.Link}}" style="color:#0A0A2E;font-size:13px;font-weight:600;text-decoration:underline;">{{.LinkText}}</a>
</div>
{{end}}`);
        break;

      default:
        parts.push(
          `<p style="color:#3D4560;font-size:15px;margin:8px 0;">{{.${key}}}</p>`,
        );
    }
  }

  return parts.join("\n\n");
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
      ${!t.builtin ? `<button class="tpl-del" data-del="${t.id}" title="Delete">×</button>` : ""}
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
   SAVE  — "go template" tab → saves {{template "base_layout.html" .}} source
            other tabs       → saves the full rendered HTML preview
══════════════════════════════════ */
async function saveToTemplates() {
  const tpl = templates.find((t) => t.id === activeId);
  if (!tpl) return;

  const isGoTab = activeTab === "gotpl";
  const content = isGoTab ? getGoTemplate(tpl) : getRenderedHTML(tpl);
  const suggestedName = tpl.name;

  if (window.showSaveFilePicker) {
    try {
      const handle = await window.showSaveFilePicker({
        suggestedName,
        types: [
          { description: "HTML Template", accept: { "text/html": [".html"] } },
        ],
        startIn: "documents",
      });
      const writable = await handle.createWritable();
      await writable.write(content);
      await writable.close();
      showSaved("saved!");
    } catch (err) {
      if (err.name !== "AbortError") fallbackDownload(content, suggestedName);
    }
  } else {
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
  if (activeId === id) activeId = templates[0]?.id ?? null;
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
      : `<p style="color:#5A607A;font-size:15px;text-align:center;margin:0;">Your email content goes here.</p>`;

  templates.push({ id, name: rawName, builtin: false, fields, content });
  closeModal();
  selectTemplate(id);
}

/* ══════════════════════════════════
   EVENTS
══════════════════════════════════ */
function bindEvents() {
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

  document.getElementById("addTplBtn").addEventListener("click", openModal);

  document.getElementById("varsForm").addEventListener("input", (e) => {
    const el = e.target;
    if (!el.dataset.key) return;
    const tpl = templates.find((t) => t.id === activeId);
    if (!tpl) return;
    getFieldData(tpl)[el.dataset.key] = el.value;
    refreshPreview();
    showSaved();
  });

  document.querySelectorAll(".ptab").forEach((btn) => {
    btn.addEventListener("click", () => switchTab(btn.dataset.tab));
  });

  document.getElementById("copyBtn").addEventListener("click", copyCurrentTab);
  document.getElementById("saveBtn").addEventListener("click", saveToTemplates);

  document.getElementById("themeToggle").addEventListener("click", () => {
    darkMode = !darkMode;
    applyTheme();
  });

  document.getElementById("modalOverlay").addEventListener("click", (e) => {
    if (e.target === document.getElementById("modalOverlay")) closeModal();
  });

  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") closeModal();
    if (
      e.key === "Enter" &&
      document.getElementById("modalOverlay").classList.contains("open")
    ) {
      createTemplate();
    }
  });

  document.getElementById("modalCancel").addEventListener("click", closeModal);
  document
    .getElementById("modalCreate")
    .addEventListener("click", createTemplate);

  document.getElementById("fieldChips").addEventListener("click", (e) => {
    const chip = e.target.closest(".chip");
    if (chip) toggleChip(chip.dataset.field);
  });

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
