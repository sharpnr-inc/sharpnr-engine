package sharp_mailer_engine

import (
	"html/template"
	"log"
	"sharpnr-engine/config"
	"sync"
)

const (
	defaultPoolSize  = 3
	defaultQueueSize = 100
)

type SharpMailerEngine struct {
	cfg           *config.Config
	queues        map[MailType]chan internalEmailJob // unexported — use Enqueue()
	templateCache *template.Template
	wg            sync.WaitGroup

	mu      sync.RWMutex
	stopped bool
}

// New creates the engine and wires up one queue per mail identity.
// Call LoadTemplates() then Start() before sending any emails.
func New(cfg *config.Config) *SharpMailerEngine {
	engine := &SharpMailerEngine{
		cfg:    cfg,
		queues: make(map[MailType]chan internalEmailJob),
	}

	for mailType, identity := range cfg.Email.Identities {
		queueSize := identity.QueueSize
		if queueSize == 0 {
			queueSize = defaultQueueSize
		}
		engine.queues[mailType] = make(chan internalEmailJob, queueSize)
	}

	return engine
}

// Start launches worker goroutines for every identity in config.
// Call this once after New() and LoadTemplates().
func (e *SharpMailerEngine) Start() {
	for mailType, identity := range e.cfg.Email.Identities {
		poolSize := identity.PoolSize
		if poolSize == 0 {
			poolSize = defaultPoolSize
		}
		e.startPool(e.queues[mailType], poolSize, identity.FromAddress)
	}
}

// Enqueue puts a job on the queue matching its mail type. Always
// non-blocking — if the queue is full or the engine has been shut
// down, the job is dropped and logged.
func (e *SharpMailerEngine) Enqueue(job internalEmailJob) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.stopped {
		log.Printf("[mailer] engine is shut down, dropping email to %s", job.To)
		return
	}

	queue, ok := e.queues[job.mailType]
	if !ok {
		log.Printf("[mailer] no queue for mail type %q — is it in config?", job.mailType)
		return
	}

	select {
	case queue <- job:
		// delivered to queue
	default:
		log.Printf("[mailer] queue full for %q, dropping email to %s", job.mailType, job.To)
	}
}

// Shutdown stops accepting new jobs, then waits for the workers to
// drain the queues and finish any in-flight sends.
func (e *SharpMailerEngine) Shutdown() {
	e.mu.Lock()
	if e.stopped {
		e.mu.Unlock()
		return
	}
	e.stopped = true
	for _, queue := range e.queues {
		close(queue)
	}
	e.mu.Unlock()

	e.wg.Wait()
}

// ── Typed job constructors ─────────────────────────────────────────
// These are the only way to create jobs — they keep Data typed at the
// call site and bind each job to its sending identity.

func NewOTPJob(to, subject string, data OTPData) internalEmailJob {
	return internalEmailJob{mailType: MailTypeOTP, To: to, Subject: subject, Template: TmplVerification, Data: data}
}

func NewNotificationJob(to, subject string, data NotificationData) internalEmailJob {
	return internalEmailJob{mailType: MailTypeNotification, To: to, Subject: subject, Template: TmplNotification, Data: data}
}

func NewBillingJob(to, subject string, data BillingData) internalEmailJob {
	return internalEmailJob{mailType: MailTypeBilling, To: to, Subject: subject, Template: TmplInvoice, Data: data}
}

func NewWelcomeJob(to, subject string, data WelcomeData) internalEmailJob {
	return internalEmailJob{mailType: MailTypeNotification, To: to, Subject: subject, Template: TmplWelcome, Data: data}
}
