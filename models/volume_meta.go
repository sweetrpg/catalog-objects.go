package models

import "time"

// VolumeMeta model.
// This model represents the stable identity of a volume record: its id, the version currently
// live, and creation/deletion audit. The record's substantive fields live on VolumeVersion.
type VolumeMeta struct {
	ID             string     `bson:"_id" json:"id" jsonapi:"primary,volume_meta"`
	CurrentVersion int        `bson:"current_version" json:"current_version" jsonapi:"attr,current_version"`
	CreatedAt      time.Time  `bson:"created_at" json:"created_at" jsonapi:"attr,created_at"`
	CreatedBy      string     `bson:"created_by" json:"created_by" jsonapi:"attr,created_by"`
	UpdatedAt      time.Time  `bson:"updated_at" json:"updated_at" jsonapi:"attr,updated_at"`
	UpdatedBy      string     `bson:"updated_by" json:"updated_by" jsonapi:"attr,updated_by"`
	DeletedAt      *time.Time `bson:"deleted_at,omitempty" json:"deleted_at" jsonapi:"attr,deleted_at,omitempty"`
	DeletedBy      *string    `bson:"deleted_by,omitempty" json:"deleted_by" jsonapi:"attr,deleted_by,omitempty"`
}
