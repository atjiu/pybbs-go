package models

import (
	"time"
	"github.com/astaxie/beego/orm"
    "strconv"
    "pybbs-go/utils"
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

func FindPermissionByUser(id int) []*Permission {
	o := orm.NewOrm()
	var permissions []*Permission
	o.Raw("select p.* from permission p " +
		"left join role_permissions rp on p.id = rp.permission_id " +
		"left join role r on rp.role_id = r.id " +
		"left join user_roles ur on r.id = ur.role_id " +
		"left join user u on ur.user_id = u.id " +
		"where u.id = ?", id).QueryRows(&permissions)
	return permissions
}

func FindPermissionByUserIdAndPermissionName(userId int, name string) bool {
	o := orm.NewOrm()
	var permission Permission
	o.Raw("select p.* from permission p " +
		"left join role_permissions rp on p.id = rp.permission_id " +
		"left join role r on rp.role_id = r.id " +
		"left join user_roles ur on r.id = ur.role_id " +
		"left join user u on ur.user_id = u.id " +
		"where u.id = ? and p.name = ?", userId, name).QueryRow(&permission)
	return permission.Id > 0
}

func DeleteUser(user *User) {
    o := orm.NewOrm()
    o.Delete(user)
}

func DeleteUserRolesByUserId(user_id int) {
    o := orm.NewOrm()
    o.Raw("delete from user_roles where user_id = ?", user_id).Exec()
}

func SaveUserRole(user_id int, role_id int) {
    o := orm.NewOrm()
    o.Raw("insert into user_roles (user_id, role_id) values (?, ?)", user_id, role_id).Exec()
}

func FindUserRolesByUserId(user_id int) []orm.Params {
    o := orm.NewOrm()
    var res []orm.Params
    o.Raw("select id, user_id, role_id from user_roles where user_id = ?", user_id).Values(&res, "id", "user_id", "role_id")
    return res
}