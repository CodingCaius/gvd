package log_stash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gvd_server/global"
	"io"
	"reflect"
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
	token    string
}

// 设置Action对象，并初始化一些字段
func NewAction(c *gin.Context) Action {
	ip := c.RemoteIP()
	addr := "局域网"
	action := Action{
		ip:   ip,
		addr: addr,
	}
	/* token := c.Request.Header.Get("token")
	jwyPayLoad := parseToken(token)
	if jwyPayLoad != nil {
		action.userID = jwyPayLoad.UserID
		action.userName = jwyPayLoad.UserName
	} */

	//不在这里解析token，现在这里拿到token
	token := c.Request.Header.Get("token")
	action.SetToken(token)
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
	//判断类型
	_type := reflect.TypeOf(value).Kind()
	switch _type {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		action.itemList = append(action.itemList, fmt.Sprintf("%s: %s", label, string(byteData)))
	default:
		action.itemList = append(action.itemList, fmt.Sprintf("%s: %v", label, value))
	}
}


func (action *Action) SetToken(token string) {
	action.token = token
}


// SetRequest 设置一组入参
func (action *Action) SetRequest(c *gin.Context) {
	// 请求头
	// 请求体
	// 请求路径，请求方法
	// 关于请求体的问题，拿了之后要还回去
	// 请求体读完之后就没了，为了日志和Login参数绑定时都能用，因此要还回去
	// 一定要在参数绑定之前调用
	method := c.Request.Method
	path := c.Request.URL.String()
	//对请求体进行 读取和重置
	byteData, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteData))

	action.itemList = append(action.itemList, fmt.Sprintf(
		`<div class="log_request">
  <div class="log_request_head">
    <span class="log_request_method">%s</span>
    <span class="log_request_path">%s</span>
  </div>
  <div class="log_request_body">
    <pre class="log_json_body">%s</pre>
  </div>
</div>`, method, path, string(byteData)))
}


// SetResponse 设置一组出参
func (action *Action) SetResponse(c *gin.Context) {
  c.Set("action", action)
}

func (action *Action) SetFlush() {
	action.level = action.model.Level
	action.save()
}


func (action *Action) save() {
	content := strings.Join(action.itemList, "\n")

	//留到这里解析token，顺便拿到用户id和用户名
	jwyPayLoad := parseToken(action.token)
	if jwyPayLoad != nil {
		action.userID = jwyPayLoad.UserID
		action.userName = jwyPayLoad.UserName
	}

	//这一步，model为空的话就创建一个并赋值，用于之后判断
	if action.model == nil {
		action.model = &LogModel{
			IP:       action.ip,
			Addr:     action.addr,
			Level:    action.level,
			Title:    action.title,
			Content:  content, //第一次的content
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
		Level: action.level,
		Title: action.title,
		//原来的content 加上 新的 content
		Content: action.model.Content + "\n" + content,
	})
}
