// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameImHandleLog = "im_handle_log"

// ImHandleLog mapped from table <im_handle_log>
type ImHandleLog struct {
	ID         int64          `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键id" json:"id"`
	Account    string         `gorm:"column:account;type:varchar(64);comment:账号" json:"account"`
	Nickname   string         `gorm:"column:nickname;type:varchar(255);comment:账号昵称" json:"nickname"`
	HandleType string         `gorm:"column:handle_type;type:varchar(64);comment:操作类型" json:"handle_type"`
	HandleDesc string         `gorm:"column:handle_desc;type:text;comment:操作描叙" json:"handle_desc"`
	RequstBody string         `gorm:"column:requst_body;type:text;comment:请求参数" json:"requst_body"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`
}

// TableName ImHandleLog's table name
func (*ImHandleLog) TableName() string {
	return TableNameImHandleLog
}