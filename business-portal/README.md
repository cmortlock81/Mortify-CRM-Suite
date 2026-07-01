# Mortify Business CRM Suite MVP

A contractor-ready SaaS-style business operations portal covering the full customer journey: Lead → Opportunity → Quote → Customer → Project → Helpdesk Ticket → Invoice → Payment. It also includes real IT monitoring for websites, SSL, DNS, TCP, HTTP/HTTPS checks, agent metrics, alerting, and ticket creation.

## Architecture

- **Frontend:** React + Vite, responsive SaaS sidebar layout, reusable tables/cards/badges.
- **Backend:** Node.js + Express REST API with JWT auth, PostgreSQL persistence, role checks, rate limits, and centralized errors.
- **Database:** PostgreSQL. SQL migrations in `backend/migrations` are the source of truth; Prisma schema is a reference only.
- **Worker:** `backend/src/workers/monitoringWorker.ts` runs out-of-band HTTP/HTTPS, DNS, TCP, SSL, keyword, and stale agent checks.
- **Agent:** Go + gopsutil collector for Windows/Linux host inventory and metrics, sending to `/api/agent/metrics` with an asset-scoped token.

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

## Database migrations and seed

```bash
cd business-portal/backend
npm install
npm run migrate
npm run seed
```

Migrations create users, CRM, sales, projects, helpdesk, finance, monitoring, indexes, and `global_search`.

## Run backend

```bash
cd business-portal/backend
npm run dev
```

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

## Build and run Go agent

```bash
cd business-portal/agent
go mod tidy
go build -o mortify-agent ./cmd/agent
cp config.example.yaml config.yaml
./mortify-agent -config config.yaml
```

Create a monitoring asset and token in the portal first. Put the raw token shown once into `config.yaml`. The backend stores only a bcrypt hash.

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

## MVP limitations

- No native accounting synchronization yet; integration placeholders are included.
- No Stripe payment processing/webhooks yet; finance records manual payments.
- No native email ingestion, document management, mobile apps, remote control, patch management, or SIEM correlation.
- ICMP depends on host permissions and is reported as unsupported by the worker; use TCP/HTTP checks for portable uptime monitoring.
- Service detection in the Go MVP is intentionally conservative and can be extended with OS-specific service APIs.
