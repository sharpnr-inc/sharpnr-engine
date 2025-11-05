package sharpnreventbus

// Right Now i am using the interface but later we will use the
// Generics
type Event struct {
	Name    string
	Payload interface{}
}

type EventHandler func(event Event)
