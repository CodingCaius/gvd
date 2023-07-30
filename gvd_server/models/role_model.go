package models

type RoleModel struct {
	Model
	//RoleID uint    `gorm:"colmun:rolID;comment:角色id" json:"roleID"`
	Title string `gorm:"size:16;not null;comment:角色的名称" json:"title"` //角色的名称
	Pwd string `gorm:"size:64" json:"-"` //角色密码
	IsSystem bool `gorm:"column:isSystem" json:"isSystem"` //是否是系统角色
	//角色拥有的文档
	DocList []DocModel `gorm:"many2many:role_doc_models;joinForeignKey:RoleID;JoinReferences:DocID" json:"-"`
}