# AGENTS.md

This file provides guidance to Claude Code, Codex, GitHub Copilot, and other AI coding agents
working in this repository.

## About This Project

`catalog-objects.go` defines the persistence models (`models`) and API value objects (`vo`) for
the Catalog microservice's domain: licenses, contributions, persons, publishers, reviews,
studios, systems, and volumes. Pure data types - no business logic, no I/O.

## Dependencies

Depends on `model-core.go` (base `Auditable`/`Property`/`Tag` types embedded into every model
here). Depended on by `catalog-data.go` and `catalog-api`.

## Committing Code

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>
```

## Branches and Workflow

* `develop` - integration branch, default branch, target for all PRs.
* `master` - latest released state, nothing committed directly.
* `feature/*`, `fix/*` branched from `develop`; `hotfix/*` branched from `master`.

See `CONTRIBUTING.md` for the full workflow.

## Running Checks Locally

```bash
go build -v ./...
go vet ./...
go test -v ./...
```

## Releases

Merges to `develop` auto-tag a patch release via CI (`.github/workflows/go-ci.yml`). Use the
"Bump version" workflow (`.github/workflows/bump-version.yml`, manually dispatched) for a minor
or major bump instead.
