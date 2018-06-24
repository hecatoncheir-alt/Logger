package logger

import (
	"testing"
	"time"

	"github.com/hecatoncheir/Broker"
	"github.com/hecatoncheir/Configuration"
)

func TestLoggerCanWriteLogData(test *testing.T) {
	conf := configuration.New()

	bro := broker.New(conf.APIVersion, "First test service for write log data")
	err := bro.Connect(conf.Development.EventBus.Host, conf.Development.EventBus.Port)
	if err != nil {
		test.Error(err)
	}

	logWriter := New(conf.APIVersion, conf.ServiceName, conf.Development.LogunaTopic, bro)
	logData := LogData{Message: "test message", Time: time.Now().UTC()}
	go func() {
		err := logWriter.Write(logData)
		if err != nil {
			test.Error(err)
		}
	}()

	otherBrokerConnection := broker.New(conf.APIVersion, "Second test service for receive log data")
	err = bro.Connect(conf.Development.EventBus.Host, conf.Development.EventBus.Port)
	if err != nil {
		test.Error(err)
	}

	for event := range otherBrokerConnection.InputChannel {
		if event.Message == "test message" {
			break
		}

		if event.Message != "test message" {
			test.Fail()
			break
		}

		if event.ServiceName == "First test service for write log data" {
			break
		}

		if event.ServiceName != "First test service for write log data" {
			test.Fail()
			break
		}
	}

}
