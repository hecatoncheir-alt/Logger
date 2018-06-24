package logger

import (
	"encoding/json"
	"time"

	"fmt"
	"github.com/hecatoncheir/Broker"
)

type LogData struct {
	Message, Level string
	Time           time.Time
}

type Writer interface {
	Write(data LogData) error
}

type LogWriter struct {
	APIVersion  string
	LoggerTopic string
	ServiceName string
	bro         *broker.Broker
}

func New(apiVersion, serviceName, topicForWriteLog string, broker *broker.Broker) *LogWriter {
	logger := LogWriter{LoggerTopic: topicForWriteLog, bro: broker, ServiceName: serviceName, APIVersion: apiVersion}
	return &logger
}

func (logWriter *LogWriter) Write(data LogData) error {
	if data.Time.IsZero() {
		println(fmt.Sprintf("LogData: %v without time", data))
		data.Time = time.Now()
	}

	item := map[string]interface{}{
		"Service": data.Level, "Time": data.Time}

	eventData, err := json.Marshal(item)
	if err != nil {
		return err
	}

	event := broker.EventData{
		Message: data.Message, Data: string(eventData)}

	logWriter.bro.Write(event)

	return nil
}
