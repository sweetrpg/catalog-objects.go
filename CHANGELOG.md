
## 0.4.1 - 2026-08-18

### Fixed
- Serialize publisher/studio/license website as a string, not url.URL



## 0.4.0 - 2026-08-18

### Added
- Add meta+version models for publisher/studio/person/license/system



## 0.3.0 - 2026-08-14

### Added
- Add VolumeMeta/VolumeVersion models and shared VersionState


### Fixed
- Use camelCase json tags on VolumeVersionVO/VolumeMetaVO



## 0.2.0 - 2026-08-12

### Added
- Add CoverAssetId and SampleAssetIds fields


# Changelog

All notable changes to this project will be documented in this file.

## 0.1.0 - 2026-08-12

### Added
- Add Format field

## 0.0.197 - 2026-08-12

### Added
- CONTRIBUTING.md, CODE_OF_CONDUCT.md, AGENTS.md/CLAUDE.md repo scaffolding.

### Fixed
- `Review.VolumeId` and `System.GameSystem` were missing `bson` struct tags. Without them,
  MongoDB's default codec persists the fields as `volumeid`/`gamesystem` (no underscore),
  inconsistent with every sibling field's explicit snake_case tag (e.g. `Contribution.VolumeId`
  -> `volume_id`), silently breaking any filter/query built against the expected field name.
- `VolumeVO.Systems/Publishers/Studios/Licenses` were value slices (`[]T`); the `jsonapi`
  library's relationship marshaling panics on a non-pointer element. Changed to `[]*T`.

## 0.0.196 - 2026-07-23

### Documentation
- Update README (#42)

### Fixed
- Add missing bson tags on Review.VolumeId and System.GameSystem (#41)
