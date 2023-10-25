// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameImUser = "im_users"

// ImUser mapped from table <im_users>
type ImUser struct {
	UserID           string    `gorm:"column:user_id;type:varchar(64);primaryKey;comment:用户ID" json:"user_id"`
	Account          string    `gorm:"column:account;type:varchar(64);comment:帐户" json:"account"`
	UserSig          string    `gorm:"column:user_sig;type:varchar(255);comment:私钥" json:"user_sig"`
	PhoneNumber      string    `gorm:"column:phone_number;type:varchar(32);comment:联系电话" json:"phone_number"`
	AreaCode         string    `gorm:"column:area_code;type:varchar(8);comment:区号" json:"area_code"`
	Email            string    `gorm:"column:email;type:varchar(64);comment:email" json:"email"`
	Nickname         string    `gorm:"column:nickname;type:varchar(64);comment:昵称" json:"nickname"`
	FaceURL          string    `gorm:"column:face_url;type:varchar(255);comment:头像" json:"face_url"`
	Gender           int32     `gorm:"column:gender;type:tinyint(2);comment:性别 0=女性，1=男性" json:"gender"`
	BirthTime        time.Time `gorm:"column:birth_time;type:datetime(3);comment:生日" json:"birth_time"`
	Level            int32     `gorm:"column:level;type:int(11);default:1;comment:等级" json:"level"`
	AllowVibration   int32     `gorm:"column:allow_vibration;type:tinyint(4);default:1;comment:允许振动通知" json:"allow_vibration"`
	AllowBeep        int32     `gorm:"column:allow_beep;type:tinyint(4);default:1;comment:允许声音通知" json:"allow_beep"`
	AllowAddFriend   int32     `gorm:"column:allow_add_friend;type:tinyint(4);default:1;comment:允许添加好友" json:"allow_add_friend"`
	GlobalRecvMsgOpt int32     `gorm:"column:global_recv_msg_opt;type:tinyint(4);comment:允许通知信息" json:"global_recv_msg_opt"`
	Ex               string    `gorm:"column:ex;type:varchar(1024);comment:备注" json:"ex"`
	AppMangerLevel   int32     `gorm:"column:app_manger_level;type:tinyint(4);comment:管理等级" json:"app_manger_level"`
	Motto            string    `gorm:"column:motto;type:varchar(1024);comment:签名" json:"motto"`
	LastLoginIP      string    `gorm:"column:last_login_ip;type:varchar(120);comment:最近登录IP" json:"last_login_ip"`
	LastLoginTime    time.Time `gorm:"column:last_login_time;type:datetime(3);comment:最近登录时间" json:"last_login_time"`
	Status           int32     `gorm:"column:status;type:tinyint(2);default:1;comment:状态 0=冻结，1=正常" json:"status"`
	GroupMute        int32     `gorm:"column:group_mute;type:tinyint(4);comment:群聊禁言 0=未禁言，1=禁言中" json:"group_mute"`
	FriendMute       int32     `gorm:"column:friend_mute;type:tinyint(4);comment:私聊禁言 0=未禁言，1=禁言中" json:"friend_mute"`
	IsPrivilegedUser int32     `gorm:"column:is_privileged_user;type:tinyint(2);comment:是否是特权用户, 0=不是，1=是" json:"is_privileged_user"`
	AllowCreateGroup int32     `gorm:"column:allow_create_group;type:tinyint(2);default:1;comment:允许创建群聊 0=不允许，1=允许" json:"allow_create_group"`
	CreateBy         string    `gorm:"column:create_by;type:varchar(128);not null;comment:创建人" json:"create_by"`
	UpdateBy         string    `gorm:"column:update_by;type:varchar(128);not null;comment:更新人" json:"update_by"`
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
	FreezeBy         string    `gorm:"column:freeze_by;type:varchar(128);not null;comment:冻结操作人" json:"freeze_by"`
	FreezeAt         time.Time `gorm:"column:freeze_at;type:datetime;comment:冻结时间" json:"freeze_at"`
	FreezeRemark     string    `gorm:"column:freeze_remark;type:varchar(500);not null;comment:冻结原因" json:"freeze_remark"`
	InIPWhitelist    int32     `gorm:"column:in_ip_whitelist;type:tinyint(2);comment:是否在IP白名单, 0=不是，1=是" json:"in_ip_whitelist"`
	FriendCount      int32     `gorm:"column:friend_count;type:int(8);default:1;comment:好友数" json:"friend_count"`
	GroupCount       int32     `gorm:"column:group_count;type:int(8);default:1;comment:群聊数" json:"group_count"`
	FreezeEndtime    time.Time `gorm:"column:freeze_endtime;type:datetime;comment:冻结截止时间" json:"freeze_endtime"`
}

// TableName ImUser's table name
func (*ImUser) TableName() string {
	return TableNameImUser
}
