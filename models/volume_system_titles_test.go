package models

import (
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

// SystemTitles must survive a BSON and a JSON round trip keyed by system ID, alongside an
// unchanged SystemIds list.
func TestVolumeVersionSystemTitlesRoundTrip(t *testing.T) {
	in := VolumeVersion{
		ID:           "vv1",
		RecordID:     "vol1",
		Version:      3,
		Title:        "Numenera",
		SystemIds:    []string{"sysA", "sysB"},
		SystemTitles: map[string]string{"sysA": "Numenera", "sysB": "The Strange"},
		State:        VersionStateLive,
	}

	bsonBytes, err := bson.Marshal(in)
	if err != nil {
		t.Fatalf("bson marshal: %v", err)
	}
	var fromBSON VolumeVersion
	if err := bson.Unmarshal(bsonBytes, &fromBSON); err != nil {
		t.Fatalf("bson unmarshal: %v", err)
	}
	assertTitles(t, "bson", fromBSON.SystemIds, fromBSON.SystemTitles)

	jsonBytes, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("json marshal: %v", err)
	}
	var fromJSON VolumeVersion
	if err := json.Unmarshal(jsonBytes, &fromJSON); err != nil {
		t.Fatalf("json unmarshal: %v", err)
	}
	assertTitles(t, "json", fromJSON.SystemIds, fromJSON.SystemTitles)
}

// An absent system_titles field decodes to a nil map, not an error.
func TestVolumeVersionSystemTitlesAbsentIsNil(t *testing.T) {
	doc, err := bson.Marshal(VolumeVersion{ID: "vv1", SystemIds: []string{"sysA"}})
	if err != nil {
		t.Fatalf("bson marshal: %v", err)
	}
	var out VolumeVersion
	if err := bson.Unmarshal(doc, &out); err != nil {
		t.Fatalf("bson unmarshal: %v", err)
	}
	if out.SystemTitles != nil {
		t.Fatalf("absent system_titles decoded to %v, want nil", out.SystemTitles)
	}
}

func assertTitles(t *testing.T, kind string, ids []string, titles map[string]string) {
	t.Helper()
	if len(ids) != 2 || ids[0] != "sysA" || ids[1] != "sysB" {
		t.Errorf("%s: system_ids = %v, want [sysA sysB]", kind, ids)
	}
	if titles["sysA"] != "Numenera" || titles["sysB"] != "The Strange" {
		t.Errorf("%s: system_titles = %v", kind, titles)
	}
}
