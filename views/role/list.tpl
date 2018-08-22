<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">
        角色管理
        <a href="/role/add" class="pull-right">添加角色</a>
      </div>
      <div class="table-responsive">
        <table class="table table-striped table-responsive">
          <tbody>
          {{range .Roles}}
          <tr>
            <td>{{.Id}}</td>
            <td>{{.Name}}</td>
            <td>
              <a href="/role/edit/{{.Id}}" class="btn btn-xs btn-warning">配置权限</a>
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/role/delete/{{.Id}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>