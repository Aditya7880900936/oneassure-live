
const API_URL = "http://localhost:8080/markdown";

async function parseResponse(response) {
  const data = await response.json().catch(() => ({}));

  if (!response.ok) {
    throw new Error(data.error || `Request failed: ${response.status}`);
  }

  return data;
}

export async function fetchMarkdown(markdown) {
  const url = new URL(API_URL);
  url.searchParams.set("content", markdown);

  const response = await fetch(url);
  return parseResponse(response);
}

export async function saveMarkdown(markdown) {
  const response = await fetch(API_URL, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      content: markdown,
    }),
  });

  return parseResponse(response);
}
