# Kivia Frontend

Dashboard for the Kivia API observability platform. View projects, manage API keys, and browse request logs with filtering and pagination.

## Tech Stack

| Category         | Technology                                  |
| ---------------- | ------------------------------------------- |
| Framework        | Next.js 16 (App Router, Turbopack)          |
| Language         | TypeScript, React 19                        |
| Styling          | Tailwind CSS 4 (oklch color space)          |
| Components       | shadcn/ui (base-nova style)                 |
| State            | React Query v5 (@tanstack/react-query)      |
| Forms            | React Hook Form + Zod                       |
| Icons            | Lucide React                                |
| Notifications    | Sonner                                      |
| Fonts            | Roboto, JetBrains Mono                      |

## Getting Started

### Prerequisites

- Node.js 18+
- Backend running at `http://localhost:8080` (see root README)

### Install & Run

```bash
npm install
npm run dev
```

Open [http://localhost:3000](http://localhost:3000).

### Scripts

```bash
npm run dev     # Development server
npm run build   # Production build
npm start       # Production server
npm run lint    # ESLint
```

## Pages

| Route              | Description                                      | Guard     |
| ------------------ | ------------------------------------------------ | --------- |
| /                  | Landing page                                     | Public    |
| /login             | Login form                                       | Guest     |
| /register          | Registration form                                | Guest     |
| /dashboard         | Overview with stats and recent projects           | Auth      |
| /projects          | Project list with create dialog                  | Auth      |
| /projects/[id]     | Project detail — API keys tab and logs tab        | Auth      |

- **AuthGuard** — Redirects unauthenticated users to `/login`
- **GuestGuard** — Redirects authenticated users to `/dashboard`

## Project Structure

```
src/
├── app/
│   ├── layout.tsx              # Root layout (providers, fonts, dark mode)
│   ├── page.tsx                # Landing page
│   ├── globals.css             # Theme variables & global styles
│   ├── login/page.tsx
│   ├── register/page.tsx
│   ├── dashboard/
│   │   ├── layout.tsx          # Sidebar layout (protected)
│   │   └── page.tsx
│   └── projects/
│       ├── layout.tsx          # Sidebar layout (protected)
│       ├── page.tsx
│       └── [id]/page.tsx
├── components/
│   ├── AuthGuard.tsx
│   ├── GuestGuard.tsx
│   ├── QueryProvider.tsx
│   ├── AppSidebar.tsx
│   └── ui/                     # shadcn/ui components
└── lib/
    ├── api.ts                  # API client with token refresh
    ├── auth.ts                 # Token storage helpers
    └── utils.ts                # cn() utility
```

## API Client

The API client in `lib/api.ts` handles:

- **Bearer token injection** from localStorage on every request
- **Automatic token refresh** on 401 responses with race condition protection (single shared refresh promise)
- **Redirect to /login** when refresh fails

## Configuration

The API base URL is currently set to `http://localhost:8080` in `lib/api.ts`. Update this for other environments.
