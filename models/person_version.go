package models

import modelcore "github.com/sweetrpg/model-core.go/models"

// PersonVersion model.
// One full-snapshot version of a person's substantive fields, plus the shared submission/review
// audit trail (VersionLifecycle).
type PersonVersion struct {
	ID               string               `bson:"_id" json:"id" jsonapi:"primary,person_version"`
	RecordID         string               `bson:"record_id" json:"record_id" jsonapi:"attr,record_id"`
	Version          int                  `bson:"version" json:"version" jsonapi:"attr,version"`
	Name             string               `json:"name" jsonapi:"attr,name"`
	Notes            string               `json:"notes" jsonapi:"attr,notes"`
	Properties       []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags             []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycle `bson:",inline"`
}
