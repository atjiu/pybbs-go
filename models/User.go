package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"pybbs-go/utils"
	"github.com/casbin/casbin"
)

type User struct {
	Id        int `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
	Token     string `orm:"unique"`
	Avatar    string
	Email     string `orm:"null"`
	Url       string `orm:"null"`
	Signature string `orm:"null;size(1000)"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
	Roles     []*Role `orm:"rel(m2m)"`
}

var Enforcer *casbin.Enforcer = nil

func getURL(name string) string {
	permissions := FindPermissions()
	for _, permission := range permissions {
		if name == strconv.Itoa(permission.Id) {
			return permission.Url
		}
	}
	return ""
}

func getURLFunc(args ...interface{}) (interface{}, error) {
	name := args[0].(string)

	return (string)(getURL(name)), nil
}

func Init() {
	Enforcer = &casbin.Enforcer{}
	Enforcer.InitWithFile("rbac_model.conf", "")
	Enforcer.AddFunction("getURL", getURLFunc)

	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select user_id, role_id from user_roles").Values(&res, "user_id", "role_id")
	for _, param := range res {
		Enforcer.AddRoleForUser(param["user_id"].(string), param["role_id"].(string))
	}

	o = orm.NewOrm()
	o.Raw("select role_id, permission_id from role_permissions").Values(&res, "role_id", "permission_id")
	for _, param := range res {
		Enforcer.AddPermissionForUser(param["role_id"].(string), param["permission_id"].(string))
	}
}

func FindUserById(id int) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func Login(username string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

func PageUser(p int, size int) utils.Page {
	o := orm.NewOrm()
	var user User
	var list []User
	qs := o.QueryTable(user)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-InTime").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func FindPermissionByUserIdAndPermissionName(userId int, name string) bool {
	permissions := FindPermissions()
	for _, permission := range permissions {
		if name == permission.Name {
			return Enforcer.Enforce(strconv.Itoa(userId), permission.Url)
		}
	}

	return false
}

func DeleteUser(user *User) {
	Enforcer.DeleteUser(strconv.Itoa(user.Id))

	o := orm.NewOrm()
	o.Delete(user)
}

func DeleteUserRolesByUserId(user_id int) {
	Enforcer.DeleteRolesForUser(strconv.Itoa(user_id))

	o := orm.NewOrm()
	o.Raw("delete from user_roles where user_id = ?", user_id).Exec()
}

func SaveUserRole(user_id int, role_id int) {
	Enforcer.AddRoleForUser(strconv.Itoa(user_id), strconv.Itoa(role_id))

	o := orm.NewOrm()
	o.Raw("insert into user_roles (user_id, role_id) values (?, ?)", user_id, role_id).Exec()
}

func FindUserRolesByUserId(user_id int) []orm.Params {
	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select id, user_id, role_id from user_roles where user_id = ?", user_id).Values(&res, "id", "user_id", "role_id")
	return res
}