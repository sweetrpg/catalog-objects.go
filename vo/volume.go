package vo

import (
	"time"

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
	ID             string      `json:"id" jsonapi:"primary,volume"`
	Title          string      `json:"title" jsonapi:"attr,title"`
	Description    string      `json:"description" jsonapi:"attr,description"`
	Notes          string      `json:"notes" jsonapi:"attr,notes"`
	Format         string      `json:"format" jsonapi:"attr,format"`
	CoverAssetId   string      `json:"coverAssetId" jsonapi:"attr,coverAssetId"`
	SampleAssetIds []string    `json:"sampleAssetIds" jsonapi:"attr,sampleAssetIds"`
	Systems        []*SystemVO `json:"systems" jsonapi:"relation,system"`
	// SystemTitles: denormalized system display names keyed by system ID. On a read it carries
	// the stored snapshot; on a write it is an advisory hint (catalog-api overwrites it from an
	// authoritative game-systems-api lookup when reachable).
	SystemTitles map[string]string      `json:"systemTitles,omitempty" jsonapi:"attr,systemTitles,omitempty"`
	Publishers   []*PublisherVO         `json:"publishers" jsonapi:"relation,publisher"`
	Studios      []*StudioVO            `json:"studios" jsonapi:"relation,studio"`
	Licenses     []*LicenseVO           `json:"licenses" jsonapi:"relation,license"`
	Properties   []modelcore.PropertyVO `json:"properties" jsonapi:"attr,properties"`
	Tags         []modelcore.TagVO      `json:"tags" jsonapi:"attr,tags"`
	// Audit fields are direct fields, not an embedded modelcore.AuditableVO: google/jsonapi
	// v1.0.0 skips untagged struct fields, so an embedded struct never reaches the wire.
	// Times are RFC 3339 strings, which is what catalog-web's decoder expects.
	CreatedAt time.Time  `json:"created_at" jsonapi:"attr,created_at,rfc3339"`
	CreatedBy string     `json:"created_by" jsonapi:"attr,created_by"`
	UpdatedAt time.Time  `json:"updated_at" jsonapi:"attr,updated_at,rfc3339"`
	UpdatedBy string     `json:"updated_by" jsonapi:"attr,updated_by"`
	DeletedAt *time.Time `json:"deleted_at" jsonapi:"attr,deleted_at,omitempty,rfc3339"`
	DeletedBy *string    `json:"deleted_by" jsonapi:"attr,deleted_by,omitempty"`
}
