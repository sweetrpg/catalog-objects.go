package models

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// LicenseVersion model.
// One full-snapshot version of a license's substantive fields, plus the shared submission/review
// audit trail (VersionLifecycle).
type LicenseVersion struct {
	ID               string               `bson:"_id" json:"id" jsonapi:"primary,license_version"`
	RecordID         string               `bson:"record_id" json:"record_id" jsonapi:"attr,record_id"`
	Version          int                  `bson:"version" json:"version" jsonapi:"attr,version"`
	Title            string               `json:"title" jsonapi:"attr,title"`
	ShortTitle       string               `bson:"short_title" json:"short_title" jsonapi:"attr,short_title"`
	LicenseVer       string               `bson:"version_label" json:"version_label" jsonapi:"attr,version_label"`
	Deed             string               `json:"deed" jsonapi:"attr,deed"`
	LegalCode        string               `bson:"legal_code" json:"legal_code" jsonapi:"attr,legal_code"`
	Website          url.URL              `json:"website" jsonapi:"attr,website"`
	Status           string               `json:"status" jsonapi:"attr,status"`
	Availability     string               `json:"availability" jsonapi:"attr,availability"`
	Notes            string               `json:"notes" jsonapi:"attr,notes"`
	Properties       []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags             []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycle `bson:",inline"`
}
