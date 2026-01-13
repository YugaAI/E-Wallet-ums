package helpers

import "github.com/sirupsen/logrus"

func SetupLog() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	log.Info("Logger initialited with logrus")
	return log
}
