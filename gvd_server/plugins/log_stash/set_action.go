package log_stash

import (
	"fmt"
	"gvd_server/global"
	"strings"

	"github.com/gin-gonic/gin"
)

type Action struct {
	ip       string
	addr     string
	userName string
	userID   uint
	level    Level
	title    string
	itemList []string
	model    *LogModel //创建之后赋值给它，用于后期更新
}

// 设置Action对象，并初始化一些字段
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

// SetItem 方法用于向 Action 结构体的 itemList 字段中添加附加项
func (action *Action) SetItem(label string, value any) {
	action.itemList = append(action.itemList, fmt.Sprintf("%s: %v", label, value))
}

func (action *Action) save() {
	content := strings.Join(action.itemList, "\n")
	//这一步，model为空的话就创建一个并赋值，用于之后判断
	if action.model == nil {
		action.model = &LogModel{
			IP:       action.ip,
			Addr:     action.addr,
			Level:    action.level,
			Title:    action.title,
			Content:  content,  //第一次的content
			UserID:   action.userID,
			UserName: action.userName,
			Type:     ActionType,
		}
		global.DB.Create(action.model)
		// 如果不对content进行置空，那么content会重复
		action.itemList = []string{}
		return
	}
	//如果model不为空，说明是同一个log，就执行更新操作
	global.DB.Model(action.model).Updates(LogModel{
		Level:    action.level,
		Title:    action.title,
		//原来的content 加上 新的 content
		Content:  action.model.Content + "\n" + content,
	})
}
