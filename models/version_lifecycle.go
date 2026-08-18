package models

import "time"

// VersionLifecycle is the submission/review audit trail shared by every versioned entity type's
// version record - state, provenance, and review outcome. Embedded into each type's own Version
// struct (PublisherVersion, StudioVersion, ...) alongside that type's substantive fields, the
// same way modelcore.Auditable is embedded into every entity model. VolumeVersion predates this
// type and carries the same fields inline rather than via embedding - not worth reshaping
// already-shipped, tested code to match.
type VersionLifecycle struct {
	State            VersionState `bson:"state" json:"state" jsonapi:"attr,state"`
	BaseVersion      *int         `bson:"base_version,omitempty" json:"base_version" jsonapi:"attr,base_version,omitempty"`
	SubmittedBy      string       `bson:"submitted_by" json:"submitted_by" jsonapi:"attr,submitted_by"`
	SubmittedAt      time.Time    `bson:"submitted_at" json:"submitted_at" jsonapi:"attr,submitted_at"`
	ReviewedBy       *string      `bson:"reviewed_by,omitempty" json:"reviewed_by" jsonapi:"attr,reviewed_by,omitempty"`
	ReviewedAt       *time.Time   `bson:"reviewed_at,omitempty" json:"reviewed_at" jsonapi:"attr,reviewed_at,omitempty"`
	ReviewNote       *string      `bson:"review_note,omitempty" json:"review_note" jsonapi:"attr,review_note,omitempty"`
	ResultingVersion *int         `bson:"resulting_version,omitempty" json:"resulting_version" jsonapi:"attr,resulting_version,omitempty"`
}
