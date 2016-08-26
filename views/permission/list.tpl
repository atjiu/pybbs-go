<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        权限管理
        <a class="pull-right" href="/permission/add?pid={{.Pid}}">添加权限</a>
      </div>
      <div class="table-responsive">
        <table class="table table-striped table-responsive">
          <tbody>
          {{range .Permissions}}
          <tr>
            <td>{{.Id}}</td>
            <td>{{.Name}}</td>
            <td>{{.Url}}</td>
            <td>{{.Description}}</td>
            <td>
              <a href="/permission/edit/{{.Id}}" class="btn btn-xs btn-warning">编辑</a>
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/permission/delete/{{.Id}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
    <div class="panel panel-default">
      <div class="panel-heading">父节点</div>
      <div class="list-group">
        {{range .ParantPermissions}}
        <li class="list-group-item permission-item" id="list-group-item-{{.Id}}">
          <a href="javascript:if(confirm('确认删除吗?'))location.href='/permission/delete/{{.Id}}'">删除</a>
          <a href="/permission/list?pid={{.Id}}">
            {{.Description}}
          </a>
        </li>
        {{end}}
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
  $(function () {
    $("#list-group-item-{{.Pid}}").addClass("active");
  });
</script>