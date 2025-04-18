package models

import (
	"net/url"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// License model.
// This model represents the license information for any given RPG resource.
type License struct {
	ID           string               `bson:"_id" json:"id" jsonapi:"primary,license"`
	Title        string               `json:"title" jsonapi:"attr,title"`
	ShortTitle   string               `bson:"short_title" json:"short_title" jsonapi:"attr,short_title"`
	Version      string               `json:"version" jsonapi:"attr,version"`
	Deed         string               `json:"deed" jsonapi:"attr,deed"`
	LegalCode    string               `bson:"legal_code" json:"legal_code" jsonapi:"attr,legal_code"`
	Website      url.URL              `json:"website" jsonapi:"attr,website"`
	Status       string               `json:"status" jsonapi:"attr,status"`
	Availability string               `json:"availability" jsonapi:"attr,availability"`
	Notes        string               `json:"notes" jsonapi:"attr,notes"`
	Properties   []modelcore.Property `json:"properties" jsonapi:"attr,properties"`
	Tags         []modelcore.Tag      `json:"tags" jsonapi:"attr,tags"`
	modelcore.Auditable
}
