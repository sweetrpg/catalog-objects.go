# catalog-objects.go

[![CI](https://github.com/sweetrpg/catalog-objects.go/actions/workflows/ci.yaml/badge.svg)](https://github.com/sweetrpg/catalog-objects.go/actions/workflows/ci.yaml)
[![License](https://img.shields.io/github/license/sweetrpg/catalog-objects.go.svg)](https://img.shields.io/github/license/sweetrpg/catalog-objects.go.svg)
[![Issues](https://img.shields.io/github/issues/sweetrpg/catalog-objects.go.svg)](https://img.shields.io/github/issues/sweetrpg/catalog-objects.go.svg)
[![PRs](https://img.shields.io/github/issues-pr/sweetrpg/catalog-objects.go.svg)](https://img.shields.io/github/issues-pr/sweetrpg/catalog-objects.go.svg)
[![Dependabot](https://badgen.net/github/dependabot/sweetrpg/catalog-objects.go)](https://badgen.net/github/dependabot/sweetrpg/catalog-objects.go)

Persistence models and API value objects for the Catalog microservice's domain: licenses,
contributions, persons, publishers, reviews, studios, systems, and volumes. Pure data types -
no business logic, no I/O.

## Install

```bash
go get github.com/sweetrpg/catalog-objects.go
```

## Packages

- `models` - persistence-layer structs (`License`, `Contribution`, `Person`, `Publisher`,
  `Review`, `Studio`, `System`, `Volume`), each embedding `model-core.go`'s `Auditable`
- `vo` - the matching API-facing value objects, each embedding `model-core.go`'s `AuditableVO`

## Documentation

Package documentation: [pkg.go.dev/github.com/sweetrpg/catalog-objects.go](https://pkg.go.dev/github.com/sweetrpg/catalog-objects.go).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the development workflow and
[RELEASE.md](RELEASE.md) for how versions get cut.
