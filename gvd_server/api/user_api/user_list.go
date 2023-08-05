package user_api

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"`
}

func (UserApi) UserListView(c *gin.Context) {
	var cr UserListRequest
	c.ShouldBindQuery(&cr)

	var count int

	if cr.Limit < 0 {
		cr.Limit = 10
	}

	offset := (cr.Page - 1) * cr.Limit

	//使用参数化查询防止sql注入
	//避免直接将参数插入sql语句，或者拼接sql语句
	//前端传过来的值，一定要用 问号
	query := global.DB.Where("nickName like ?", fmt.Sprintf("%%%s%%", cr.Key))

	var users []models.UserModel
	global.DB.Model(models.UserModel{}).Where(query).Select("count(id)").Scan(&count)
	global.DB.Limit(cr.Limit).Offset(offset).Find(&users)

	res.OKWithList(users, count, c)
}
