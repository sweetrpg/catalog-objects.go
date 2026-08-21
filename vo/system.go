package vo

import (
	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// System value object.
// This value object is a serializable representation of the System model.
type SystemVO struct {
	ID string `json:"id" jsonapi:"primary,system"`
	// GameSystem holds the game system's display name (e.g. "Dungeons & Dragons") - tagged
	// "name" on the wire, not "game_system", so catalog-api-client.swift's generic
	// NamedAttributes decoder (name ?? title ?? "Untitled") picks it up the same way it
	// already does for Publisher/Studio/Person/License.
	GameSystem string            `json:"game_system" jsonapi:"attr,name"`
	Edition    string            `json:"edition" jsonapi:"attr,edition"`
	Notes      string            `json:"notes" jsonapi:"attr,notes"`
	Tags       []modelcore.TagVO `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
