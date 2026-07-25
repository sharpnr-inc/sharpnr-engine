package sharp_mailer_engine

import "log"

func (e *SharpMailerEngine) startPool(queue chan internalEmailJob, count int, from string) {
	for i := range count {
		e.wg.Add(1)

		go func(workerID int) {
			defer e.wg.Done()

			var client *smtpClientWrapper

			// Loop ends when Shutdown() closes the queue — after all
			// pending jobs have been drained.
			for job := range queue {
				// Render the HTML template with the job's data
				body, err := e.parse(job.Template, job.Data)
				if err != nil {
					log.Printf("[worker %d] template error for %s: %v", workerID, job.Template, err)
					continue
				}

				// One retry with a fresh connection covers the common
				// case of a pooled connection going stale between the
				// IsClosed() check and the actual send.
				var sendErr error
				for range 2 {
					if client == nil || client.IsClosed() {
						client, sendErr = e.dialNewClient()
						if sendErr != nil {
							log.Printf("[worker %d] dial error: %v", workerID, sendErr)
							client = nil
							continue
						}
					}

					sendErr = e.sendPooled(client, job.To, job.Subject, body, from)
					if sendErr == nil {
						break
					}

					// Send failed — close and force a redial on retry
					client.smtp.Close()
					client = nil
				}
				if sendErr != nil {
					log.Printf("[worker %d] send failed to %s: %v", workerID, job.To, sendErr)
				}
			}

			// Queue closed and drained — close the SMTP connection cleanly
			if client != nil {
				client.smtp.Quit()
			}
		}(i)
	}
}
