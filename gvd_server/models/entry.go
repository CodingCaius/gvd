package models //表结构

import "time"

/*
JSON 标签的作用是指定在将结构体实例转换为 JSON 字符串（序列化）时，使用指定的字段名称来表示结构体字段。同样，在从 JSON 字符串中解析数据（反序列化）时，也会根据 JSON 标签来找到对应的字段并赋值
*/


type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt" json:"updatedAt"`
}