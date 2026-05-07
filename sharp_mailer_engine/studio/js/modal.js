/* ══════════════════════════════════════════════════
   modal.js — new template creation wizard
══════════════════════════════════════════════════ */

let selectedStruct = null;
let selectedFields = [];

function openModal() {
  selectedStruct = null;
  selectedFields = [];

  const nameInput = document.getElementById("newTplName");
  nameInput.value = "";
  nameInput.classList.remove("error");
  document.getElementById("nameError").classList.remove("visible");

  renderStructPicker();
  renderFieldChips();

  document.getElementById("modalOverlay").classList.add("open");
  setTimeout(() => nameInput.focus(), 150);
}

function closeModal() {
  document.getElementById("modalOverlay").classList.remove("open");
}

/* ── Step 1: struct picker chips ── */
function renderStructPicker() {
  const wrap = document.getElementById("structPicker");
  const structs = Object.keys(state.fields || {});

  if (structs.length === 0) {
    wrap.innerHTML = `<span class="empty-chips">No Go structs found — is file_server.py running?</span>`;
    return;
  }

  wrap.innerHTML = structs
    .map(
      (name) => `
    <div class="chip ${selectedStruct === name ? "on" : ""}" data-struct="${name}">${name}</div>
  `,
    )
    .join("");
}

/* ── Step 2: field chips for selected struct ── */
function renderFieldChips() {
  const wrap = document.getElementById("fieldChips");

  if (!selectedStruct) {
    wrap.innerHTML = `<span class="empty-chips">← pick a data type first</span>`;
    return;
  }

  selectedFields = state.fields[selectedStruct] || [];

  if (selectedFields.length === 0) {
    wrap.innerHTML = `<span class="empty-chips">No fields found in this struct.</span>`;
    return;
  }

  wrap.innerHTML = selectedFields
    .map(
      (f) => `
    <div class="chip on" data-field="${f.key}">
      .${f.key}${f.optional ? " <span style='opacity:.5;font-size:10px;'>(opt)</span>" : ""}
    </div>
  `,
    )
    .join("");
}

/* ── Create template → scaffold → save to disk ── */
async function createTemplate() {
  const nameInput = document.getElementById("newTplName");
  const errorEl = document.getElementById("nameError");

  // Normalise name
  let rawName = nameInput.value
    .trim()
    .toLowerCase()
    .replace(/\s+/g, "_")
    .replace(/[^a-z0-9_]/g, "");

  if (!rawName) {
    showFieldError(nameInput, errorEl, "Template name is required.");
    return;
  }

  if (!rawName.endsWith(".html")) rawName += ".html";
  const id = rawName.replace(".html", "");

  if (state.templates.find((t) => t.id === id)) {
    showFieldError(nameInput, errorEl, `"${rawName}" already exists.`);
    return;
  }

  if (!selectedStruct || selectedFields.length === 0) {
    errorEl.textContent = "Pick a data type and fields first.";
    errorEl.classList.add("visible");
    return;
  }

  const content = scaffoldContent(selectedStruct, selectedFields);
  const goSource = buildGoTemplate(content);

  const btn = document.getElementById("modalCreate");
  btn.textContent = "creating…";
  btn.disabled = true;
  errorEl.classList.remove("visible");

  try {
    await apiSaveTemplate(rawName, goSource);

    state.templates.push({ id, name: rawName, content, raw: goSource });
    closeModal();
    renderSidebar();
    selectTemplate(id);
    setStatus("created!", "ok");
  } catch (err) {
    errorEl.textContent = `Save failed: ${err.message}`;
    errorEl.classList.add("visible");
  } finally {
    btn.textContent = "create template";
    btn.disabled = false;
  }
}

function showFieldError(input, errorEl, msg) {
  input.classList.add("error");
  errorEl.textContent = msg;
  errorEl.classList.add("visible");
  input.focus();
}

