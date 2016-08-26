package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Reply struct {
	Id      int `orm:"pk;auto"`
	Topic   *Topic `orm:"rel(fk)"`
	Content string `orm:"type(text)"`
	User    *User `orm:"rel(fk)"`
	Up      int `orm:"default(0)"`
	InTime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func FindReplyById(id int) Reply {
	o := orm.NewOrm()
	var reply Reply
	o.QueryTable(reply).RelatedSel("Topic").Filter("Id", id).One(&reply)
	return reply
}

func FindReplyByTopic(topic *Topic) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel().Filter("Topic", topic).OrderBy("-Up", "-InTime").All(&replies)
	return replies
}

func SaveReply(reply *Reply) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(reply)
	return id
}

func UpReply(reply *Reply) {
	o := orm.NewOrm()
	reply.Up = reply.Up + 1
	o.Update(reply, "Up")
}

func FindReplyByUser(user *User, limit int) []*Reply {
	o := orm.NewOrm()
	var reply Reply
	var replies []*Reply
	o.QueryTable(reply).RelatedSel("Topic", "User").Filter("User", user).OrderBy("-InTime").Limit(limit).All(&replies)
	return replies
}

func DeleteReplyByTopic(topic *Topic) {
	o := orm.NewOrm()
	var reply Reply
	var replies []Reply
	o.QueryTable(reply).Filter("Topic", topic).All(&replies)
	for _, reply := range replies {
		o.Delete(&reply)
	}
}

func DeleteReply(reply *Reply) {
	o := orm.NewOrm()
	o.Delete(reply)
}

func DeleteReplyByUser(user *User) {
    o := orm.NewOrm()
    o.Raw("delete form reply where user_id = ?", user.Id).Exec()
}