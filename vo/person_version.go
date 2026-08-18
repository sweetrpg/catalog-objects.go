package vo

import modelcore "github.com/sweetrpg/model-core.go/vo"

// PersonVersionVO value object - see models.PersonVersion.
type PersonVersionVO struct {
	ID         string                 `json:"id" jsonapi:"primary,person_version"`
	RecordID   string                 `json:"recordId" jsonapi:"attr,record_id"`
	Version    int                    `json:"version" jsonapi:"attr,version"`
	Name       string                 `json:"name" jsonapi:"attr,name"`
	Notes      string                 `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycleVO
}
