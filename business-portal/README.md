# Mortify Business CRM Suite MVP

A contractor-ready SaaS-style business operations portal covering the full customer journey: Lead → Opportunity → Quote → Customer → Project → Helpdesk Ticket → Invoice → Payment. It also includes real IT monitoring for websites, SSL, DNS, TCP, HTTP/HTTPS, keyword, ICMP, and stale agent checks, plus agent metrics, alerting, ticket creation, and operations modules for email ingestion, documents, mobile devices, remote control, patch management, and SIEM correlation.

## Architecture

- **Frontend:** React + Vite, responsive SaaS sidebar layout, reusable tables/cards/badges, and Vite environment typing for `import.meta.env`.
- **Backend:** Node.js + Express REST API with JWT auth, PostgreSQL persistence, role checks, rate limits, centralized errors, and ESM TypeScript compilation.
- **Database:** PostgreSQL. SQL migrations in `backend/migrations` are the source of truth; Prisma schema is a reference only.
- **Worker:** `backend/src/workers/monitoringWorker.ts` runs out-of-band HTTP/HTTPS, DNS, TCP, SSL, keyword, ICMP, and stale agent checks.
- **Agent:** Go + gopsutil collector for Windows/Linux host inventory and metrics, sending to `/api/agent/metrics` with an asset-scoped token.

## Repository hygiene

The repo includes a top-level `.gitignore` for generated dependencies and local artifacts, including:

- `node_modules/`
- build outputs such as `dist/`, `build/`, `coverage/`, and `.vite/`
- log files
- local `.env` overrides while preserving `.env.example` files
- common OS/editor files

Do not commit generated dependency directories. Commit lockfiles and module checksum files, including the Go agent `go.sum`, when dependencies change.

## Quick start with Docker Compose

```bash
cd business-portal
docker compose up --build
```

Services:

- Frontend: <http://localhost:5173>
- Backend API: <http://localhost:4000/api>
- PostgreSQL: `localhost:5432`

Default seeded admin:

- Email: `admin@example.com`
- Password: `ChangeMe123!`

## Environment variables

Copy examples if running outside Docker:

```bash
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
```

The backend includes placeholders for future Xero, QuickBooks, Sage, Stripe, and AI provider integrations. The MVP AI provider is deterministic templates and requires no external API key.

The frontend reads `VITE_API_URL` via Vite's `import.meta.env`. The default frontend API client falls back to `http://localhost:4000/api` when `VITE_API_URL` is not set.

## Database migrations and seed

```bash
cd business-portal/backend
npm install
npm run migrate
npm run seed
```

Migrations create users, CRM, sales, projects, helpdesk, finance, monitoring, indexes, `global_search`, and operations tables for email ingestion, documents, mobile devices, remote control sessions, patch jobs, and SIEM correlations.

## Run backend

```bash
cd business-portal/backend
npm install
npm run dev
```

The backend package is configured as ESM (`"type": "module"`) to match the NodeNext TypeScript configuration and the use of `import.meta`/top-level `await` in scripts such as migrations and seed.

## Run frontend

```bash
cd business-portal/frontend
npm install
npm run dev
```

## Run monitoring worker

```bash
cd business-portal/backend
npm run worker
```

The worker loads due enabled checks, stores historical results, evaluates alert rules, deduplicates open alerts, resolves healthy alerts, and auto-creates helpdesk tickets when configured.

ICMP checks use the host `ping` command with platform-specific flags. The worker records a critical result if the command is unavailable or the target cannot be reached within the configured timeout.

## Build and run Go agent

```bash
cd business-portal/agent
go mod tidy
go build -o mortify-agent ./cmd/agent
cp config.example.yaml config.yaml
./mortify-agent -config config.yaml
```

Create a monitoring asset and token in the portal first. Put the raw token shown once into `config.yaml`. The backend stores only a bcrypt hash.

The agent uses gopsutil for CPU, memory, disk, host, network, and process telemetry. Service checks use OS-specific commands where available (`systemctl` on Linux, `sc query` on Windows) and fall back to process-name detection.

## Operations modules

The portal now includes database-backed API surfaces for managed service provider / IT operations workflows:

| Capability | Table | API prefix |
| --- | --- | --- |
| Email ingestion | `email_ingestions` | `/api/monitoring/email-ingestions` |
| Document management | `documents` | `/api/monitoring/documents` |
| Mobile devices | `mobile_devices` | `/api/monitoring/mobile-devices` |
| Remote control sessions | `remote_control_sessions` | `/api/monitoring/remote-control-sessions` |
| Patch management | `patch_jobs` | `/api/monitoring/patch-jobs` |
| SIEM correlation | `siem_correlations` | `/api/monitoring/siem-correlations` |

All routes use the existing authenticated monitoring permissions. SIEM correlations also expose a resolve action:

```bash
curl -X POST http://localhost:4000/api/monitoring/siem-correlations/$CORRELATION_ID/resolve \
  -H "authorization: Bearer $TOKEN"
```

## Example API calls

Login:

```bash
curl -s -X POST http://localhost:4000/api/auth/login \
  -H 'content-type: application/json' \
  -d '{"email":"admin@example.com","password":"ChangeMe123!"}'
```

Create a company:

```bash
curl -X POST http://localhost:4000/api/companies \
  -H "authorization: Bearer $TOKEN" -H 'content-type: application/json' \
  -d '{"name":"Acme Ltd","email":"ops@acme.example","status":"prospect"}'
```

Create a monitoring asset token:

```bash
curl -X POST http://localhost:4000/api/monitoring/assets/$ASSET_ID/tokens \
  -H "authorization: Bearer $TOKEN" -H 'content-type: application/json' \
  -d '{"name":"default agent"}'
```

Record a payment and recalculate invoice status:

```bash
curl -X POST http://localhost:4000/api/invoices/$INVOICE_ID/payments \
  -H "authorization: Bearer $TOKEN" -H 'content-type: application/json' \
  -d '{"amount":100,"payment_date":"2026-06-30","payment_method":"bank"}'
```

Create a document record:

```bash
curl -X POST http://localhost:4000/api/monitoring/documents \
  -H "authorization: Bearer $TOKEN" -H 'content-type: application/json' \
  -d '{"title":"Network diagram","document_type":"diagram","storage_url":"s3://bucket/network-diagram.pdf"}'
```

Create a patch job:

```bash
curl -X POST http://localhost:4000/api/monitoring/patch-jobs \
  -H "authorization: Bearer $TOKEN" -H 'content-type: application/json' \
  -d '{"asset_id":"'$ASSET_ID'","title":"Install security rollup","patch_identifier":"KB-EXAMPLE","severity":"high"}'
```

## Validation commands

```bash
cd business-portal/backend && npm test
cd business-portal/frontend && npm test
cd business-portal/agent && go test ./...
```

## MVP limitations

- No native accounting synchronization yet; integration placeholders are included.
- No Stripe payment processing/webhooks yet; finance records manual payments.
- Email ingestion, document management, mobile device, remote control, patch management, and SIEM correlation APIs are schema-backed MVP surfaces; external provider connectors and front-end workflows are still future work.
- ICMP depends on the host `ping` command and host/network permissions.
- Go service detection now uses OS commands plus process fallback, but deeper service-manager integrations can still be added per platform.
