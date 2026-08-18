package models

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// StudioVersion model.
// One full-snapshot version of a studio's substantive fields, plus the shared submission/review
// audit trail (VersionLifecycle).
type StudioVersion struct {
	ID               string               `bson:"_id" json:"id" jsonapi:"primary,studio_version"`
	RecordID         string               `bson:"record_id" json:"record_id" jsonapi:"attr,record_id"`
	Version          int                  `bson:"version" json:"version" jsonapi:"attr,version"`
	Name             string               `json:"name" jsonapi:"attr,name"`
	Website          url.URL              `json:"website" jsonapi:"attr,website"`
	Notes            string               `json:"notes" jsonapi:"attr,notes"`
	Properties       []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags             []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycle `bson:",inline"`
}
