package image_api

import (
	"gvd_server/global"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var ImageWhiteList = []string{
	"jpg",
	"png",
	"jpeg",
	"gif",
	"svg",
	"webp",
}

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

	//白名单判断
	if !InImageWhiteList(file.Filename, ImageWhiteList) {
		res.FailWithMsg("文件非法", c)
		return
	}

	//文件大小判断

	//重复文件判断

	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		global.Log.Errorf("%s 文件保存错误 %s", savePath, err)
		res.FailWithMsg("上传文件错误", c)
		return
	}

	res.OK("/"+savePath, "图片上传成功", c)
}

//白名单校验
// InImageWhiteList 判断一个图片是否在白名单中
func InImageWhiteList(fileName string, whiteList []string) bool {
	// 截取文件后缀
	_list := strings.Split(fileName, ".") // xxx  1.2 xxx.png xxx.PNG  xxx.png   xxx.1.2.png  xxxx.png.exe
	if len(_list) < 2 {
		return false
	}
	suffix := strings.ToLower(_list[len(_list)-1])
	for _, s := range whiteList {
		if suffix == s {
			return true
		}
	}
	return false
}