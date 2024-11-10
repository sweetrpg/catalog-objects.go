package models

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core/models"
)

// Publisher model.
// This model represents a publisher involved in the production of an RPG resource.
type Publisher struct {
	ID         string               `bson:"_id" json:"id" jsonapi:"primary,publisher"`
	Name       string               `json:"name" jsonapi:"attr,name"`
	Address    string               `json:"address" jsonapi:"attr,address"`
	Website    url.URL              `json:"website" jsonapi:"attr,website"`
	Notes      string               `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
