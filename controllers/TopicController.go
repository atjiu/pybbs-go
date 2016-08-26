package controllers

import (
	"github.com/astaxie/beego"
	"pybbs-go/models"
	"strconv"
	"pybbs-go/filters"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Create() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
	c.Data["PageTitle"] = "发布话题"
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.tpl"
	c.TplName = "topic/create.tpl"
}

func (c *TopicController) Save() {
	flash := beego.NewFlash()
	title, content, sid := c.Input().Get("title"), c.Input().Get("content"), c.Input().Get("sid")
	if len(title) == 0 || len(title) > 120 {
		flash.Error("话题标题不能为空且不能超过120个字符")
		flash.Store(&c.Controller)
		c.Redirect("/topic/create", 302)
	} else if len(sid) == 0 {
		flash.Error("请选择话题版块")
		flash.Store(&c.Controller)
		c.Redirect("/topic/create", 302)
	} else {
		s, _ := strconv.Atoi(sid)
		section := models.Section{Id: s}
		_, user := filters.IsLogin(c.Ctx)
		topic := models.Topic{Title: title, Content: content, Section: &section, User: &user}
		id := models.SaveTopic(&topic)
		c.Redirect("/topic/" + strconv.FormatInt(id, 10), 302)
	}
}

func (c *TopicController) Detail() {
	id := c.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	if tid > 0 {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
		topic := models.FindTopicById(tid)
		models.IncrView(&topic)//查看+1
		c.Data["PageTitle"] = topic.Title
		c.Data["Topic"] = topic
		c.Data["Replies"] = models.FindReplyByTopic(&topic)
		c.Layout = "layout/layout.tpl"
		c.TplName = "topic/detail.tpl"
	} else {
		c.Ctx.WriteString("话题不存在")
	}
}

func (c *TopicController) Edit() {
	beego.ReadFromRequest(&c.Controller)
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		topic := models.FindTopicById(id)
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Controller.Ctx)
		c.Data["PageTitle"] = "编辑话题"
		c.Data["Sections"] = models.FindAllSection()
		c.Data["Topic"] = topic
		c.Layout = "layout/layout.tpl"
		c.TplName = "topic/edit.tpl"
	} else {
		c.Ctx.WriteString("话题不存在")
	}
}

func (c *TopicController) Update() {
	flash := beego.NewFlash()
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title, content, sid := c.Input().Get("title"), c.Input().Get("content"), c.Input().Get("sid")
	if len(title) == 0 || len(title) > 120 {
		flash.Error("话题标题不能为空且不能超过120个字符")
		flash.Store(&c.Controller)
		c.Redirect("/topic/edit/" + strconv.Itoa(id), 302)
	} else if len(sid) == 0 {
		flash.Error("请选择话题版块")
		flash.Store(&c.Controller)
		c.Redirect("/topic/edit/" + strconv.Itoa(id), 302)
	} else {
		s, _ := strconv.Atoi(sid)
		section := models.Section{Id: s}
		topic := models.FindTopicById(id)
		topic.Title = title
		topic.Content = content
		topic.Section = &section
		models.UpdateTopic(&topic)
		c.Redirect("/topic/" + strconv.Itoa(id), 302)
	}
}

func (c *TopicController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if id > 0 {
		topic := models.FindTopicById(id)
		models.DeleteTopic(&topic)
		models.DeleteReplyByTopic(&topic)
		c.Redirect("/", 302)
	} else {
		c.Ctx.WriteString("话题不存在")
	}
}