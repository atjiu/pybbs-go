<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">
        编辑权限
      </div>
      <div class="panel-body">
        <div class="panel-body">
          {{template "../components/flash_error.tpl" .}}
          <form action="/permission/edit/{{.Permission.Id}}" method="post">
            <div class="form-group">
              <label for="pid">父节点</label>
              <select name="pid" id="pid" class="form-control">
                <option value="0">添加父节点</option>
                {{range .ParantPermissions}}
                <option value="{{.Id}}">
                  {{.Description}}
                </option>
                {{end}}
              </select>
            </div>
            <div class="form-group">
              <label for="name">权限标识</label>
              <input type="text" id="name" name="name" value="{{.Permission.Name}}" class="form-control" placeholder="权限标识">
            </div>
            <div class="form-group">
              <label for="url">授权地址</label>
              <input type="text" id="url" name="url" value="{{.Permission.Url}}" class="form-control" placeholder="授权地址">
            </div>
            <div class="form-group">
              <label for="description">权限描述</label>
              <input type="text" id="description" name="description" value="{{.Permission.Description}}" class="form-control" placeholder="权限描述">
            </div>
            <button type="submit" class="btn btn-default btn-sm">保存</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript">
  $(function () {
    $("#pid").val('{{.Permission.Pid}}');
  })
</script>