package flags

import (
	"gvd_server/global"
	"gvd_server/models"

	"github.com/sirupsen/logrus"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.RoleModel{},
			&models.DocModel{},
			//如果有自定义连接表，一定要放在另外两张表的后面
			&models.UserCollDocModel{},
			&models.RoleDocModel{},
			&models.ImageModel{},
			&models.UserPwdDocModel{},
			&models.LoginModel{},
		)

	if err != nil {
		logrus.Fatalf("数据库迁移失败 err:%s", err.Error())
	}
	logrus.Infof("数据库迁移成功!")

}
