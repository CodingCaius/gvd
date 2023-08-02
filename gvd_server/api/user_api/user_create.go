package user_api

import (
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
	RoleID uint `json:"roleID" binding:"required"` //角色IDjson:""userName" binding:"required"

}

//处理 创建用户 的 API 请求
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	//将接收到的 JSON 数据解析到 cr 中
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数校验错误", c)
		return
	}
	res.OKWithMsg("创建成功", c)
}
