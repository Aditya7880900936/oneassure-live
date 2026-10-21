
import { useState } from "react";
import DOMPurify from "dompurify";
import { saveMarkdown } from "./api/api";
import "./App.css";

function App() {
  const [markdown, setMarkdown] = useState("# Hello World\n\nWrite some **Markdown** here.");
  const [html, setHtml] = useState("");
  const [isPreviewing, setIsPreviewing] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    setIsLoading(true);
    setError("");

    try {
      const result = await saveMarkdown(markdown);

      if (typeof result.html !== "string") {
        throw new Error("Invalid response received from server.");
      }

      setHtml(result.html);
      setIsPreviewing(true);
    } catch (err) {
      setError(err.message || "Unable to convert Markdown.");
      setIsPreviewing(false);
    } finally {
      setIsLoading(false);
    }
  };

  const safeHTML = DOMPurify.sanitize(html);

  return (
    <main className="app">
      <h1>Markdown Previewer</h1>
      <p>Write Markdown and preview the converted HTML.</p>

      <form onSubmit={handleSubmit}>
        <label htmlFor="markdown">Markdown input</label>

        <textarea
          id="markdown"
          name="markdown"
          rows={14}
          value={markdown}
          onChange={(e) => setMarkdown(e.target.value)}
          placeholder="Enter your Markdown here..."
        />

        <div className="actions">
          <button type="submit" disabled={isLoading}>
            {isLoading ? "Converting..." : "Preview"}
          </button>

          {isPreviewing && (
            <button
              type="button"
              className="secondary"
              onClick={() => setIsPreviewing(false)}
            >
              Hide Preview
            </button>
          )}
        </div>
      </form>

      {error && (
        <p className="error" role="alert">
          {error}
        </p>
      )}

      {isPreviewing && (
        <section className="preview">
          <h2>Preview</h2>
          <div
            className="preview-content"
            dangerouslySetInnerHTML={{ __html: safeHTML }}
          />
        </section>
      )}
    </main>
  );
}

export default App;
