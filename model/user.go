package model

import "time"

const (
	SEX_WOMEN="W"
	SEX_MEN="M"
	SEX_UNKNOW="U"
)
type User struct {
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`//用户ID
	Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // 加密密码
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`   // 昵称
	Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`   //加盐随机字符串6
	Online     int	`xorm:"int(10)" form:"online" json:"online"`   //是否在线
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   //前端鉴权因子
	Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`   // 简介
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 创建时间
}