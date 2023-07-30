package models

type UserPwdDocModel struct {
	Model
	UserID uint `grom:"column:userID" json:"userID"`
	DocID  uint `grom:"column:docID" json:"docID"`
}
