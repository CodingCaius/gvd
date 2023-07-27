package global

import (
	"gvd_server/config"

	"github.com/sirupsen/logrus"
)

var (
	Config *config.Config
	Log *logrus.Logger
)
