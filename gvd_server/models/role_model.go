package models

type RoleModel struct {
	Model
	Title string `gorm:"size:16;not null;comment:角色的名称" json:"title"` //角色的名称
	Pwd string `gorm:"size:64" json:"pwd"` //角色密码
	IsSystem bool `gorm:"column:isSystem" json:"isSystem"` //是否是系统角色
	//角色拥有的文档
	DocsList []DocModel `gorm:"many2many:role_doc_models;joinForeignKey:RoleID;JoinReferences:DocID" json:"-"` // joinForeignKey:RoleID 自定义外键
	UserList []UserModel `foreignKey:"RoleID" json:"-"`
}