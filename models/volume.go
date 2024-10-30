package models

import modelcore "github.com/sweetrpg/model-core/models"

// Volume model.
// This model represents an RPG volume.
type Volume struct {
	ID           string               `bson:"_id" json:"id" jsonapi:"primary,volume"`
	Title        string               `json:"title" jsonapi:"attr,title"`
	Description  string               `bson:"description" json:"description" jsonapi:"attr,description"`
	Notes        string               `json:"notes" jsonapi:"attr,notes"`
	SystemIds    []string             `bson:"system_ids" json:"system_ids" jsonapi:"relation,system"`
	PublisherIds []string             `bson:"publisher_ids" json:"publisher_ids" jsonapi:"relation,publisher"`
	StudioIds    []string             `bson:"studio_ids" json:"studio_ids" jsonapi:"relation,studio"`
	LicenseIds   []string             `bson:"license_ids" json:"license_ids" jsonapi:"relation,license"`
	Properties   []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags         []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
