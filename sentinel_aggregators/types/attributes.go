// Package sentinelAggregatorTypes holds the value types shared by the resource
// aggregator and every module's mod_resources.
//
// They exist to stop the catalog being stringly-typed. A resource attribute
// declared as "boolean" instead of "bool" produces no compile error and no
// runtime error — it silently renders a broken picker in the policy builder.
// Making these named types with a Valid() check lets the aggregator reject them
// at bootstrap instead.
package sentinelAggregatorTypes

// AttrType is the Cedar primitive of a resource attribute a policy condition
// can test (`resource.<name> == ...`).
type AttrType string

const (
	AttrString AttrType = "string"
	AttrLong   AttrType = "long"
	AttrBool   AttrType = "bool"
)

// Valid reports whether t is a Cedar primitive the policy builder can render.
func (t AttrType) Valid() bool {
	switch t {
	case AttrString, AttrLong, AttrBool:
		return true
	}
	return false
}

// AllAttrTypes is every valid attribute type, for error messages and tests.
var AllAttrTypes = []AttrType{AttrString, AttrLong, AttrBool}
