package routers

import (
	"pybbs-go/controllers"
	"github.com/astaxie/beego"
    "pybbs-go/filters"
)

func init() {
    beego.Router("/", &controllers.IndexController{}, "GET:Index")
    beego.Router("/login", &controllers.IndexController{}, "GET:LoginPage")
    beego.Router("/login", &controllers.IndexController{}, "POST:Login")
    beego.Router("/register", &controllers.IndexController{}, "GET:RegisterPage")
    beego.Router("/register", &controllers.IndexController{}, "POST:Register")
    beego.Router("/logout", &controllers.IndexController{}, "GET:Logout")
    beego.Router("/about", &controllers.IndexController{}, "GET:About")

    beego.InsertFilter("/topic/create", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/topic/create", &controllers.TopicController{}, "GET:Create")
    beego.Router("/topic/create", &controllers.TopicController{}, "POST:Save")
    beego.Router("/topic/:id([0-9]+)", &controllers.TopicController{}, "GET:Detail")
    beego.InsertFilter("/topic/edit/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "GET:Edit")
    beego.Router("/topic/edit/:id([0-9]+)", &controllers.TopicController{}, "POST:Update")
    beego.InsertFilter("/topic/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/topic/delete/:id([0-9]+)", &controllers.TopicController{}, "GET:Delete")

    beego.InsertFilter("/reply/save", beego.BeforeRouter, filters.FilterUser)
    beego.Router("/reply/save", &controllers.ReplyController{}, "POST:Save")
	beego.InsertFilter("/reply/up", beego.BeforeRouter, filters.FilterUser)
    beego.Router("/reply/up", &controllers.ReplyController{}, "GET:Up")
	beego.InsertFilter("/reply/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/reply/delete/:id([0-9]+)", &controllers.ReplyController{}, "GET:Delete")

    beego.Router("/user/:username", &controllers.UserController{}, "GET:Detail")
    beego.Router("/user/setting", &controllers.UserController{}, "GET:ToSetting")
    beego.InsertFilter("/user/setting", beego.BeforeRouter, filters.FilterUser)
    beego.Router("/user/setting", &controllers.UserController{}, "POST:Setting")
    beego.InsertFilter("/user/updatepwd", beego.BeforeRouter, filters.FilterUser)
    beego.Router("/user/updatepwd", &controllers.UserController{}, "POST:UpdatePwd")
    beego.InsertFilter("/user/updateavatar", beego.BeforeRouter, filters.FilterUser)
    beego.Router("/user/updateavatar", &controllers.UserController{}, "POST:UpdateAvatar")
    beego.InsertFilter("/user/list", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/user/list", &controllers.UserController{}, "GET:List")
    beego.InsertFilter("/user/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "GET:Delete")
    beego.InsertFilter("/user/edit/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "GET:Edit")
    beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "POST:Update")

    beego.InsertFilter("/role/list", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/role/list", &controllers.RoleController{}, "GET:List")
    beego.InsertFilter("/role/add", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/role/add", &controllers.RoleController{}, "GET:Add")
    beego.Router("/role/add", &controllers.RoleController{}, "Post:Save")
    beego.InsertFilter("/role/edit/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "GET:Edit")
    beego.Router("/role/edit/:id([0-9]+)", &controllers.RoleController{}, "Post:Update")
    beego.InsertFilter("/role/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/role/delete/:id([0-9]+)", &controllers.RoleController{}, "GET:Delete")

    beego.InsertFilter("/permission/list", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/permission/list", &controllers.PermissionController{}, "GET:List")
    beego.InsertFilter("/permission/add", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/permission/add", &controllers.PermissionController{}, "GET:Add")
    beego.Router("/permission/add", &controllers.PermissionController{}, "Post:Save")
    beego.InsertFilter("/permission/edit", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/permission/edit/:id([0-9]+)", &controllers.PermissionController{}, "GET:Edit")
    beego.Router("/permission/edit/:id([0-9]+)", &controllers.PermissionController{}, "Post:Update")
    beego.InsertFilter("/permission/delete/:id([0-9]+)", beego.BeforeRouter, filters.HasPermission)
    beego.Router("/permission/delete/:id([0-9]+)", &controllers.PermissionController{}, "GET:Delete")

}
