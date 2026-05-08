package mailer

type Identity struct {
	FromAddress string // "Sharpnr
	ReplyTo     string // optional
	PoolSize    int    // optional, defaults to 5 if zero
}
