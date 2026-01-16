package sharp_mailer_engine

import "log"

func (e *SharpMailerEngine) Start() {
	for mType, identity := range e.cfg.Email.Identities {
		// Start 3 persistent workers per email type (verification, receipts, etc.)
		e.startPool(e.Queues[MailType(mType)], 3, identity.FromAddress)
	}
}

func (e *SharpMailerEngine) startPool(queue chan EmailJob, count int, from string) {
	for i := 0; i < count; i++ {
		e.wg.Add(1)
		go func(workerID int) {
			defer e.wg.Done()
			var client *smtpClientWrapper // We'll define this in sender.go

			for {
				select {
				case job := <-queue:
					body, err := e.parse(job.Template, job.Data)
					if err != nil {
						log.Printf("[Worker %d] Template Error: %v", workerID, err)
						continue
					}

					// If connection is dead or nil, dial a new one
					if client == nil || client.IsClosed() {
						client, err = e.dialNewClient()
						if err != nil {
							log.Printf("[Worker %d] Dial Error: %v", workerID, err)
							continue
						}
					}

					// Send using the persistent connection
					if err := e.sendPooled(client.smtp, job.To, job.Subject, body, from); err != nil {
						log.Printf("[Worker %d] Send Failed: %v. Retrying connection...", workerID, err)
						client.smtp.Close()
						client = nil // Force redial on next job
					}

				case <-e.quit:
					if client != nil {
						client.smtp.Quit()
					}
					return
				}
			}
		}(i)
	}
}
