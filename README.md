# Markdown to HTML Converter

A full-stack Markdown Previewer built with **Go, Gin, Goldmark, and React**. Write Markdown and convert it into HTML through a REST API, then preview the rendered output in the browser.

## Features

- Convert Markdown into HTML using Goldmark.
- Preview formatted Markdown in a React interface.
- Support GitHub Flavored Markdown (GFM), including tables, strikethrough, and task lists.
- REST API built with Go and Gin.
- Validate empty and oversized Markdown input.
- Handle malformed requests and API errors.
- Sanitize HTML before rendering it in the frontend.
- Unit tests for Markdown conversion and input validation.

## Tech Stack

| Technology | Purpose |
|---|---|
| Go | Backend language |
| Gin | HTTP routing and API handling |
| Goldmark | Markdown parsing and HTML rendering |
| React | Frontend interface |
| JavaScript | API integration and UI interactions |
| DOMPurify | Sanitizing HTML before preview rendering |

## Project Structure

```text
practice-1/
├── backend/
│   ├── main.go
│   ├── go.mod
│   └── internal/
│       ├── handler/
│       │   └── markdownHandler.go
│       ├── model/
│       │   └── markdown.go
│       └── services/
│           ├── markdownService.go
│           └── markdown_service_test.go
├── frontend/
│   ├── src/
│   │   ├── api/
│   │   │   └── api.js
│   │   ├── App.jsx
│   │   └── App.css
│   └── package.json
├── .gitignore
└── README.md
```

## Prerequisites

Install the following before running the project:

- [Go](https://go.dev/dl/)
- [Node.js and npm](https://nodejs.org/)
- [Git](https://git-scm.com/)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Aditya7880900936/oneassure-live.git
cd oneassure-live
```

### 2. Start the backend

Open a terminal:

```bash
cd backend
go mod tidy
go run .
```

The backend starts on:

`http://localhost:8080`

### 3. Start the frontend

Open another terminal from the repository root:

```bash
cd frontend
npm install
npm run dev
```

Open the local development URL displayed by Vite, typically `http://localhost:5173`.

Ensure the backend is running before using the preview feature.

## API Documentation

### Health Check

**Endpoint:** `GET /health`

Checks whether the backend is running.

Example:

```bash
curl http://localhost:8080/health
```

Example response:

```json
{
  "status": "ok"
}
```

### Convert Markdown to HTML

**Endpoint:** `POST /markdown`

**Content-Type:** `application/json`

Converts Markdown provided in the request body into HTML.

Example request:

```json
{
  "content": "# Hello World\n\nThis is **bold** text."
}
```

Example response:

```json
{
  "html": "<h1>Hello World</h1>\n<p>This is <strong>bold</strong> text.</p>\n"
}
```

### Convert Markdown Using GET

**Endpoint:** `GET /markdown?content=...`

Accepts Markdown through the `content` query parameter.

Example:

```bash
curl -G http://localhost:8080/markdown \
  --data-urlencode 'content=# Hello World'
```

The response contains the converted HTML in the `html` field.

## Example Markdown

Input:

```markdown
# Markdown Previewer

A simple **Markdown to HTML converter**.

## Features

- Headings and paragraphs
- **Bold** and *italic* text
- Ordered and unordered lists
- [Links](https://example.com)

| Feature | Supported |
|---|---|
| Markdown conversion | Yes |
| HTML preview | Yes |

- [x] Convert Markdown
- [x] Render HTML
- [ ] Add more features
```

The frontend sends the Markdown to the backend and displays the returned HTML in the preview panel.

## Running Tests

Run the backend tests from the backend directory:

```bash
go test ./...
```

Run the service tests with verbose output:

```bash
go test ./internal/services -v
```

The service tests cover Markdown headings, bold text, lists, strikethrough, raw HTML handling, empty input, and oversized input.

## Error Handling

The API validates incoming requests and handles invalid input.

| Scenario | Expected behavior |
|---|---|
| Invalid JSON | `400 Bad Request` |
| Missing content | `400 Bad Request` |
| Empty Markdown | `400 Bad Request` |
| Markdown exceeding the configured size limit | `413 Request Entity Too Large` |
| Successful conversion | `200 OK` |

## Security Considerations

- Raw HTML is not enabled by default in the Goldmark renderer.
- The frontend sanitizes generated HTML with DOMPurify before inserting it into the preview.
- The API applies a size limit to Markdown input.
- Production deployments should configure CORS to allow only trusted frontend origins.

## Future Improvements

- Add automated frontend tests.
- Add Docker support for easier deployment.
- Introduce CI workflows for formatting, testing, and builds.
- Improve preview styling and accessibility.
- Add persistent storage only if saving and retrieving Markdown documents becomes a requirement.

## Author

**Aditya Sanskar Srivastava**

GitHub: [Aditya7880900936](https://github.com/Aditya7880900936)

## License

This project was developed as a live coding exercise and demonstration of full-stack development using Go and React.
