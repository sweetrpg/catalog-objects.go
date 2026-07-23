# Changelog

All notable changes to this project will be documented in this file.

## 0.0.196 - 2026-07-23

### Documentation
- Update README (#42)

### Fixed
- Add missing bson tags on Review.VolumeId and System.GameSystem (#41)

## [Unreleased]

### Added

- CONTRIBUTING.md, CODE_OF_CONDUCT.md, AGENTS.md/CLAUDE.md repo scaffolding.

### Fixed

- `Review.VolumeId` and `System.GameSystem` were missing `bson` struct tags. Without them,
  MongoDB's default codec persists the fields as `volumeid`/`gamesystem` (no underscore),
  inconsistent with every sibling field's explicit snake_case tag (e.g. `Contribution.VolumeId`
  -> `volume_id`), silently breaking any filter/query built against the expected field name.
