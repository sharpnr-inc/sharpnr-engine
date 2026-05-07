package sharp_mailer_engine

import "log"

func (e *SharpMailerEngine) startPool(queue chan internalEmailJob, count int, from string) {
	for i := 0; i < count; i++ {
		e.wg.Add(1)

		go func(workerID int) {
			defer e.wg.Done()

			var client *smtpClientWrapper

			for {
				select {
				case job := <-queue:
					// Render the HTML template with the job's data
					body, err := e.parse(job.Template, job.Data)
					if err != nil {
						log.Printf("[worker %d] template error for %s: %v", workerID, job.Template, err)
						continue
					}

					// Dial a new SMTP connection if needed
					if client == nil || client.IsClosed() {
						client, err = e.dialNewClient()
						if err != nil {
							log.Printf("[worker %d] dial error: %v", workerID, err)
							continue
						}
					}

					// Send — on failure close and force redial next iteration
					if err := e.sendPooled(client.smtp, job.To, job.Subject, body, from); err != nil {
						log.Printf("[worker %d] send failed to %s: %v", workerID, job.To, err)
						client.smtp.Close()
						client = nil
					}

				case <-e.quit:
					// Graceful shutdown — close the SMTP connection cleanly
					if client != nil {
						client.smtp.Quit()
					}
					return
				}
			}
		}(i)
	}
}
