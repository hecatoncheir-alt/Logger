package logger

import (
	"testing"
	"time"

	"github.com/hecatoncheir/Broker"
	"github.com/hecatoncheir/Configuration"
)

func TestLoggerCanWriteLogData(test *testing.T) {
	conf := configuration.New()

	bro := broker.New(conf.APIVersion, conf.ServiceName)
	bro.Connect(conf.Development.Broker.Host, conf.Development.Broker.Port)

	logWriter := New(conf.APIVersion, conf.ServiceName, conf.Development.LogunaTopic, bro)
	logData := LogData{Message: "test message", Time: time.Now().UTC()}
	go logWriter.Write(logData)

	logunaTopic, err := bro.ListenTopic(conf.Development.LogunaTopic, conf.APIVersion)
	if err != nil {
		test.Fatal(err)
	}

	for event := range logunaTopic {
		if event.Message != "test message" {
			test.Fail()
		}

		break
	}

}
