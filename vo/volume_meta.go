package vo

import "time"

// VolumeMeta value object.
// This value object is a serializable representation of the VolumeMeta model.
type VolumeMetaVO struct {
	ID             string     `json:"id" jsonapi:"primary,volume_meta"`
	CurrentVersion int        `json:"current_version" jsonapi:"attr,current_version"`
	CreatedAt      time.Time  `json:"created_at" jsonapi:"attr,created_at"`
	CreatedBy      string     `json:"created_by" jsonapi:"attr,created_by"`
	DeletedAt      *time.Time `json:"deleted_at" jsonapi:"attr,deleted_at,omitempty"`
	DeletedBy      *string    `json:"deleted_by" jsonapi:"attr,deleted_by,omitempty"`
}
