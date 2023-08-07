package user_api

import (
	"gvd_server/global"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"gvd_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UserUpdatePasswordRequest struct {
	OldPwd   string `json:"oldPwd" binding:"required" label:"之前的密码"`
	Password string `json:"password" binding:"required" label:"新密码"`
}

func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UserUpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithValidError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	//从声明中获取userID，并在数据库中查询
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMsg("原密码错误", c)
		return
	}

	//更新密码
	hashPwd := pwd.HashPwd(cr.Password)
	global.DB.Model(user).Update("password", hashPwd)

	res.OKWithMsg("用户密码修改成功", c)
}
