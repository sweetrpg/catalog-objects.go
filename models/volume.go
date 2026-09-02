package models

import modelcore "github.com/sweetrpg/model-core.go/models"

// Volume model.
// This model represents an RPG volume.
type Volume struct {
	ID             string   `bson:"_id" json:"id" jsonapi:"primary,volume"`
	Title          string   `json:"title" jsonapi:"attr,title"`
	Description    string   `bson:"description" json:"description" jsonapi:"attr,description"`
	Notes          string   `json:"notes" jsonapi:"attr,notes"`
	Format         string   `bson:"format" json:"format" jsonapi:"attr,format"`
	CoverAssetId   string   `bson:"cover_asset_id" json:"cover_asset_id" jsonapi:"attr,coverAssetId"`
	SampleAssetIds []string `bson:"sample_asset_ids" json:"sample_asset_ids" jsonapi:"attr,sampleAssetIds"`
	SystemIds      []string `bson:"system_ids" json:"system_ids" jsonapi:"relation,system"`
	// SystemTitles: denormalized system display names keyed by system ID, parallel to SystemIds.
	// Captured at write time so reads render the name without a game-systems-api call; nil is
	// valid (renders the ID).
	SystemTitles map[string]string    `bson:"system_titles,omitempty" json:"system_titles,omitempty"`
	PublisherIds []string             `bson:"publisher_ids" json:"publisher_ids" jsonapi:"relation,publisher"`
	StudioIds    []string             `bson:"studio_ids" json:"studio_ids" jsonapi:"relation,studio"`
	LicenseIds   []string             `bson:"license_ids" json:"license_ids" jsonapi:"relation,license"`
	Properties   []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags         []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
