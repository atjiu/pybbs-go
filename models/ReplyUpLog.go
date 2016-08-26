package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type ReplyUpLog struct {
	Id     int `orm:"pk;auto"`
	User   *User `orm:"rel(fk)"`
	Reply  *Reply `orm:"rel(fk)"`
	InTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func SaveReplyUpLog(replyUpLog *ReplyUpLog) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(replyUpLog)
	return id
}

func FindReplyUpLogByUserAndReply(user *User, reply *Reply) ReplyUpLog {
	o := orm.NewOrm()
	var replyUpLog ReplyUpLog
	o.QueryTable(replyUpLog).Filter("User", user).Filter("Reply", reply).One(&replyUpLog)
	return replyUpLog
}
