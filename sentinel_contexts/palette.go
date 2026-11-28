package sentinelContexts

// The policy builder's context palette: which context.* and principal.* keys a
// policy condition can be authored against, with enough shape (type, closed
// value sets) for the FE to render pickers instead of free-text fields. It
// lives HERE — next to the key constants and the structs the engine stamps —
// so adding a key and publishing it to the builder is one change in one
// package; the aggregator only relays it.

// KeyDescriptor describes one authorable attribute key.
type KeyDescriptor struct {
	Key         string `json:"key"`   // attribute name, e.g. "channel"
	Cedar       string `json:"cedar"` // how policy text reads it, e.g. `context.channel`
	Type        string `json:"type"`  // "string" | "bool" | "datetime"
	Label       string `json:"label"`
	Description string `json:"description"`
	// Values is the closed set of valid values, when there is one; empty means
	// free-form. The builder should render a picker for closed sets.
	Values []string `json:"values,omitempty"`
}

// Palette is the full context vocabulary for the builder, split by where the
// key lives in Cedar: request environment (`context.*`) vs principal facts
// (`principal.*`).
//
// principal.orgId and resource.orgId are deliberately NOT listed: the engine
// enforces the tenant boundary structurally, so authoring conditions on them
// is redundant and only invites confusion.
type Palette struct {
	Request   []KeyDescriptor `json:"request"`
	Principal []KeyDescriptor `json:"principal"`
}

// BuilderPalette returns the palette of every canonical key the engine stamps.
// Keep it in lockstep with RequestContext/UserContext (a test guards this).
func BuilderPalette() Palette {
	return Palette{
		Request: []KeyDescriptor{
			{
				Key:         ContextKeyChannel,
				Cedar:       "context." + ContextKeyChannel,
				Type:        "string",
				Label:       "Channel",
				Description: "Which client the request came from",
				Values:      []string{ChannelWeb, ChannelMobile, ChannelMCP},
			},
			{
				Key:         ContextKeyIP,
				Cedar:       "context." + ContextKeyIP,
				Type:        "string",
				Label:       "Client IP",
				Description: "IP address the request originated from",
			},
			{
				Key:         ContextKeyRequestedAt,
				Cedar:       "context." + ContextKeyRequestedAt,
				Type:        "datetime",
				Label:       "Requested at",
				Description: "When the request happened (for time-window rules)",
			},
			{
				Key:         ContextKeyMFAVerified,
				Cedar:       "context." + ContextKeyMFAVerified,
				Type:        "bool",
				Label:       "MFA verified",
				Description: "Whether the session completed multi-factor authentication",
			},
		},
		Principal: []KeyDescriptor{
			{
				Key:         PrincipalKeyDepartment,
				Cedar:       "principal." + PrincipalKeyDepartment,
				Type:        "string",
				Label:       "Department",
				Description: "The employee's department",
			},
			{
				Key:         PrincipalKeyDesignation,
				Cedar:       "principal." + PrincipalKeyDesignation,
				Type:        "string",
				Label:       "Designation",
				Description: "The employee's designation",
			},
		},
	}
}
