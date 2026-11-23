package vo

import (
	"time"

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
	// Audit fields are direct fields, not an embedded modelcore.AuditableVO: google/jsonapi
	// v1.0.0 skips untagged struct fields, so an embedded struct never reaches the wire.
	// Times are RFC 3339 strings, which is what catalog-web's decoder expects.
	CreatedAt time.Time  `json:"created_at" jsonapi:"attr,created_at,rfc3339"`
	CreatedBy string     `json:"created_by" jsonapi:"attr,created_by"`
	UpdatedAt time.Time  `json:"updated_at" jsonapi:"attr,updated_at,rfc3339"`
	UpdatedBy string     `json:"updated_by" jsonapi:"attr,updated_by"`
	DeletedAt *time.Time `json:"deleted_at" jsonapi:"attr,deleted_at,omitempty,rfc3339"`
	DeletedBy *string    `json:"deleted_by" jsonapi:"attr,deleted_by,omitempty"`
}
