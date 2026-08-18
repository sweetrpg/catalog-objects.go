package models

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// PublisherVersion model.
// One full-snapshot version of a publisher's substantive fields, plus the shared submission/
// review audit trail (VersionLifecycle).
type PublisherVersion struct {
	ID               string               `bson:"_id" json:"id" jsonapi:"primary,publisher_version"`
	RecordID         string               `bson:"record_id" json:"record_id" jsonapi:"attr,record_id"`
	Version          int                  `bson:"version" json:"version" jsonapi:"attr,version"`
	Name             string               `json:"name" jsonapi:"attr,name"`
	Address          string               `json:"address" jsonapi:"attr,address"`
	Website          url.URL              `json:"website" jsonapi:"attr,website"`
	Notes            string               `json:"notes" jsonapi:"attr,notes"`
	Properties       []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags             []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycle `bson:",inline"`
}
