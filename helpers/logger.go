package helpers

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func SetupLog() {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	log.Info("Logger initialited with logrus")
	Logger = log
}
