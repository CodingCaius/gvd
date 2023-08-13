package user_api

import (
	"encoding/json"
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
	"gvd_server/utils/pwd"
	"time"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	//bingding字段 定义字段在数据绑定（比如 HTTP 请求参数绑定）时的行为
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID" binding:"required"` //角色ID

}

// UserCreateView 处理 创建用户 的 API 请求
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户，只能管理员创建
// @Param data body UserCreateRequest true "参数"
// @Param token header string true "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	//将接收到的 JSON 数据解析到 cr 中
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//创建日志对象
	log := log_stash.NewAction(c)

	// 将cr转换为json格式
	byteData, _ := json.Marshal(cr)

	// 将cr转换为字符串，并设置到log中
	log.SetItem("创建参数", string(byteData))



	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err == nil {
		log.SetItem("userName", cr.UserName)
		log.Warn("创建用户错误，用户名已存在")
		res.FailWithMsg("用户名已存在", c)
		return

	}

	if cr.NickName == "" {
		//昵称如果没有，就自己拼接一个
		var maxID uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxID)
		cr.NickName = fmt.Sprintf("用户_%d", maxID+1)
		log.SetItem("自动生成昵称", cr.NickName)
	}
	err = global.DB.Create(&models.UserModel{
		UserName:  cr.UserName,
		Password:  pwd.HashPwd(cr.Password),
		NickName:  cr.NickName,
		IP:        c.RemoteIP(),
		RoleID:    cr.RoleID,
		LastLogin: time.Now(),
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户创建失败", c)
		log.SetItem("错误原因", err.Error())
		log.Error("用户创建失败")
		return
	}

	log.Info("用户创建成功")
	res.OKWithMsg("用户创建成功", c)
}
