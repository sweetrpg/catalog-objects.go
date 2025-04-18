package models

import modelcore "github.com/sweetrpg/model-core.go/models"

// System model.
// This model represents a game system that RPG resources can be associated with.
type System struct {
	ID         string          `bson:"_id" json:"id" jsonapi:"primary,system"`
	GameSystem string          `json:"game_system" jsonapi:"attr,game_system"`
	Edition    string          `bson:"edition" json:"edition" jsonapi:"attr,edition"`
	Notes      string          `json:"notes" jsonapi:"attr,notes"`
	Tags       []modelcore.Tag `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
