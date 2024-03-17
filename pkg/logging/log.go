package logging

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type (
	Level  int
	Fields map[string]interface{}
)

const (
	DEBUG Level = iota
	INFO
	ERROR
)

type Message struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Time    string                 `json:"time"`
	Fields  map[string]interface{} `json:"fields,omitempty"`
}

type Logger struct {
	level   Level
	encoder Encoder
	// loc     *time.Location
}

type FieldEncoder struct {
	Encoder
	Fields map[string]interface{}
}

type Encoder interface {
	Encode(msg Message)
}

type JsonEncoder struct {
	enc *json.Encoder
}

func (je JsonEncoder) Encode(msg Message) {
	err := je.enc.Encode(msg)
	if err != nil {
		fmt.Printf("{\"level\":\"ERROR\",\"message\":\"failed to encode message\",\"time\":\"%s\"}", time.Now().Format(time.RFC3339))
	}
}

func New(level Level) *Logger {
	// loc, err := time.LoadLocation("America/New_York") // Eastern Time zone
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// currentTime := time.Now().In(loc)
	return &Logger{
		level:   level,
		encoder: JsonEncoder{enc: json.NewEncoder(os.Stdout)},
		// loc:     loc,
	}
}

func (l *Logger) Debug(msg string) {
	if l.level <= DEBUG {
		l.output("DEBUG", msg, nil)
	}
}

func (l *Logger) Info(msg string) {
	if l.level <= INFO {
		l.output("INFO", msg, nil)
	}
}

func (l *Logger) Error(msg string) {
	if l.level <= ERROR {
		l.output("ERROR", msg, nil)
	}
}

func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	return &Logger{
		level:   l.level,
		encoder: &FieldEncoder{Encoder: l.encoder, Fields: fields},
	}
}

func (fe *FieldEncoder) Encode(msg Message) {
	msg.Fields = fe.Fields
	fe.Encoder.Encode(msg)
}

func (l *Logger) output(level string, msg string, fields map[string]interface{}) {
	loc, _ := time.LoadLocation("America/New_York") // Eastern Time zone

	message := Message{
		Time:    time.Now().In(loc).Format(time.RFC3339),
		Level:   level,
		Message: msg,
		Fields:  fields,
	}

	l.encoder.Encode(message)
}
