package models

import "github.com/astaxie/beego/orm"

type Permission struct {
    Id          int `orm:"pk;auto"`
    Pid         int
    Url         string
    Name        string
    Description string
    Roles       []*Role `orm:"reverse(many)"`
    ChildPermissions []*Permission `orm:"-"`
}

func FindPermissionById(id int) Permission {
    o := orm.NewOrm()
    var permission Permission
    o.QueryTable(permission).Filter("Id", id).One(&permission)
    return permission
}

func FindPermissions() []*Permission {
    o := orm.NewOrm()
    var permission Permission
    var permissions []*Permission
    o.QueryTable(permission).All(&permissions)
    return permissions
}

func FindPermissionsByPid(pid int) []*Permission {
    o := orm.NewOrm()
    var permission Permission
    var permissions []*Permission
    o.QueryTable(permission).Filter("Pid", pid).All(&permissions)
    return permissions
}

func SavePermission(permission *Permission) int64 {
    o := orm.NewOrm()
    id, _ := o.Insert(permission)
    return id
}

func UpdatePermission(permission *Permission) int64 {
    o := orm.NewOrm()
    id, _ := o.Update(permission)
    return id
}

func DeletePermission(permission *Permission) {
    o := orm.NewOrm()
    o.Delete(permission)
}

func DeleteRolePermissionByPermissionId(permission_id int) {
    o := orm.NewOrm()
    o.Raw("delete from role_permissions where permission_id = ?", permission_id).Exec()
}