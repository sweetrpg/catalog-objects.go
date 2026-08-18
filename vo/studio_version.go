package vo

import (
	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// StudioVersionVO value object - see models.StudioVersion. Website is a plain string, not
// url.URL - see PublisherVO's comment.
type StudioVersionVO struct {
	ID         string                 `json:"id" jsonapi:"primary,studio_version"`
	RecordID   string                 `json:"recordId" jsonapi:"attr,record_id"`
	Version    int                    `json:"version" jsonapi:"attr,version"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Website    string                 `json:"website" jsonapi:"attr,website"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycleVO
}
