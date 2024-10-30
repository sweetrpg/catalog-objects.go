package vo

import (
	modelcore "github.com/sweetrpg/model-core/vo"
)

// System value object.
// This value object is a serializable representation of the System model.
type SystemVO struct {
	ID         string            `json:"id" jsonapi:"primary,system"`
	GameSystem string            `json:"game_system" jsonapi:"attr,game_system"`
	Edition    string            `json:"edition" jsonapi:"attr,edition"`
	Notes      string            `json:"notes" jsonapi:"attr,notes"`
	Tags       []modelcore.TagVO `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
