package filters

import (
    "github.com/astaxie/beego/context"
    "pybbs-go/models"
    "github.com/astaxie/beego"
    "regexp"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
    token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
    var user models.User
    if flag {
        flag, user = models.FindUserByToken(token)
    }
    return flag, user
}

var HasPermission = func(ctx *context.Context) {
    ok, user := IsLogin(ctx)
    if !ok {
        ctx.Redirect(302, "/login")
    } else {
        permissions := models.FindPermissionByUser(user.Id)
        url := ctx.Request.RequestURI
        beego.Debug("url: ", url)
        var flag = false
        for _, v := range permissions {
            if a, _ := regexp.MatchString(v.Url, url); a {
                flag = true
                break
            }
        }
        if !flag {
            ctx.WriteString("你没有权限访问这个页面")
        }
    }
}

var FilterUser = func(ctx *context.Context) {
    ok, _ := IsLogin(ctx)
    if !ok {
        ctx.Redirect(302, "/login")
    }
}
