package cookieManager

var registry map[Type]Definition

func Init(defs map[Type]Definition) {
	registry = defs
}

// This Init function is Initiated in main.go
// it Takes Cookie config from the from the config package
