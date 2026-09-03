package cablesEngineService

import (
	"strings"
	"testing"

	cablesEngineConstants "sharpnr-engine/cables_engine/constants"
	cablesEngineErrs "sharpnr-engine/cables_engine/errs"
)

// The slug is the channel's @mention and URL key, and it is derived rather than
// accepted from the caller — so this function is the only thing standing
// between a display name and a key the client has to be able to route on.
func TestDeriveChannelSlug(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"plain name", "Design Team", "design-team"},
		{"already lowercase", "general", "general"},
		{"trims outer space", "  Front Office  ", "front-office"},
		{"collapses separator runs", "Design   --  Team", "design-team"},
		{"drops punctuation", "Q3 Planning (2026)!", "q3-planning-2026"},
		{"keeps digits", "team 42", "team-42"},
		{"underscores and dots are separators", "eng_team.core", "eng-team-core"},
		{"no leading or trailing hyphen", "-- Ops --", "ops"},

		// Nothing to mention or link: the caller turns this into
		// ErrChannelNameRequired rather than storing an empty slug.
		{"punctuation only", "###", ""},
		{"empty", "", ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := deriveChannelSlug(tc.in); got != tc.want {
				t.Errorf("deriveChannelSlug(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// A name long enough to be truncated must not come back with a dangling
// separator: "design-team-" is not a slug anyone can link to, and it would
// collide differently than the untruncated name.
func TestDeriveChannelSlugTruncatesWithoutTrailingHyphen(t *testing.T) {
	long := strings.Repeat("ab ", cablesEngineConstants.MaxChannelSlugLength)

	slug := deriveChannelSlug(long)

	if len(slug) > cablesEngineConstants.MaxChannelSlugLength {
		t.Fatalf("slug is %d chars, want at most %d", len(slug), cablesEngineConstants.MaxChannelSlugLength)
	}
	if strings.HasPrefix(slug, "-") || strings.HasSuffix(slug, "-") {
		t.Errorf("slug %q has a dangling separator", slug)
	}
}

/* ---------------------------------- <***> --------------------------------- */

// The binding tags validate these too, but the service is also reachable from
// the orchestrator: an enum must not be able to reach the database unchecked.
func TestResolveChannelType(t *testing.T) {
	cases := []struct {
		in      string
		want    cablesEngineConstants.ChannelType
		wantErr error
	}{
		{"", cablesEngineConstants.DefaultChannelType, nil},
		{"TEXT", cablesEngineConstants.ChannelTypeText, nil},
		{"video", cablesEngineConstants.ChannelTypeVideo, nil},
		{"  audio  ", cablesEngineConstants.ChannelTypeAudio, nil},
		{"HOLOGRAM", "", cablesEngineErrs.ErrInvalidChannelType},
	}

	for _, tc := range cases {
		got, err := resolveChannelType(tc.in)
		if err != tc.wantErr {
			t.Errorf("resolveChannelType(%q) error = %v, want %v", tc.in, err, tc.wantErr)
		}
		if got != tc.want {
			t.Errorf("resolveChannelType(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// An empty visibility defaults to PUBLIC; PRIVATE is never implicit.
func TestResolveChannelVisibility(t *testing.T) {
	cases := []struct {
		in      string
		want    cablesEngineConstants.ChannelVisibility
		wantErr error
	}{
		{"", cablesEngineConstants.DefaultChannelVisibility, nil},
		{"PUBLIC", cablesEngineConstants.ChannelVisibilityPublic, nil},
		{"private", cablesEngineConstants.ChannelVisibilityPrivate, nil},
		{"SECRET", "", cablesEngineErrs.ErrInvalidChannelVisibility},
	}

	for _, tc := range cases {
		got, err := resolveChannelVisibility(tc.in)
		if err != tc.wantErr {
			t.Errorf("resolveChannelVisibility(%q) error = %v, want %v", tc.in, err, tc.wantErr)
		}
		if got != tc.want {
			t.Errorf("resolveChannelVisibility(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
