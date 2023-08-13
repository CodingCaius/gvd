package log_stash

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"strings"
)

type Action struct {
	ip       string
	addr     string
	userName string
	userID   uint
	level    Level
	title    string
	itemList []string
}

//设置Action对象，并初始化一些字段
func NewAction(c *gin.Context) Action {
	ip := c.RemoteIP()
	addr := "局域网"
	action := Action{
		ip:   ip,
		addr: addr,
	}
	token := c.Request.Header.Get("token")
	jwyPayLoad := parseToken(token)
	if jwyPayLoad != nil {
		action.userID = jwyPayLoad.UserID
		action.userName = jwyPayLoad.UserName
	}
	return action
}

// Info Warn Error三个属于中断式函数
func (action *Action) Info(title string) {
	action.level = Info
	action.title = title
	action.save()
}
func (action *Action) Warn(title string) {
	action.level = Warning
	action.title = title
	action.save()
}
func (action *Action) Error(title string) {
	action.level = Error
	action.title = title
	action.save()
}

//SetItem 方法用于向 Action 结构体的 itemList 字段中添加附加项
func (action *Action) SetItem(label string, value any) {
	action.itemList = append(action.itemList, fmt.Sprintf("%s: %v", label, value))
}

func (action *Action) save() {
	content := strings.Join(action.itemList, "\n")
	global.DB.Create(&LogModel{
		IP:       action.ip,
		Addr:     action.addr,
		Level:    action.level,
		Title:    action.title,
		Content:  content,
		UserID:   action.userID,
		UserName: action.userName,
		Type:     ActionType,
	})
}