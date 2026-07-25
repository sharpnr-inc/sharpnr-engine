package sentinelContexts

import (
	"context"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

// The palette is what the FE builder authors conditions against, and CedarKeys
// is what the engine actually stamps. If they drift, authored policies condition
// on keys that never exist — and in Cedar a missing attribute silently disables
// the policy. This test pins them together.
func TestBuilderPaletteCoversAllStampedKeys(t *testing.T) {
	palette := BuilderPalette()

	paletteKeys := func(ds []KeyDescriptor) map[string]KeyDescriptor {
		out := make(map[string]KeyDescriptor, len(ds))
		for _, d := range ds {
			out[d.Key] = d
		}
		return out
	}

	requestPalette := paletteKeys(palette.Request)
	for key := range (RequestContext{}).CedarKeys() {
		d, ok := requestPalette[key]
		require.True(t, ok, "stamped context key %q is missing from the builder palette", key)
		assert.Equal(t, "context."+key, d.Cedar)
		assert.NotEmpty(t, d.Type, "context key %q needs a type", key)
	}
	assert.Len(t, palette.Request, len((RequestContext{}).CedarKeys()),
		"palette lists a context key the engine never stamps")

	principalPalette := paletteKeys(palette.Principal)
	for key := range (UserContext{}).CedarKeys() {
		d, ok := principalPalette[key]
		require.True(t, ok, "stamped principal key %q is missing from the builder palette", key)
		assert.Equal(t, "principal."+key, d.Cedar)
		assert.NotEmpty(t, d.Type, "principal key %q needs a type", key)
	}
	assert.Len(t, palette.Principal, len((UserContext{}).CedarKeys()),
		"palette lists a principal key the engine never stamps")
}

func TestInjectExtractRoundtrip(t *testing.T) {
	authz := AuthzContext{
		Request: RequestContext{Channel: ChannelMobile, IPAddress: "10.0.0.1", MFAVerified: true},
		User:    UserContext{Department: "Finance", Designation: "Manager"},
	}

	got, ok := Extract(Inject(context.Background(), authz))
	require.True(t, ok)
	assert.Equal(t, authz, got)

	_, ok = Extract(context.Background())
	assert.False(t, ok, "a bare context carries no AuthzContext")
}
