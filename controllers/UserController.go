package controllers

import (
	"github.com/astaxie/beego"
	"pybbs-go/models"
	"pybbs-go/filters"
	"regexp"
    "strconv"
	"net/http"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Detail() {
	username := c.Ctx.Input.Param(":username")
	ok, user := models.FindUserByUserName(username)
	if ok {
		c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
		c.Data["PageTitle"] = "个人主页"
		c.Data["CurrentUserInfo"] = user
		c.Data["Topics"] = models.FindTopicByUser(&user, 7)
		c.Data["Replies"] = models.FindReplyByUser(&user, 7)
	}
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/detail.tpl"
}

func (c *UserController) ToSetting() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
	c.Data["PageTitle"] = "用户设置"
	c.Layout = "layout/layout.tpl"
	c.TplName = "user/setting.tpl"
}

func (c *UserController) Setting() {
	flash := beego.NewFlash()
	email, url, signature := c.Input().Get("email"), c.Input().Get("url"), c.Input().Get("signature")
	if len(email) > 0 {
		ok, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|_|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", email)
		if !ok {
			flash.Error("请输入正确的邮箱地址");
			flash.Store(&c.Controller)
			c.Redirect("/user/setting", 302)
			return
		}
	}
	if len(signature) > 1000 {
		flash.Error("个人签名长度不能超过1000字符");
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	_, user := filters.IsLogin(c.Ctx)
	user.Email = email
	user.Url = url
	user.Signature = signature
	models.UpdateUser(&user)
	flash.Success("更新资料成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *UserController) UpdatePwd() {
	flash := beego.NewFlash()
	oldpwd, newpwd := c.Input().Get("oldpwd"), c.Input().Get("newpwd")
	_, user := filters.IsLogin(c.Ctx)
	if user.Password != oldpwd {
		flash.Error("旧密码不正确")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	if len(newpwd) == 0 {
		flash.Error("新密码都不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	}
	user.Password = newpwd
	models.UpdateUser(&user)
	flash.Success("密码修改成功")
	flash.Store(&c.Controller)
	c.Redirect("/user/setting", 302)
}

func (c *UserController) UpdateAvatar() {
	flash := beego.NewFlash()
	f, h, err := c.GetFile("avatar")
	if err == http.ErrMissingFile {
		flash.Error("请选择文件")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	}
	defer f.Close()
	if err != nil {
		flash.Error("上传失败")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
		return
	} else {
		c.SaveToFile("avatar", "static/upload/avatar/" + h.Filename)
		_, user := filters.IsLogin(c.Ctx)
		user.Avatar = "/static/upload/avatar/" + h.Filename
		models.UpdateUser(&user)
		flash.Success("上传成功")
		flash.Store(&c.Controller)
		c.Redirect("/user/setting", 302)
	}
}

func (c *UserController) List() {
    c.Data["PageTitle"] = "用户列表"
    c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
    p, _ := strconv.Atoi(c.Ctx.Input.Query("p"))
    if p == 0 {
        p = 1
    }
    size, _ := beego.AppConfig.Int("page.size")
    c.Data["Page"] = models.PageUser(p, size)
    c.Layout = "layout/layout.tpl"
    c.TplName = "user/list.tpl"
}

func (c *UserController) Edit() {
    c.Data["PageTitle"] = "配置角色"
    c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    if id > 0 {
        ok, user := models.FindUserById(id)
        if ok {
            c.Data["User"] = user
            c.Data["Roles"] = models.FindRoles()
            c.Data["UserRoles"] = models.FindUserRolesByUserId(id)
            c.Layout = "layout/layout.tpl"
            c.TplName = "user/edit.tpl"
        } else {
            c.Ctx.WriteString("用户不存在")
        }
    } else {
        c.Ctx.WriteString("用户不存在")
    }
}

func (c *UserController) Update() {
    c.Data["PageTitle"] = "配置角色"
    c.Data["IsLogin"], c.Data["UserInfo"] = filters.IsLogin(c.Ctx)
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    roleIds := c.GetStrings("roleIds")
    if id > 0 {
        models.DeleteUserRolesByUserId(id)
        for _, v := range roleIds {
            roleId, _ := strconv.Atoi(v)
            models.SaveUserRole(id, roleId)
        }
        c.Redirect("/user/list", 302)
    } else {
        c.Ctx.WriteString("用户不存在")
    }
}

func (c *UserController) Delete() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    if id > 0 {
        ok, user := models.FindUserById(id)
        if ok {
            models.DeleteTopicByUser(&user)
            models.DeleteReplyByUser(&user)
            models.DeleteUser(&user)
        }
        c.Redirect("/user/list", 302)
    } else {
        c.Ctx.WriteString("用户不存在")
    }
}