// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameImPrivilegeUser = "im_privilege_user"

// ImPrivilegeUser mapped from table <im_privilege_user>
type ImPrivilegeUser struct {
	ID               int64          `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键id" json:"id"`
	UserID           string         `gorm:"column:user_id;type:varchar(64);comment:用户ID 对应着腾讯的id" json:"user_id"`
	Account          string         `gorm:"column:account;type:varchar(64);comment:帐户" json:"account"`
	Nickname         string         `gorm:"column:nickname;type:varchar(64);comment:昵称" json:"nickname"`
	FriendPrivileges int32          `gorm:"column:friend_privileges;type:int(11);comment:好有特权 1站点禁止加好友时不受限制 2禁用" json:"friend_privileges"`
	GroupPrivileges  int32          `gorm:"column:group_privileges;type:int(11);comment:群组特权 1不受限制创建群聊 2禁用" json:"group_privileges"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`
	CheckType        int32          `gorm:"column:check_type;type:tinyint(4);not null;comment:登录校验方式 0=不校验，1=IP校验，2=设备号校验" json:"check_type"`
}

// TableName ImPrivilegeUser's table name
func (*ImPrivilegeUser) TableName() string {
	return TableNameImPrivilegeUser
}