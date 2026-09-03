package vo

import "time"

// VolumeMeta value object.
// This value object is a serializable representation of the VolumeMeta model.
type VolumeMetaVO struct {
	ID             string     `json:"id" jsonapi:"primary,volume_meta"`
	CurrentVersion int        `json:"currentVersion" jsonapi:"attr,current_version"`
	CreatedAt      time.Time  `json:"createdAt" jsonapi:"attr,created_at"`
	CreatedBy      string     `json:"createdBy" jsonapi:"attr,created_by"`
	UpdatedAt      time.Time  `json:"updatedAt" jsonapi:"attr,updated_at"`
	UpdatedBy      string     `json:"updatedBy" jsonapi:"attr,updated_by"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty" jsonapi:"attr,deleted_at,omitempty"`
	DeletedBy      *string    `json:"deletedBy,omitempty" jsonapi:"attr,deleted_by,omitempty"`
}
