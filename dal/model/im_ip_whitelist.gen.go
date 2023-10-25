// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameImIPWhitelist = "im_ip_whitelist"

// ImIPWhitelist mapped from table <im_ip_whitelist>
type ImIPWhitelist struct {
	ID        int64     `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:ID" json:"id"`
	UserID    string    `gorm:"column:user_id;type:varchar(64);not null;comment:用户ID" json:"user_id"`
	Account   string    `gorm:"column:account;type:varchar(64);comment:帐户ID" json:"account"`
	Nickname  string    `gorm:"column:nickname;type:varchar(120);comment:昵称" json:"nickname"`
	Role      int32     `gorm:"column:role;type:tinyint(4);default:1;comment:身份 1=普通用户、2=特权用户" json:"role"`
	Remark    string    `gorm:"column:remark;type:varchar(1024);comment:原因" json:"remark"`
	CreateBy  string    `gorm:"column:create_by;type:varchar(128);not null;comment:创建人" json:"create_by"`
	UpdateBy  string    `gorm:"column:update_by;type:varchar(128);not null;comment:更新人" json:"update_by"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
}

// TableName ImIPWhitelist's table name
func (*ImIPWhitelist) TableName() string {
	return TableNameImIPWhitelist
}