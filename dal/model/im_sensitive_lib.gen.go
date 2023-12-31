// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameImSensitiveLib = "im_sensitive_lib"

// ImSensitiveLib mapped from table <im_sensitive_lib>
type ImSensitiveLib struct {
	ID         string    `gorm:"column:id;type:varchar(128);primaryKey;comment:库ID" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(255);comment:库名称" json:"name"`
	Describe   string    `gorm:"column:Describe;type:varchar(255);comment:库描述" json:"Describe"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间" json:"create_time"`
	Suggestion string    `gorm:"column:Suggestion;type:varchar(255);comment:库建议" json:"Suggestion"`
}

// TableName ImSensitiveLib's table name
func (*ImSensitiveLib) TableName() string {
	return TableNameImSensitiveLib
}
