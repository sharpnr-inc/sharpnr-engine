package sharp_mailer_engine

import (
	"html/template"
	"sharpnr-engine/config"
	"sync"
)

type SharpMailerEngine struct {
	cfg           *config.Config
	Queues        map[MailType]chan EmailJob
	templateCache *template.Template
	wg            sync.WaitGroup
	quit          chan struct{}
}

func New(cfg *config.Config) *SharpMailerEngine {
	engine := &SharpMailerEngine{
		cfg:    cfg,
		Queues: map[MailType]chan EmailJob{},
		quit:   make(chan struct{}),
	}

	for mailType := range cfg.Email.Identities {
		engine.Queues[MailType(mailType)] = make(chan EmailJob, 100)
	}

	return engine
}

func (e *SharpMailerEngine) Shutdown() {
	close(e.quit)
	e.wg.Wait()
}
