const API = "http://127.0.0.1:5000/api";

async function apiFetch(path, options = {}) {
  const res = await fetch(`${API}${path}`, {
    headers: { "Content-Type": "application/json" },
    ...options,
  });
  if (!res.ok) throw new Error(`API error: ${res.status} ${path}`);
  return res.json();
}

async function loadBaseLayout() {
  const data = await apiFetch("/base_layout");
  return data.content;
}

async function loadTemplates() {
  const list = await apiFetch("/templates");
  const fetches = list
    .filter((t) => !t.builtin)
    .map((t) => apiFetch(`/templates/${t.name}`));
  const files = await Promise.all(fetches);
  return files.map((f) => ({
    id: f.name.replace(".html", ""),
    name: f.name,
    content: extractContent(f.content),
    raw: f.content,
  }));
}

async function apiSaveTemplate(name, goTemplateContent) {
  return apiFetch(`/templates/${name}`, {
    method: "PUT",
    body: JSON.stringify({ content: goTemplateContent }),
  });
}

async function apiDeleteTemplate(name) {
  return apiFetch(`/templates/${name}`, { method: "DELETE" });
}
