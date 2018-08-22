<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">配置角色</div>
      <div class="panel-body">
        <form action="/user/edit/{{.User.Id}}" method="post">
          <div class="form-group">
            <label for="username">用户名</label>
            <input type="text" disabled id="username" value="{{.User.Username}}" class="form-control">
          </div>
          <div class="form-group">
            <label>角色</label>
            <div>
              {{range .Roles}}
                <input type="checkbox" name="roleIds" value="{{.Id}}" id="role_{{.Id}}">
                <label for="role_{{.Id}}">{{.Name}}</label>&nbsp;
              {{end}}
            </div>
          </div>
          <button type="submit" class="btn btn-sm btn-default">保存</button>
        </form>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
  $(function () {
    {{range .UserRoles}}
      $("#role_{{.role_id}}").attr("checked", true);
    {{end}}
  })
</script>