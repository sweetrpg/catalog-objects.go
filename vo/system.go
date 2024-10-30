package vo

import (
	modelcore "github.com/sweetrpg/model-core/vo"
)

type SystemVO struct {
	ID         string            `json:"id" jsonapi:"primary,system"`
	GameSystem string            `json:"game_system" jsonapi:"attr,game_system"`
	Edition    string            `json:"edition" jsonapi:"attr,edition"`
	Notes      string            `json:"notes" jsonapi:"attr,notes"`
	Tags       []modelcore.TagVO `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
