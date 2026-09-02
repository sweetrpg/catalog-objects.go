package models

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// VolumeVersion model.
// This model represents one full-snapshot version of a volume's substantive fields, plus the
// submission/review audit trail that moves it through the version lifecycle (see VersionState).
type VolumeVersion struct {
	ID             string   `bson:"_id" json:"id" jsonapi:"primary,volume_version"`
	RecordID       string   `bson:"record_id" json:"record_id" jsonapi:"attr,record_id"`
	Version        int      `bson:"version" json:"version" jsonapi:"attr,version"`
	Title          string   `json:"title" jsonapi:"attr,title"`
	Description    string   `bson:"description" json:"description" jsonapi:"attr,description"`
	Notes          string   `json:"notes" jsonapi:"attr,notes"`
	Format         string   `bson:"format" json:"format" jsonapi:"attr,format"`
	CoverAssetId   string   `bson:"cover_asset_id" json:"cover_asset_id" jsonapi:"attr,coverAssetId"`
	SampleAssetIds []string `bson:"sample_asset_ids" json:"sample_asset_ids" jsonapi:"attr,sampleAssetIds"`
	// StagedCoverAssetId/StagedSampleAssetIds hold not-yet-promoted edit-session asset uploads
	// on a submitted version - see design.md's "Staged edit-session assets ride on the submitted
	// version as separate staged fields". CoverAssetId/SampleAssetIds above stay at their current
	// live values until promotion happens at accept time.
	StagedCoverAssetId   *string  `bson:"staged_cover_asset_id,omitempty" json:"staged_cover_asset_id,omitempty" jsonapi:"attr,staged_cover_asset_id,omitempty"`
	StagedSampleAssetIds []string `bson:"staged_sample_asset_ids,omitempty" json:"staged_sample_asset_ids,omitempty" jsonapi:"attr,staged_sample_asset_ids,omitempty"`
	SystemIds            []string `bson:"system_ids" json:"system_ids" jsonapi:"relation,system"`
	// SystemTitles is a denormalized snapshot of each referenced system's display name, keyed by
	// system ID, captured at write time so volume reads render the system name without a
	// game-systems-api call. Parallel to SystemIds; a nil/absent map is valid (renders the ID).
	SystemTitles map[string]string    `bson:"system_titles,omitempty" json:"system_titles,omitempty"`
	PublisherIds []string             `bson:"publisher_ids" json:"publisher_ids" jsonapi:"relation,publisher"`
	StudioIds    []string             `bson:"studio_ids" json:"studio_ids" jsonapi:"relation,studio"`
	LicenseIds   []string             `bson:"license_ids" json:"license_ids" jsonapi:"relation,license"`
	Properties   []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags         []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`

	State            VersionState `bson:"state" json:"state" jsonapi:"attr,state"`
	BaseVersion      *int         `bson:"base_version,omitempty" json:"base_version" jsonapi:"attr,base_version,omitempty"`
	SubmittedBy      string       `bson:"submitted_by" json:"submitted_by" jsonapi:"attr,submitted_by"`
	SubmittedAt      time.Time    `bson:"submitted_at" json:"submitted_at" jsonapi:"attr,submitted_at"`
	ReviewedBy       *string      `bson:"reviewed_by,omitempty" json:"reviewed_by" jsonapi:"attr,reviewed_by,omitempty"`
	ReviewedAt       *time.Time   `bson:"reviewed_at,omitempty" json:"reviewed_at" jsonapi:"attr,reviewed_at,omitempty"`
	ReviewNote       *string      `bson:"review_note,omitempty" json:"review_note" jsonapi:"attr,review_note,omitempty"`
	ResultingVersion *int         `bson:"resulting_version,omitempty" json:"resulting_version" jsonapi:"attr,resulting_version,omitempty"`
}
