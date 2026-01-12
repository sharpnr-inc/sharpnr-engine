package sharp_mailer_engine

import "log"

func (e *SharpMailerEngine) Start() {
	for mType, identity := range e.cfg.Email.Identities {
		e.startPool(e.Queues[MailType(mType)], 3, identity.FromAddress)
	}
}

func (e *SharpMailerEngine) startPool(queue chan EmailJob, count int, from string) {
	for i := 0; i < count; i++ {
		e.wg.Add(1)
		go func() {
			defer e.wg.Done()
			for {
				select {
				case job := <-queue:
					body, err := e.parse(job.Template, job.Data)
					if err != nil {
						log.Printf("[Mailer] Template Error %v", err)
						continue
					}
					if err := e.send(job.To, job.Subject, body, from); err != nil {
						log.Printf("[Mailer] Critical Error from %s: %v", from, err)
					}
				case <-e.quit:
					return
				}
			}
		}()
	}
}
