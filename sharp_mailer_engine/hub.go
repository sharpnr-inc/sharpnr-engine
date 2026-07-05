package sharp_mailer_engine

import (
	"html/template"
	"log"
	"sharpnr-engine/config"
	mailerEngineTypes "sharpnr-engine/sharp_mailer_engine/types"
	"sync"
)

type SharpMailerEngine struct {
	cfg           *config.Config
	queues        map[MailType]chan internalEmailJob // unexported — use Enqueue()
	templateCache *template.Template
	wg            sync.WaitGroup
	quit          chan struct{}
}

// New creates the engine and wires up one queue per mail identity.
// Call LoadTemplates() then Start() before sending any emails.
func New(cfg *config.Config) *SharpMailerEngine {
	engine := &SharpMailerEngine{
		cfg:    cfg,
		queues: make(map[MailType]chan internalEmailJob),
		quit:   make(chan struct{}),
	}

	for mailType := range cfg.Email.Identities {
		engine.queues[mailType] = make(chan internalEmailJob, 100)
	}

	return engine
}

// Start launches worker goroutines for every identity in config.
// Call this once after New() and LoadTemplates().
func (e *SharpMailerEngine) Start() {
	for mailType, identity := range e.cfg.Email.Identities {
		poolSize := identity.PoolSize
		if poolSize == 0 {
			poolSize = 3 // sensible default
		}
		e.startPool(e.queues[mailType], poolSize, identity.FromAddress)
	}
}

// Enqueue puts a job on the correct queue. Always non-blocking —
// if the queue is full the job is dropped and logged.
func (e *SharpMailerEngine) Enqueue(mailType MailType, job internalEmailJob) {
	queue, ok := e.queues[mailType]
	if !ok {
		log.Printf("[mailer] no queue for mail type %q — is it in config?", mailType)
		return
	}

	select {
	case queue <- job:
		// delivered to queue
	default:
		log.Printf("[mailer] queue full for %q, dropping email to %s", mailType, job.To)
	}
}

// Shutdown signals all workers to stop and waits for them to finish.
func (e *SharpMailerEngine) Shutdown() {
	close(e.quit)
	e.wg.Wait()
}

// ── Typed job constructors ─────────────────────────────────────────
// These are the only way to create jobs — keeps Data typed at the call site.

func NewOTPJob(to, subject string, data OTPData) internalEmailJob {
	return internalEmailJob{To: to, Subject: subject, Template: TmplVerification, Data: data}
}

func NewNotificationJob(to, subject string, data NotificationData) internalEmailJob {
	return internalEmailJob{To: to, Subject: subject, Template: TmplNotification, Data: data}
}

func NewBillingJob(to, subject string, data BillingData) internalEmailJob {
	return internalEmailJob{To: to, Subject: subject, Template: TmplInvoice, Data: data}
}

func NewWelcomeJob(to, subject string, data WelcomeData) internalEmailJob {
	return internalEmailJob{To: to, Subject: subject, Template: TmplWelcome, Data: data}
}

// keep compiler happy if mailer import is only used for the alias
var _ mailerEngineTypes.Type
