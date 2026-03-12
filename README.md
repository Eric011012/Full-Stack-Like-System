# Full-Stack Like System (Prototype)

A lightweight full-stack application built to practice the integration of a React frontend, a Go backend, and a PostgreSQL database. This project serves as a foundational sandbox for learning end-to-end data flow and persistent storage.

## 🛠 Tech Stack
- Frontend: React.js (Vite)
- Backend: Golang (Standard Library net/http)
- Database: PostgreSQL (Containerized via Docker)
- Communication: RESTful API / JSON

## 🚀 Environment Setup

### 1. Database Initialization
Ensure your PostgreSQL instance is running. Execute the following SQL to prepare the table:

CREATE TABLE IF NOT EXISTS like_history (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

### 2. Backend Service (Go)
Navigate to the go-api directory and start the server:
go run main.go
(The API will be available at http://localhost:8080)

### 3. Frontend Application (React)
Navigate to the react-app directory, install dependencies, and start the development server:
npm install
npm run dev
(The UI will be available at http://localhost:5173)

## 🧠 Engineering Journal & Troubleshooting

1. The "Compiled Language" Hurdle (Go)
- Issue: Changes in Go routes or logic didn't reflect in the browser.
- Learning: Unlike interpreted languages (like JavaScript), Go must be re-compiled/re-started (Ctrl+C) to apply changes.

2. Database Connectivity (Error 500)
- Issue: "relation like_history does not exist".
- Learning: This occurs when the Go backend queries a table that hasn't been created yet. Always verify schema via SQL.

3. State Synchronization
- Issue: Page refresh would reset the counter to zero.
- Learning: Shifted to "Database-backed state." The frontend now fetches the COUNT(*) from PostgreSQL on load.

## 📂 Project Structure
- /react-app: Frontend UI components and API fetch logic.
- /go-api: Backend handlers and DB connection.
- /db: SQL initialization scripts.

---

## 🙈 Recommended .gitignore (Create a separate file for this)
node_modules/
dist/
main
*.exe
.vscode/
.env
.DS_Store
