package user_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Page int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
	Key string `json:"key" form:"key"`
}

func (UserApi) UserListView(c *gin.Context) {
	var cr UserListRequest
	c.ShouldBindQuery(&cr)

	var count int
	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&count)

	offset := (cr.Page - 1) * cr.Limit

	var users []models.UserModel
	global.DB.Limit(cr.Limit).Offset(offset).Find(&users)

	res.OKWithList(users, count, c)
}