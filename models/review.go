package models

import modelcore "github.com/sweetrpg/model-core.go/models"

// Review model.
// This model represents a review on an RPG resource.
type Review struct {
	ID       string          `bson:"_id" json:"id" jsonapi:"primary,review"`
	Title    string          `json:"title" jsonapi:"attr,title"`
	Body     string          `bson:"body" json:"body" jsonapi:"attr,body"`
	Language string          `json:"language" jsonapi:"attr,language"`
	VolumeId string          `bson:"volume_id" json:"volume_id" jsonapi:"relation,volume"`
	Notes    string          `json:"notes" jsonapi:"attr,notes"`
	Tags     []modelcore.Tag `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
