package models

import modelcore "github.com/sweetrpg/model-core.go/models"

// SystemVersion model.
// One full-snapshot version of a system's substantive fields, plus the shared submission/review
// audit trail (VersionLifecycle).
type SystemVersion struct {
	ID               string          `bson:"_id" json:"id" jsonapi:"primary,system_version"`
	RecordID         string          `bson:"record_id" json:"record_id" jsonapi:"attr,record_id"`
	Version          int             `bson:"version" json:"version" jsonapi:"attr,version"`
	GameSystem       string          `bson:"game_system" json:"game_system" jsonapi:"attr,game_system"`
	Edition          string          `bson:"edition" json:"edition" jsonapi:"attr,edition"`
	Notes            string          `json:"notes" jsonapi:"attr,notes"`
	Tags             []modelcore.Tag `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycle `bson:",inline"`
}
