package models

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core/models"
)

// Studio model.
// This model represents a studio that can produce an RPG resource.
type Studio struct {
	ID         string               `bson:"_id" json:"id" jsonapi:"primary,studio"`
	Name       string               `json:"name" jsonapi:"attr,name"`
	URL        url.URL              `bson:"url" json:"url" jsonapi:"attr,url"`
	Notes      string               `json:"notes" jsonapi:"attr,notes"`
	Properties []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags       []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
