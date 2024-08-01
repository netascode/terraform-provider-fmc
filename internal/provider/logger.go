package provider

import (
	"fmt"
	"log"
	"os"
)

type logger struct {
	logger *log.Logger
}

var Log *logger

func init() {
	var err error
	Log, err = newLogger("", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		log.Fatal(err)
	}
}

func newLogger(prefix string, flag int) (*logger, error) {
	logDir := "/tmp/netascode/fmc/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	logFileName := fmt.Sprintf("%s/%s.logs", logDir, "latest")
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &logger{
		logger: log.New(file, prefix, flag),
	}, nil
}

func (l *logger) Debug(v ...interface{}) {
	l.logger.SetPrefix("[DEBUG] ")
	l.logger.Println(v...)

}
