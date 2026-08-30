// Package sentinelAggregatorTypes holds the value types shared by the resource
// aggregator, every module's mod_resources, and the context palette.
//
// They exist to stop the catalog being stringly-typed. An attribute declared as
// "boolean" instead of "bool" produces no compile error and no runtime error —
// it silently renders a broken picker in the policy builder. Making these named
// types with a Valid() check lets the aggregator reject them at bootstrap.
package sentinelAggregatorTypes

/* -------------------------------------------------------------------------- */
/*                              Attribute types                               */
/* -------------------------------------------------------------------------- */

// AttrType is the Cedar shape of a value a policy condition can test.
type AttrType string

const (
	AttrString   AttrType = "string"
	AttrLong     AttrType = "long"
	AttrBool     AttrType = "bool"
	AttrDatetime AttrType = "datetime"
)

// Valid reports whether t is a type the policy builder can render.
func (t AttrType) Valid() bool {
	switch t {
	case AttrString, AttrLong, AttrBool, AttrDatetime:
		return true
	}
	return false
}

// AllAttrTypes is every valid attribute type, for error messages and tests.
var AllAttrTypes = []AttrType{AttrString, AttrLong, AttrBool, AttrDatetime}

/* -------------------------------------------------------------------------- */
/*                                 Operators                                  */
/* -------------------------------------------------------------------------- */

// Operator is how a restriction compares an attribute to its value(s).
//
// There is deliberately no "unless" concept anywhere in this vocabulary:
// `unless X` is exactly `when !X`. The second keyword exists so hand-written
// Cedar reads well, and a generator does not need it. Negation lives HERE, as
// an operator, so the person building a rule never chooses between two keywords
// that mean the same thing — and every rule emits a single `when { ... }`.
type Operator string

const (
	OpIs       Operator = "is"       // x == v
	OpIsNot    Operator = "isNot"    // x != v
	OpIsOneOf  Operator = "isOneOf"  // [v...].contains(x)
	OpIsNoneOf Operator = "isNoneOf" // ![v...].contains(x)
)

// Valid reports whether o is an operator the generator knows how to render.
func (o Operator) Valid() bool {
	switch o {
	case OpIs, OpIsNot, OpIsOneOf, OpIsNoneOf:
		return true
	}
	return false
}

// DefaultOperators is the operator set for an attribute of type t.
//
// A bool only ever reads as itself or its negation, so offering "is one of" on
// one would render a multi-select over {yes, no} — a control that cannot express
// anything the single picker cannot.
func DefaultOperators(t AttrType) []Operator {
	switch t {
	case AttrBool:
		return []Operator{OpIs}
	case AttrLong, AttrDatetime:
		return []Operator{OpIs, OpIsNot}
	default:
		return []Operator{OpIs, OpIsNot, OpIsOneOf, OpIsNoneOf}
	}
}
