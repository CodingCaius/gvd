package user_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	//bingding字段 定义字段在数据绑定（比如 HTTP 请求参数绑定）时的行为
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID" binding:"required" label:"角色ID"` //角色ID

}

// 处理 创建用户 的 API 请求
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	//将接收到的 JSON 数据解析到 cr 中
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err == nil {
		res.FailWithMsg("用户名已存在", c)
		return
	}
	err = global.DB.Create(&models.UserModel{
		UserName: cr.UserName,
		Password: pwd.HashPwd(cr.Password),
		NickName: cr.NickName,
		IP: c.RemoteIP(),
		RoleID: cr.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户创建失败", c)
		return
	}

	res.OKWithMsg("用户创建成功", c)
}
