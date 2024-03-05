package workerpool

import (
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

func newWorker(timeout time.Duration) *worker {
	return &worker{
		&http.Client{
			Timeout: timeout,
		},
	}
}

func (w worker) process(j Job) Result {
	result := Result{URL: j.URL}

	now := time.Now()

	response, err := w.client.Get(j.URL)
	if err != nil {
		result.Error = err

		return result
	}

	result.StatusCode = response.StatusCode
	result.ResponseTime = time.Since(now)

	return result
}
