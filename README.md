# Mortify-CRM-Suite

Modular CRM and business operations suite.

The main application lives in [`business-portal/`](business-portal/). See [`business-portal/README.md`](business-portal/README.md) for architecture, setup, validation commands, monitoring worker details, Go agent build instructions, and the latest operations-module API notes.

## Repository hygiene

This repo ignores generated dependency and build artifacts such as `node_modules/`, `dist/`, `build/`, `coverage/`, `.vite/`, logs, local `.env` overrides, and common editor/OS files. Commit source, migrations, package manifests, and dependency checksum/lock files; do not commit generated dependency directories.
