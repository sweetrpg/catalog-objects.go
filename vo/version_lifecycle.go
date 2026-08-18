package vo

import "time"

// VersionLifecycleVO mirrors models.VersionLifecycle - see that type's doc comment.
type VersionLifecycleVO struct {
	State            VersionState `json:"state" jsonapi:"attr,state"`
	BaseVersion      *int         `json:"baseVersion,omitempty" jsonapi:"attr,base_version,omitempty"`
	SubmittedBy      string       `json:"submittedBy" jsonapi:"attr,submitted_by"`
	SubmittedAt      time.Time    `json:"submittedAt" jsonapi:"attr,submitted_at"`
	ReviewedBy       *string      `json:"reviewedBy,omitempty" jsonapi:"attr,reviewed_by,omitempty"`
	ReviewedAt       *time.Time   `json:"reviewedAt,omitempty" jsonapi:"attr,reviewed_at,omitempty"`
	ReviewNote       *string      `json:"reviewNote,omitempty" jsonapi:"attr,review_note,omitempty"`
	ResultingVersion *int         `json:"resultingVersion,omitempty" jsonapi:"attr,resulting_version,omitempty"`
}
