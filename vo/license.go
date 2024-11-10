package vo

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core/vo"
)

// License value object.
// This value object is a serializable representation of the License model.
type LicenseVO struct {
	ID           string                 `json:"id" jsonapi:"primary,license"`
	Title        string                 `json:"title" jsonapi:"attr,title"`
	ShortTitle   string                 `json:"short_title" jsonapi:"attr,short_title"`
	Version      string                 `json:"version" jsonapi:"attr,version"`
	Deed         string                 `json:"deed" jsonapi:"attr,deed"`
	LegalCode    string                 `json:"legal_code" jsonapi:"attr,legal_code"`
	Website      url.URL                `json:"website" jsonapi:"attr,website"`
	Status       string                 `json:"status" jsonapi:"attr,status"`
	Availability string                 `json:"availability" jsonapi:"attr,availability"`
	Notes        string                 `json:"notes" jsonapi:"attr,notes"`
	Properties   []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags         []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
