package mailerEngineTypes

type Identity struct {
	FromAddress string // e.g. "noreply@sharpnr.com"
	ReplyTo     string // optional
	PoolSize    int    // optional, defaults to 3 if zero
	QueueSize   int    // optional, defaults to 100 if zero
}
