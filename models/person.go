package models

import modelcore "github.com/sweetrpg/model-core/models"

// Person model.
// This model represents a person involved in the production of an RPG resource.
type Person struct {
	ID         string               `bson:"_id" json:"id" jsonapi:"primary,person"`
	Name       string               `json:"name" jsonapi:"attr,name"`
	Notes      string               `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
