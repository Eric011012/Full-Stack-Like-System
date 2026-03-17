import React, { useState } from "react";
import FeedbackPage from "./pages/FeedbackPage";
import MessagesPage from "./pages/MessagesPage";

export default function App() {
  const [page, setPage] = useState("feedback");

  return (
    <div>
      <nav style={{ background: "#ddd", padding: "10px", display: "flex", gap: "10px" }}>
        <button
          onClick={() => setPage("feedback")}
          style={{ fontWeight: page === "feedback" ? "bold" : "normal" }}
        >
          Feedback
        </button>
        <button
          onClick={() => setPage("messages")}
          style={{ fontWeight: page === "messages" ? "bold" : "normal" }}
        >
          Messages
        </button>
      </nav>
      {page === "feedback" && <FeedbackPage />}
      {page === "messages" && <MessagesPage />}
    </div>
  );
}
