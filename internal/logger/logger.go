package logger

import "github.com/sirupsen/logrus"

func Init() *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	log.ReportCaller = true
	return log
}
