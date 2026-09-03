package models

import "time"

// EntityMeta is the stable-identity record shared by every versioned entity type introduced by
// this change (publisher, studio, person, license, system): its id, the version currently live,
// and creation/deletion audit. Its shape never varies by type, so it's reused directly rather
// than duplicated per type - unlike VolumeMeta, which predates this shared type and stays as its
// own struct rather than being reshaped after shipping.
type EntityMeta struct {
	ID             string     `bson:"_id" json:"id"`
	CurrentVersion int        `bson:"current_version" json:"current_version"`
	CreatedAt      time.Time  `bson:"created_at" json:"created_at"`
	CreatedBy      string     `bson:"created_by" json:"created_by"`
	UpdatedAt      time.Time  `bson:"updated_at" json:"updated_at"`
	UpdatedBy      string     `bson:"updated_by" json:"updated_by"`
	DeletedAt      *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	DeletedBy      *string    `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
}
