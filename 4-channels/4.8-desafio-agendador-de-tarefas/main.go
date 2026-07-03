package main

import (
	"fmt"
	"time"
)

type JobMessage struct {
	message string
	date    time.Time
	isSent  bool
}

func hasPendingJobs(messages []JobMessage) bool {
	for _, job := range messages {
		if !job.isSent {
			return true
		}
	}
	return false
}

func cron(messages []JobMessage, now time.Time) {
	for i := range messages {
		job := &messages[i]

		if !job.isSent && !now.Before(job.date) {
			fmt.Printf("[%s] Executando JOB: %s\n",
				now.Format("15:04:05"),
				job.message,
			)

			job.isSent = true
		}
	}
}

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	jobMessages := []JobMessage{
		{
			message: "JOB 1",
			date:    time.Now().Add(10 * time.Second),
		},
		{
			message: "JOB 2",
			date:    time.Now().Add(20 * time.Second),
		},
		{
			message: "JOB 3",
			date:    time.Now().Add(30 * time.Second),
		},
	}

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				cron(jobMessages, time.Now())

				if !hasPendingJobs(jobMessages) {
					fmt.Println("Todos os JOBS executados.")
					close(done)
					return
				}

			case <-done:
				return
			}
		}
	}()

	<-done
	fmt.Println("Processo encerrado.")
}
