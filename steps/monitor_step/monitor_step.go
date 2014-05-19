package monitor_step

import (
	"net/http"
	"net/url"
	"time"

	"github.com/cloudfoundry-incubator/executor/sequence"
)

type monitorStep struct {
	check sequence.Step

	interval           time.Duration
	healthyThreshold   uint
	unhealthyThreshold uint

	healthyHook   *url.URL
	unhealthyHook *url.URL
}

func New(
	check sequence.Step,
	interval time.Duration,
	healthyThreshold, unhealthyThreshold uint,
	healthyHook, unhealthyHook *url.URL,
) sequence.Step {
	return &monitorStep{
		check:              check,
		interval:           interval,
		healthyThreshold:   healthyThreshold,
		unhealthyThreshold: unhealthyThreshold,
		healthyHook:        healthyHook,
		unhealthyHook:      unhealthyHook,
	}
}

func (step *monitorStep) Perform() error {
	timer := time.NewTicker(step.interval)

	var healthyCount uint
	var unhealthyCount uint

	for {
		<-timer.C

		healthy := step.check.Perform() == nil

		if healthy {
			healthyCount++
			unhealthyCount = 0
		} else {
			unhealthyCount++
			healthyCount = 0
		}

		var request *http.Request

		if healthyCount >= step.healthyThreshold {
			healthyCount = 0

			request = &http.Request{
				Method: "PUT",
				URL:    step.healthyHook,
			}
		}

		if unhealthyCount >= step.unhealthyThreshold {
			unhealthyCount = 0

			request = &http.Request{
				Method: "PUT",
				URL:    step.unhealthyHook,
			}
		}

		if request != nil {
			resp, err := http.DefaultClient.Do(request)
			if err == nil {
				resp.Body.Close()
			}
		}
	}

	return nil
}

func (step *monitorStep) Cancel() {
}

func (step *monitorStep) Cleanup() {
}