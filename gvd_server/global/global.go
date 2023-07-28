package global

import (
	"gvd_server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Log    *logrus.Logger
	DB     *gorm.DB
)
