package steps

import (
	"fmt"
	"os"

	loggingclient "code.cloudfoundry.org/diego-logging-client"
	"code.cloudfoundry.org/executor"
	"github.com/tedsuo/ifrit"
)

type emitCheckFailureMetricStep struct {
	subStep       ifrit.Runner
	checkProtocol executor.CheckProtocol
	checkType     executor.HealthcheckType
	metronClient  loggingclient.IngressClient
}

const (
	CheckFailedCount = "ChecksFailedCount"
)

func NewEmitCheckFailureMetricStep(
	subStep ifrit.Runner,
	checkProtocol executor.CheckProtocol,
	checkType executor.HealthcheckType,
	metronClient loggingclient.IngressClient) ifrit.Runner {
	return &emitCheckFailureMetricStep{
		subStep:       subStep,
		checkProtocol: checkProtocol,
		checkType:     checkType,
		metronClient:  metronClient,
	}
}

func (step *emitCheckFailureMetricStep) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	if step.subStep == nil {
		return nil
	}

	subStepErr := step.subStep.Run(signals, ready)

	if subStepErr != nil {
		step.emitFailureMetric()
	}

	return subStepErr
}

func (step *emitCheckFailureMetricStep) emitFailureMetric() {
	metricName := constructMetricName(step.checkProtocol, step.checkType)
	go step.metronClient.IncrementCounter(metricName)
}

func constructMetricName(checkProtocol executor.CheckProtocol, checkType executor.HealthcheckType) string {
	return fmt.Sprintf("%s%s%s", checkProtocol, checkType, CheckFailedCount)
}
