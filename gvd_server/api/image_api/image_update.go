package image_api

import (
	"gvd_server/global"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"path"

	"github.com/gin-gonic/gin"
)

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片
// @Param token header string true "token"
// @Param image formData file true "文件上传"
// @Router /api/image [post]
// @Accept multipart/formdata
// @Produce json
// @Success 200 {object} res.Response{}
func (ImageApi) ImageUploadView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		res.FailWithMsg("图片参数错误", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)

	savePath := path.Join("updates", claims.NickName, file.Filename)

	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		global.Log.Errorf("%s 文件保存错误 %s", savePath, err)
		res.FailWithMsg("上传文件错误", c)
		return
	}

	res.OK("/" + savePath, "图片上传成功", c)
}
