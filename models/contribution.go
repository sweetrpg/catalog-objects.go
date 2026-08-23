package models

import (
	modelcore "github.com/sweetrpg/model-core.go/models"
)

// Contribution model.
// This model represents the association of a person to a volume.
type Contribution struct {
	ID       string `bson:"_id" json:"id" jsonapi:"primary,contribution"`
	PersonId string `bson:"person_id" json:"person_id" jsonapi:"relation,person"`
	VolumeId string `bson:"volume_id" json:"volume_id" jsonapi:"relation,volume"`
	Role     string `bson:"role" json:"role" jsonapi:"attr,role"`
	Notes    string `json:"notes" jsonapi:"attr,notes"`
	modelcore.Auditable
}
