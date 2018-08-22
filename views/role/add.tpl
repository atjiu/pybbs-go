<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">添加角色</div>
      <div class="panel-body">
        {{template "../components/flash_error.tpl" .}}
        <form action="/role/add" method="post">
          <div class="form-group">
            <label for="name">角色名称</label>
            <input type="text" id="name" name="name" class="form-control">
          </div>
          <div class="form-group">
            <label>权限</label>
            <div>
              {{range .Permissions}}
                <h4><b>{{.Description}}</b></h4>
                {{range .ChildPermissions}}
                  <input type="checkbox" name="permissionIds" value="{{.Id}}" id="permission_{{.Id}}">
                  <label for="permission_{{.Id}}">{{.Description}}</label>&nbsp;
                {{end}}
              {{end}}
            </div>
          </div>
          <button type="submit" class="btn btn-sm btn-default">保存</button>
        </form>
      </div>
    </div>
  </div>
</div>