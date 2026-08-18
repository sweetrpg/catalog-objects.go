package vo

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// PublisherVersionVO value object - see models.PublisherVersion.
type PublisherVersionVO struct {
	ID         string                 `json:"id" jsonapi:"primary,publisher_version"`
	RecordID   string                 `json:"recordId" jsonapi:"attr,record_id"`
	Version    int                    `json:"version" jsonapi:"attr,version"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Address    string                 `json:"address" jsonapi:"attr,address"`
	Website    url.URL                `json:"website" jsonapi:"attr,website"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycleVO
}
