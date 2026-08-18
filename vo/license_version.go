package vo

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// LicenseVersionVO value object - see models.LicenseVersion. Note the internal model field
// LicenseVer/version_label maps back onto this VO's Version/"version" - the public read shape is
// unaffected by the internal rename needed to avoid colliding with the version record's own
// Version int field.
type LicenseVersionVO struct {
	ID           string                 `json:"id" jsonapi:"primary,license_version"`
	RecordID     string                 `json:"recordId" jsonapi:"attr,record_id"`
	Version      int                    `json:"version" jsonapi:"attr,version"`
	Title        string                 `json:"title" jsonapi:"attr,title"`
	ShortTitle   string                 `json:"shortTitle" jsonapi:"attr,short_title"`
	LicenseVer   string                 `json:"licenseVersion" jsonapi:"attr,license_version"`
	Deed         string                 `json:"deed" jsonapi:"attr,deed"`
	LegalCode    string                 `json:"legalCode" jsonapi:"attr,legal_code"`
	Website      url.URL                `json:"website" jsonapi:"attr,website"`
	Status       string                 `json:"status" jsonapi:"attr,status"`
	Availability string                 `json:"availability" jsonapi:"attr,availability"`
	Notes        string                 `json:"notes" jsonapi:"attr,notes"`
	Properties   []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags         []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycleVO
}
