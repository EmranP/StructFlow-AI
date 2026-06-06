# StructFlow — AI Project Structure Generator

Frontend Angular 17 app for generating AI-powered project structures.

## Quick Start

```bash
npm install
npm start
# Opens at http://localhost:4200
```

## Configuration

Change API URL in each service file or create an environment:

```
src/app/services/auth.service.ts      → apiUrl
src/app/services/project.service.ts   → apiUrl
src/app/services/generation.service.ts → apiUrl
```

Default: `http://localhost:3000/api`

## Pages & Flow

```
/auth              → Login / Register
/projects          → Dashboard — list of all projects
/projects/:id      → Project detail + start generation
/generation/:id    → Live generation status (polls every 15s)
/structure/:genId  → View 3 templates (Simple / Medium / Enterprise) + download ZIP
```

## Stack

- Angular 17 (standalone components, signals)
- SCSS with CSS variables design system
- HttpClient with JWT interceptor
- Polling for generation status

## Features

- Auth (JWT stored in localStorage)
- Create / edit / delete projects
- Start AI generation
- Live status polling every 15 seconds
- Interactive file tree viewer (collapsible)
- 3 template variants (Simple, Medium, Enterprise)
- Download generated structure as ZIP
