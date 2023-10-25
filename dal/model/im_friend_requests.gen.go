// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameImFriendRequest = "im_friend_requests"

// ImFriendRequest mapped from table <im_friend_requests>
type ImFriendRequest struct {
	FromUserID    string    `gorm:"column:from_user_id;type:varchar(64);primaryKey;comment:申请者 UserID" json:"from_user_id"`
	ToUserID      string    `gorm:"column:to_user_id;type:varchar(64);primaryKey;comment:被申请者 UserID" json:"to_user_id"`
	HandleResult  int32     `gorm:"column:handle_result;type:int(11);comment:处理结果" json:"handle_result"`
	ReqMsg        string    `gorm:"column:req_msg;type:varchar(255);comment:处理结果" json:"req_msg"`
	CreateTime    time.Time `gorm:"column:create_time;type:datetime(3);comment:申请时间" json:"create_time"`
	HandlerUserID string    `gorm:"column:handler_user_id;type:varchar(64);comment:处理者的UserID" json:"handler_user_id"`
	HandleMsg     string    `gorm:"column:handle_msg;type:varchar(255);comment:处理信息" json:"handle_msg"`
	HandleTime    time.Time `gorm:"column:handle_time;type:datetime(3);comment:处理时间" json:"handle_time"`
	Ex            string    `gorm:"column:ex;type:varchar(1024);comment:拓展字段" json:"ex"`
	ComeFrom      int32     `gorm:"column:come_from;type:int(10);comment:来源，1-ID,2-扫码,3-名片，4-群聊，5-其他" json:"come_from"`
}

// TableName ImFriendRequest's table name
func (*ImFriendRequest) TableName() string {
	return TableNameImFriendRequest
}