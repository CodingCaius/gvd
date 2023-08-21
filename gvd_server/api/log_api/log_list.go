package log_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/list"
	"gvd_server/service/common/res"
)

type LogListRequest struct {
	models.Pagination
	Level log_stash.Level   `json:"level" form:"level"` // 日志查询的等级
	Type  log_stash.LogType `json:"type" form:"type"`   // 日志的类型   1 登录日志  2 操作日志  3 运行日志
}

// LogListView 日志列表 可根据type level 搜索日志
// @Tags 日志管理
// @Summary 日志列表
// @Description 日志列表
// @Param data query LogListRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[log_stash.LogModel]}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogListRequest
	c.ShouldBindQuery(&cr)
	_list, count, _ := list.QueryList(log_stash.LogModel{
		Type:  cr.Type,
		Level: cr.Level,
	}, list.Option{
		Pagination: cr.Pagination,
	})
	res.OKWithList(_list, count, c)
}