/* ── Scaffold HTML content from struct fields ── */
function scaffoldContent(structName, fields) {
  const parts = [];
  const badge = structName.replace(/Data$/i, "");

  parts.push(`<!-- Badge -->
<p style="display:inline-block;background:#EEF0FA;color:#0A0A2E;font-size:10px;font-weight:700;letter-spacing:1px;text-transform:uppercase;padding:4px 10px;border-radius:5px;margin:0 0 16px;">${badge}</p>`);

  for (const f of fields) {
    const tag = `{{.${f.key}}}`;

    switch (f.key) {
      case "Name":
        parts.push(
          `<p style="color:#5A607A;font-size:15px;line-height:1.7;margin:0 0 20px;">Hi <strong style="color:#0A0A2E;">${tag}</strong>,</p>`,
        );
        break;

      case "Title":
        parts.push(
          `<h1 style="color:#0A0A2E;font-size:22px;font-weight:700;letter-spacing:-0.5px;line-height:1.3;margin:0 0 12px;">${tag}</h1>`,
        );
        break;

      case "Subtitle":
        parts.push(
          `{{if .Subtitle}}\n<p style="color:#5A607A;font-size:15px;line-height:1.7;margin:0 0 24px;">${tag}</p>\n{{end}}`,
        );
        break;

      case "Message":
        parts.push(`<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-bottom:24px;">
  <tr>
    <td style="background-color:#F7F8FC;border:1px solid #EEF0F8;border-radius:10px;padding:22px 20px;">
      <p style="color:#3D4560;font-size:15px;line-height:1.7;margin:0;">${tag}</p>
    </td>
  </tr>
</table>`);
        break;

      case "Code":
        parts.push(`<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin-bottom:28px;">
  <tr>
    <td style="background-color:#F7F8FC;border:1.5px solid #E2E5F0;border-radius:10px;padding:24px 16px;text-align:center;">
      <p style="color:#9CA3C0;font-size:10px;font-weight:700;letter-spacing:1.5px;text-transform:uppercase;margin:0 0 14px;">One-Time Password</p>
      <p style="color:#0A0A2E;font-size:38px;font-weight:800;letter-spacing:10px;font-family:'Courier New',Courier,monospace;margin:0;line-height:1;">${tag}</p>
    </td>
  </tr>
</table>`);
        break;

      case "ExpiryMinutes":
        parts.push(
          `<p style="color:#9CA3C0;font-size:12px;text-align:center;margin:-16px 0 24px;">Expires in ${tag} minutes</p>`,
        );
        break;

      case "Amount":
        parts.push(
          `<p style="color:#0A0A2E;font-size:32px;font-weight:800;letter-spacing:-1px;text-align:center;margin:0 0 4px;">${tag}</p>`,
        );
        break;

      case "ActionURL":
        parts.push(`{{if .ActionURL}}
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin:24px 0 8px;">
  <tr>
    <td align="center">
      <a href="${tag}" style="display:inline-block;background-color:#0A0A2E;color:#ffffff;font-size:15px;font-weight:600;text-decoration:none;padding:14px 36px;border-radius:8px;line-height:1;">{{.ActionText}}</a>
    </td>
  </tr>
</table>
{{end}}`);
        break;

      case "ActionText":
        break; // rendered inside ActionURL block above

      case "FooterNote":
        parts.push(`{{if .FooterNote}}
<table cellpadding="0" cellspacing="0" border="0" width="100%" style="margin:16px 0 12px;">
  <tr><td style="border-top:1px solid #F0F2F8;font-size:0;line-height:0;">&nbsp;</td></tr>
</table>
<p style="color:#9CA3C0;font-size:12px;line-height:1.65;margin:0;">${tag}</p>
{{end}}`);
        break;

      case "Link":
        parts.push(`{{if .Link}}
<div style="text-align:center;margin-top:16px;">
  <a href="${tag}" style="color:#0A0A2E;font-size:13px;font-weight:600;text-decoration:underline;">{{.LinkText}}</a>
</div>
{{end}}`);
        break;

      case "LinkText":
        break; // rendered inside Link block above

      default:
        parts.push(`<p style="color:#3D4560;font-size:14px;margin:6px 0;">
  <strong style="color:#0A0A2E;">${f.key}:</strong> ${tag}
</p>`);
    }
  }

  return parts.join("\n\n");
}

/* ── Event bindings ── */
function bindModalEvents() {
  document.getElementById("addTplBtn").addEventListener("click", openModal);
  document.getElementById("modalCancel").addEventListener("click", closeModal);
  document
    .getElementById("modalCreate")
    .addEventListener("click", createTemplate);

  // Struct picker
  document.getElementById("structPicker").addEventListener("click", (e) => {
    const chip = e.target.closest("[data-struct]");
    if (!chip) return;
    selectedStruct = chip.dataset.struct;
    renderStructPicker();
    renderFieldChips();
  });

  // Close on overlay backdrop click
  document.getElementById("modalOverlay").addEventListener("click", (e) => {
    if (e.target === document.getElementById("modalOverlay")) closeModal();
  });

  // Keyboard shortcuts
  document.addEventListener("keydown", (e) => {
    const isOpen = document
      .getElementById("modalOverlay")
      .classList.contains("open");
    if (e.key === "Escape" && isOpen) closeModal();
    if (e.key === "Enter" && isOpen) createTemplate();
  });

  // Clear name error on input
  document.getElementById("newTplName").addEventListener("input", () => {
    document.getElementById("newTplName").classList.remove("error");
    document.getElementById("nameError").classList.remove("visible");
  });
}
