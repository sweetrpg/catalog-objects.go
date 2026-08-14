package vo

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// VolumeVersion value object.
// This value object is a serializable representation of the VolumeVersion model.
type VolumeVersionVO struct {
	ID             string                 `json:"id" jsonapi:"primary,volume_version"`
	RecordID       string                 `json:"recordId" jsonapi:"attr,record_id"`
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
	BaseVersion      *int         `json:"baseVersion,omitempty" jsonapi:"attr,base_version,omitempty"`
	SubmittedBy      string       `json:"submittedBy" jsonapi:"attr,submitted_by"`
	SubmittedAt      time.Time    `json:"submittedAt" jsonapi:"attr,submitted_at"`
	ReviewedBy       *string      `json:"reviewedBy,omitempty" jsonapi:"attr,reviewed_by,omitempty"`
	ReviewedAt       *time.Time   `json:"reviewedAt,omitempty" jsonapi:"attr,reviewed_at,omitempty"`
	ReviewNote       *string      `json:"reviewNote,omitempty" jsonapi:"attr,review_note,omitempty"`
	ResultingVersion *int         `json:"resultingVersion,omitempty" jsonapi:"attr,resulting_version,omitempty"`
}
