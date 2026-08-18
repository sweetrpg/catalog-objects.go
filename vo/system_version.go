package vo

import modelcore "github.com/sweetrpg/model-core.go/vo"

// SystemVersionVO value object - see models.SystemVersion.
type SystemVersionVO struct {
	ID         string            `json:"id" jsonapi:"primary,system_version"`
	RecordID   string            `json:"recordId" jsonapi:"attr,record_id"`
	Version    int               `json:"version" jsonapi:"attr,version"`
	GameSystem string            `json:"gameSystem" jsonapi:"attr,game_system"`
	Edition    string            `json:"edition" jsonapi:"attr,edition"`
	Notes      string            `json:"notes" jsonapi:"attr,notes"`
	Tags       []modelcore.TagVO `json:"tags" jsonapi:"attr,tags"`
	VersionLifecycleVO
}
