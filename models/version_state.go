package models

// VersionState is the lifecycle state of a version record, shared across every versioned
// entity type (volume, publisher, studio, person, license, system).
type VersionState string

const (
	VersionStateSubmitted         VersionState = "submitted"
	VersionStateLive              VersionState = "live"
	VersionStateArchived          VersionState = "archived"
	VersionStateRejected          VersionState = "rejected"
	VersionStatePartiallyAccepted VersionState = "partially_accepted"
)
