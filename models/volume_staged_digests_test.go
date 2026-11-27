package models

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestVolumeVersionStagedDigestsRoundTrip(t *testing.T) {
	cover := "cover-digest"
	want := VolumeVersion{StagedCoverDigest: &cover, StagedSampleDigests: []string{"a", "b"}}

	raw, err := bson.Marshal(want)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	var got VolumeVersion
	if err := bson.Unmarshal(raw, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	if got.StagedCoverDigest == nil || *got.StagedCoverDigest != cover {
		t.Errorf("StagedCoverDigest = %v, want %q", got.StagedCoverDigest, cover)
	}
	if !reflect.DeepEqual(got.StagedSampleDigests, want.StagedSampleDigests) {
		t.Errorf("StagedSampleDigests = %v, want %v", got.StagedSampleDigests, want.StagedSampleDigests)
	}
}

func TestVolumeVersionWithoutStagedDigestsStoresNoFields(t *testing.T) {
	raw, err := bson.Marshal(VolumeVersion{})
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	for _, key := range []string{"staged_cover_digest", "staged_sample_digests"} {
		if _, err := bson.Raw(raw).LookupErr(key); err == nil {
			t.Errorf("%s stored for a version with no digests", key)
		}
	}
}
