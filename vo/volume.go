package vo

import (
	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// Volume value object.
// This value object is a serializable representation of the Volume model.
//
// Systems/Publishers/Studios/Licenses are pointer slices, not value slices - the jsonapi
// library's relationship marshaling (visitModelNode, called recursively per relationship
// element) requires each element to be a pointer; passed a value type, reflect.Value.IsNil
// panics rather than returning an error. This never surfaced before because no Volume's
// relationships were ever actually populated through the write path until
// durable-volume-editing's publisher/studio linking exercised it for the first time.
type VolumeVO struct {
	ID          string                 `json:"id" jsonapi:"primary,volume"`
	Title       string                 `json:"title" jsonapi:"attr,title"`
	Description string                 `json:"description" jsonapi:"attr,description"`
	Notes       string                 `json:"notes" jsonapi:"attr,notes"`
	Format      string                 `json:"format" jsonapi:"attr,format"`
	Systems     []*SystemVO            `json:"systems" jsonapi:"relation,system"`
	Publishers  []*PublisherVO         `json:"publishers" jsonapi:"relation,publisher"`
	Studios     []*StudioVO            `json:"studios" jsonapi:"relation,studio"`
	Licenses    []*LicenseVO           `json:"licenses" jsonapi:"relation,license"`
	Properties  []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags        []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	modelcore.AuditableVO
}
