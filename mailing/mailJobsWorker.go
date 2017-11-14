package mailing

import (
	log "github.com/sirupsen/logrus"
	. "github.com/resurtm/boomak-server/mailing/base"
)

func mailJobsWorker(workerID byte, ch <-chan MailJob) {
	log.WithField("worker_id", workerID).Debug("mail worker started")
	client := newClient()
	for {
		job := <-ch
		log.WithFields(log.Fields{
			"worker_id": workerID,
			"job_kind":  job.Kind,
		}).Debug("started processing mail job")

		if email := emailBuilders[job.Kind](job.Payload); email == nil {
			log.WithFields(log.Fields{
				"worker_id": workerID,
				"job_kind":  job.Kind,
			}).Warn("email cannot be prepared, just skipping the job")
		} else {
			log.WithFields(log.Fields{
				"worker_id": workerID,
				"job_kind":  job.Kind,
				"email":     email,
			}).Debug("email prepared")

			if err := client.SendEmail(*email); err != nil {
				log.WithFields(log.Fields{
					"worker_id": workerID,
					"job_kind":  job.Kind,
					"error":     err,
				}).Debug("mail job process failure")
			}
		}

		log.WithFields(log.Fields{
			"worker_id": workerID,
			"job_kind":  job.Kind,
		}).Debug("finished processing mail job")
	}
	log.WithField("worker_id", workerID).Debug("mail worker stopped")
}
