package routers

import (
	"fmt"
	"pybbs-go/controllers"
	"pybbs-go/filters"

	"github.com/astaxie/beego"
)

func setPatternsNeedLogin(patterns ...string) {
	for _, pa := range patterns {
		beego.Debug(fmt.Sprintf("set url pattern: %s need login\n", pa))
		beego.InsertFilter(pa, beego.BeforeRouter, filters.HasPermission)
	}
}

func setupIndex() {
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/login", &controllers.IndexController{}, "GET:LoginPage")
	beego.Router("/login", &controllers.IndexController{}, "POST:Login")
	beego.Router("/register", &controllers.IndexController{}, "GET:RegisterPage")
	beego.Router("/register", &controllers.IndexController{}, "POST:Register")
	beego.Router("/logout", &controllers.IndexController{}, "GET:Logout")
	beego.Router("/about", &controllers.IndexController{}, "GET:About")
}

func setupTopic() {
	beego.Router("/topic/create", &controllers.TopicController{}, "GET:Create")
	beego.Router("/topic/create", &controllers.TopicController{}, "POST:Save")
	beego.Router("/topic/:id([0-9]+)", &controllers.TopicController{}, "GET:Detail")
	beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "GET:Edit")
	beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "POST:Update")
	beego.Router("/topic/delete/:id([0-9]+)", &controllers.TopicController{}, "GET:Delete")
	setPatternsNeedLogin("/topic/create", "/topic/edit/:id([0-9]+)", "/topic/delete/:id([0-9]+)")
}

func setupReply() {
	beego.Router("/reply/save", &controllers.ReplyController{}, "POST:Save")
	beego.Router("/reply/up", &controllers.ReplyController{}, "GET:Up")
	beego.Router("/reply/delete/:id([0-9]+)", &controllers.ReplyController{}, "GET:Delete")
	setPatternsNeedLogin("/reply/save", "/reply/up", "/reply/delete/:id([0-9]+)")
}

func setupUser() {
	beego.Router("/user/:username", &controllers.UserController{}, "GET:Detail")
	beego.Router("/user/setting", &controllers.UserController{}, "GET:ToSetting")
	beego.Router("/user/setting", &controllers.UserController{}, "POST:Setting")
	beego.Router("/user/updatepwd", &controllers.UserController{}, "POST:UpdatePwd")
	beego.Router("/user/updateavatar", &controllers.UserController{}, "POST:UpdateAvatar")
	beego.Router("/user/list", &controllers.UserController{}, "GET:List")
	beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "GET:Delete")
	beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "GET:Edit")
	beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "POST:Update")
	setPatternsNeedLogin("/user/edit/:id([0-9]+)", "/user/delete/:id([0-9]+)", "/user/list", "/user/updateavatar", "/user/setting", "/user/updatepwd")
}

func setupRole() {
	beego.Router("/role/list", &controllers.RoleController{}, "GET:List")
	beego.Router("/role/add", &controllers.RoleController{}, "GET:Add")
	beego.Router("/role/add", &controllers.RoleController{}, "Post:Save")
	beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "GET:Edit")
	beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "Post:Update")
	beego.Router("/role/delete/:id([0-9]+)", &controllers.RoleController{}, "GET:Delete")
	setPatternsNeedLogin("/role/list", "/role/add", "/role/edit/:id([0-9]+)", "/role/delete/:id([0-9]+)")
}

func setupPermission() {
	beego.Router("/permission/list", &controllers.PermissionController{}, "GET:List")
	beego.Router("/permission/add", &controllers.PermissionController{}, "GET:Add")
	beego.Router("/permission/add", &controllers.PermissionController{}, "Post:Save")
	beego.Router("/permission/edit/:id([0-9]+)", &controllers.PermissionController{}, "GET:Edit")
	beego.Router("/permission/edit/:id([0-9]+)", &controllers.PermissionController{}, "Post:Update")
	beego.Router("/permission/delete/:id([0-9]+)", &controllers.PermissionController{}, "GET:Delete")
	setPatternsNeedLogin("/permission/list", "/permission/add", "/permission/edit", "/permission/delete/:id([0-9]+)")
}

func init() {
	setupIndex()
	setupTopic()
	setupReply()
	setupRole()
	setupUser()
	setupPermission()
}
