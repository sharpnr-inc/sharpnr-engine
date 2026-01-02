package bar_capabilities

type BarCapability string

const (
	BarCreateTableSession BarCapability = "BAR:TABLE_SESSION_CREATE"
	BarOrderItem          BarCapability = "BAR:ORDER_ITEM"
	BarCreateTable        BarCapability = "BAR:CREATE_TABLE"
)
