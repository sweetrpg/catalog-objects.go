package vo

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// VolumeVersion value object.
// This value object is a serializable representation of the VolumeVersion model.
type VolumeVersionVO struct {
	ID             string                 `json:"id" jsonapi:"primary,volume_version"`
	RecordID       string                 `json:"record_id" jsonapi:"attr,record_id"`
	Version        int                    `json:"version" jsonapi:"attr,version"`
	Title          string                 `json:"title" jsonapi:"attr,title"`
	Description    string                 `json:"description" jsonapi:"attr,description"`
	Notes          string                 `json:"notes" jsonapi:"attr,notes"`
	Format         string                 `json:"format" jsonapi:"attr,format"`
	CoverAssetId   string                 `json:"coverAssetId" jsonapi:"attr,coverAssetId"`
	SampleAssetIds []string               `json:"sampleAssetIds" jsonapi:"attr,sampleAssetIds"`
	Systems        []*SystemVO            `json:"systems" jsonapi:"relation,system"`
	Publishers     []*PublisherVO         `json:"publishers" jsonapi:"relation,publisher"`
	Studios        []*StudioVO            `json:"studios" jsonapi:"relation,studio"`
	Licenses       []*LicenseVO           `json:"licenses" jsonapi:"relation,license"`
	Properties     []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags           []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`

	State            VersionState `json:"state" jsonapi:"attr,state"`
	BaseVersion      *int         `json:"base_version" jsonapi:"attr,base_version,omitempty"`
	SubmittedBy      string       `json:"submitted_by" jsonapi:"attr,submitted_by"`
	SubmittedAt      time.Time    `json:"submitted_at" jsonapi:"attr,submitted_at"`
	ReviewedBy       *string      `json:"reviewed_by" jsonapi:"attr,reviewed_by,omitempty"`
	ReviewedAt       *time.Time   `json:"reviewed_at" jsonapi:"attr,reviewed_at,omitempty"`
	ReviewNote       *string      `json:"review_note" jsonapi:"attr,review_note,omitempty"`
	ResultingVersion *int         `json:"resulting_version" jsonapi:"attr,resulting_version,omitempty"`
}
