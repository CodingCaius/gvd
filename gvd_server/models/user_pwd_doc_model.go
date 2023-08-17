package models

//自定义连接表

//用户文档密码模型
type UserPwdDocModel struct {
	Model
	UserID uint `grom:"column:user_id" json:"userID"`
	DocID  uint `grom:"column:doc_id" json:"docID"`
}
