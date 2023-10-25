// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameImRegister = "im_registers"

// ImRegister mapped from table <im_registers>
type ImRegister struct {
	Account    string    `gorm:"column:account;type:char(255);primaryKey" json:"account"`
	Password   string    `gorm:"column:password;type:varchar(255)" json:"password"`
	Ex         string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
	UserID     string    `gorm:"column:user_id;type:varchar(255)" json:"user_id"`
	Platform   int32     `gorm:"column:platform;type:int(10);comment:平台，IOS:1、安卓:2、window:3、OSX:4、Web:5、小程序:6、Linux:7" json:"platform"`
	RegisterIP string    `gorm:"column:register_ip;type:varchar(255)" json:"register_ip"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

// TableName ImRegister's table name
func (*ImRegister) TableName() string {
	return TableNameImRegister
}
