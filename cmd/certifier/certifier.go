package certifier

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) error {
	logrus.Info("Running certifier")
	for ctx.Err() == nil {
		time.Sleep(time.Second)
		logrus.Info("Looping !")
	}
	logrus.Info("Ciao")
	return nil
}
