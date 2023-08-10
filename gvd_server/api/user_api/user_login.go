package user_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"gvd_server/utils/pwd"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}


// UserListView 用户登录
// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录
// @Param data body UserLoginRequest true "参数"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserLoginView(c *gin.Context) {
	var ulr UserLoginRequest
	err := c.ShouldBindJSON(&ulr)
	if err != nil {
		return
	}
	// if err != nil {
	// 	res.FailWithVaildError(err, &ulr, c)
	// 	return
	// }

	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", ulr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在", ulr.UserName)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CheckPwd(user.Password, ulr.Password) {
		global.Log.Warn("用户密码错误", ulr.UserName, ulr.Password)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserName: user.UserName,
		NickName: user.NickName,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成Token失败", c)
		return
	}

	global.DB.Model(&user).Update("lastLogin", time.Now())

	res.OKWithData(token, c)
}
