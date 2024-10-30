package vo

import (
	modelcore "github.com/sweetrpg/model-core/vo"
)

type ReviewVO struct {
	ID       string            `json:"id" jsonapi:"primary,review"`
	Title    string            `json:"title" jsonapi:"attr,title"`
	Body     string            `json:"body" jsonapi:"attr,body"`
	Language string            `json:"language" jsonapi:"attr,language"`
	Volume   *VolumeVO         `json:"volume" jsonapi:"relation,volume"`
	Notes    string            `json:"notes" jsonapi:"attr,notes"`
	Tags     []modelcore.TagVO `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
