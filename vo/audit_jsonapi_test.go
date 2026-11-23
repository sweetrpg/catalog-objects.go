package vo

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/google/jsonapi"
	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// google/jsonapi v1.0.0 skips untagged struct fields, so an embedded audit struct never
// reached the wire. Every VO must expose the audit fields as top-level attributes.
func auditSample() (modelcore.AuditableVO, time.Time) {
	deletedAt := time.Date(2026, 11, 20, 8, 30, 0, 0, time.UTC)
	deletedBy := "user-9"
	return modelcore.AuditableVO{
		CreatedAt: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC),
		CreatedBy: "user-1",
		UpdatedAt: time.Date(2026, 6, 7, 8, 9, 10, 0, time.UTC),
		UpdatedBy: "user-2",
		DeletedAt: &deletedAt,
		DeletedBy: &deletedBy,
	}, deletedAt
}

func TestVOsMarshalAuditFieldsAsJSONAPIAttributes(t *testing.T) {
	a, _ := auditSample()
	tests := []struct {
		name string
		vo   any
	}{
		{"volume", &VolumeVO{ID: "1", Title: "t"}},
		{"publisher", &PublisherVO{ID: "1"}},
		{"studio", &StudioVO{ID: "1"}},
		{"person", &PersonVO{ID: "1"}},
		{"license", &LicenseVO{ID: "1"}},
		{"system", &SystemVO{ID: "1"}},
		{"review", &ReviewVO{ID: "1"}},
		{"contribution", &ContributionVO{ID: "1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setAudit(t, tt.vo, a)
			var buf bytes.Buffer
			if err := jsonapi.MarshalPayload(&buf, tt.vo); err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var doc struct {
				Data struct {
					Attributes map[string]any `json:"attributes"`
				} `json:"data"`
			}
			if err := json.Unmarshal(buf.Bytes(), &doc); err != nil {
				t.Fatal(err)
			}
			want := map[string]any{
				"created_at": "2026-01-02T03:04:05Z", "created_by": "user-1",
				"updated_at": "2026-06-07T08:09:10Z", "updated_by": "user-2",
				"deleted_at": "2026-11-20T08:30:00Z", "deleted_by": "user-9",
			}
			for k, v := range want {
				if got := doc.Data.Attributes[k]; got != v {
					t.Errorf("attributes[%q] = %v, want %v\npayload: %s", k, got, v, buf.String())
				}
			}
		})
	}
}

func TestVolumeAuditFieldsRoundTripThroughJSONAPI(t *testing.T) {
	a, deletedAt := auditSample()
	in := &VolumeVO{ID: "1", Title: "t"}
	setAudit(t, in, a)
	var buf bytes.Buffer
	if err := jsonapi.MarshalPayload(&buf, in); err != nil {
		t.Fatal(err)
	}
	var out VolumeVO
	if err := jsonapi.UnmarshalPayload(&buf, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out.DeletedAt == nil || !out.DeletedAt.Equal(deletedAt) || out.CreatedBy != "user-1" {
		t.Errorf("round trip lost audit fields: %+v", out)
	}
}

func TestLiveVolumeOmitsDeletionAttributes(t *testing.T) {
	a, _ := auditSample()
	a.DeletedAt, a.DeletedBy = nil, nil
	in := &VolumeVO{ID: "1", Title: "t"}
	setAudit(t, in, a)
	var buf bytes.Buffer
	if err := jsonapi.MarshalPayload(&buf, in); err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Data struct {
			Attributes map[string]any `json:"attributes"`
		} `json:"data"`
	}
	_ = json.Unmarshal(buf.Bytes(), &doc)
	if _, ok := doc.Data.Attributes["deleted_at"]; ok {
		t.Errorf("live record carries deleted_at: %s", buf.String())
	}
}

// setAudit copies a's fields onto v by name, so the test reads the same before and after the
// audit fields move from an embedded struct to direct fields.
func setAudit(t *testing.T, v any, a modelcore.AuditableVO) {
	t.Helper()
	rv := reflect.ValueOf(v).Elem()
	for name, val := range map[string]any{
		"CreatedAt": a.CreatedAt, "CreatedBy": a.CreatedBy, "UpdatedAt": a.UpdatedAt,
		"UpdatedBy": a.UpdatedBy, "DeletedAt": a.DeletedAt, "DeletedBy": a.DeletedBy,
	} {
		f := rv.FieldByName(name)
		if !f.IsValid() {
			t.Fatalf("%T has no field %s", v, name)
		}
		f.Set(reflect.ValueOf(val))
	}
}
