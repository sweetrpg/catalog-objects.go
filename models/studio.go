package models

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// Studio model.
// This model represents a studio that can produce an RPG resource.
type Studio struct {
	ID         string               `bson:"_id" json:"id" jsonapi:"primary,studio"`
	Name       string               `json:"name" jsonapi:"attr,name"`
	Website    url.URL              `json:"website" jsonapi:"attr,website"`
	Notes      string               `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
