# Kivia

API observability platform that monitors and logs HTTP requests from your applications. Integrate the Go SDK into your services to automatically capture request metadata and view it through the dashboard.

## Architecture

```
┌─────────────┐     ┌───────┐     ┌──────────────┐     ┌──────────┐     ┌────────────┐
│  Your App   │────▶│ NGINX │────▶│ Kivia Service │────▶│ RabbitMQ │────▶│ PostgreSQL │
│ (Kivia SDK) │     │       │     │  (Fiber v3)   │     │          │     │            │
└─────────────┘     └───────┘     └──────────────┘     └──────────┘     └────────────┘
                        │
                        ▼
                 ┌──────────────┐
                 │Email Service │
                 └──────────────┘
```

- **NGINX** — Reverse proxy with rate limiting, CORS, and JWT validation via `auth_request`
- **Kivia Service** — Go/Fiber API handling auth, projects, API keys, and log ingestion
- **RabbitMQ** — Async log processing queue (`log_queue`) for non-blocking ingestion
- **PostgreSQL** — Primary data store with auto-migrations on startup
- **Email Service** — Fiber-based email microservice

## Tech Stack

| Component   | Technology                  |
| ----------- | --------------------------- |
| Language    | Go 1.25                     |
| Framework   | Fiber v3                    |
| Database    | PostgreSQL 18 (pgxpool)     |
| Queue       | RabbitMQ (amqp091-go)       |
| Auth        | JWT (access + refresh tokens) |
| Migrations  | golang-migrate/migrate v4   |
| Proxy       | NGINX                       |
| Containers  | Docker Compose              |

## Getting Started

### Prerequisites

- Docker & Docker Compose

### Run

```bash
docker-compose up --build
```

Services will be available at:

| Service          | URL                     |
| ---------------- | ----------------------- |
| API (via NGINX)  | http://localhost:8080    |
| RabbitMQ Console | http://localhost:15672   |
| PostgreSQL       | localhost:5000           |

### Environment Variables

Create a `.env.dev` in `kivia_service/`:

```env
DATABASE_URL=postgresql://postgres:@postgres:5432/kivia
PORT=8081
JWT_ACCESS_TOKEN_SECRET=<64-char-hex>
JWT_REFRESH_TOKEN_SECRET=<64-char-hex>
RABBITMQ_CONNECTION_URL=amqp://guest:guest@rabbitmq:5672
APP_ENV=dev
```

## API Routes

### Public — Auth

| Method | Route            | Description             |
| ------ | ---------------- | ----------------------- |
| POST   | /auth/register   | Register a new user     |
| POST   | /auth/login      | Login, returns tokens   |
| POST   | /auth/refresh    | Refresh access token    |

### Protected — JWT Required

| Method | Route                          | Description                    |
| ------ | ------------------------------ | ------------------------------ |
| POST   | /api/v1/projects/create        | Create a project               |
| GET    | /api/v1/projects/all           | List user's projects           |
| POST   | /api/v1/api-keys/create        | Create an API key for project  |
| GET    | /api/v1/api-keys/all/:projectId | List API keys for project     |
| PATCH  | /api/v1/api-keys/revoke/:id    | Revoke an API key              |
| GET    | /api/v1/logs/all/:projectId    | Get paginated logs             |

### Protected — API Key Required

| Method | Route                 | Description                          |
| ------ | --------------------- | ------------------------------------ |
| POST   | /api/v1/logs/create   | Ingest a log entry (returns 202)     |

## SDK Integration

Install the Go SDK:

```bash
go get github.com/winnerx0/kivia-sdk
```

Wrap your HTTP handlers to auto-capture requests:

```go
package main

import (
    "net/http"
    kiviaSdk "github.com/winnerx0/kivia-sdk"
)

func main() {
    client := kiviaSdk.NewClient("your-api-key")

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello"))
    })

    http.ListenAndServe(":8080", client.NewLog(mux))
}
```

The SDK captures path, status code, IP address, timestamp, and latency for each request and sends it asynchronously to the Kivia API.

## Project Structure

```
kivia/
├── kivia_service/          # Main backend API
│   ├── cmd/kivia/          # Entry point
│   ├── api/                # Server & route setup
│   ├── internal/
│   │   ├── auth/           # JWT authentication
│   │   ├── user/           # User management
│   │   ├── project/        # Project CRUD
│   │   ├── api_key/        # API key management
│   │   ├── log/            # Log ingestion & retrieval
│   │   ├── refresh_token/  # Token refresh storage
│   │   ├── middleware/     # Auth & API key middleware
│   │   ├── database/       # PostgreSQL connection pool
│   │   ├── config/         # Environment config
│   │   ├── rabbitmq/       # RabbitMQ client & setup
│   │   └── utils/          # Error types
│   └── migrations/         # SQL migrations (auto-run)
├── kivia-sdk/              # Go SDK for client integration
├── email_service/          # Email microservice
├── nginx/                  # NGINX reverse proxy config
├── frontend/               # Next.js dashboard (separate repo)
└── docker-compose.yml
```

## How It Works

1. Your application uses the Kivia SDK which wraps your HTTP handlers
2. Each request is captured and sent to `POST /api/v1/logs/create` with your API key
3. The log is published to a RabbitMQ queue and a `202 Accepted` is returned immediately
4. A consumer goroutine processes the queue and persists logs to PostgreSQL
5. View your logs in the dashboard with filtering and pagination
