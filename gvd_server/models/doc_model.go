package models

import "gvd_server/global"

type DocModel struct {
	Model
	Title     string `gorm:"comment:文档标题" json:"title"`
	Content   string `gorm:"comment:文档内容" json:"-"`
	DiggCount int    `gorm:"comment:点赞量" json:"diggcount"`
	LookCount int    `gorm:"comment:浏览量" json:"lookCount"`
	Key       string `gorm:"comment:key;not null;unique" json:"key"`
	ParentID  *uint  `gorm:"comment:父文档id column:parent_id" json:"parentID"`
	//通过 ParentID 字段与父文档DocModel 关联起来的
	ParentModel *DocModel   `gorm:"foreignKey:ParentID" json:"-"` //父文档
	Child       []*DocModel `gorm:"foreignKey:ParentID" json:"-"` //子孙文档
	FreeContent string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`

	//收藏该文档的用户列表
	/*连接表 user_coll_doc_models：
	DocID: 外键，关联到文档表的 DocID 字段。
	UserID: 外键，关联到用户表的 UserID 字段。*/
	UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;joinForeignKey:DocID;JoinReferences:UserID" json:"-"`
}

// FindAllParentDocList 找一个文档的所有父文档
func FindAllParentDocList(doc DocModel, docList *[]DocModel) {
	// 不管谁来，先把自己放进去
	*docList = append(*docList, doc)
	if doc.ParentID != nil {
		// 说明有父文档
		var parentDoc DocModel
		global.DB.Take(&parentDoc, *doc.ParentID)
		FindAllParentDocList(parentDoc, docList)
	}
}
