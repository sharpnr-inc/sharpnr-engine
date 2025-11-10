package sharpnreventbus

// Subscriber defines a module that listens for certain event types.
type Subscriber interface {
	SubscribedEvents() []string
	HandleEvent(event SharpEvent)
}
