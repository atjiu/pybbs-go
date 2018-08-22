<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">
        用户管理
        <span class="pull-right">{{.Page.TotalCount}}个用户</span>
      </div>
      <div class="table-responsive">
        <table class="table table-striped">
          <tbody>
          {{range .Page.List}}
          <tr>
            <td>{{.Id}}</td>
            <td><a href="/user/{{.Username}}" target="_blank">{{.Username}}</a></td>
            <td>{{if .Url}}<a href="{{.Url}}" target="_blank">{{.Url}}</a>{{end}}</td>
            <td>
              <a href="/user/edit/{{.Id}}" class="btn btn-xs btn-warning">配置角色</a>
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/user/delete/{{.Id}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
      <div class="panel-body" style="padding: 0 15px;">
        <ul id="page"></ul>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
<script type="text/javascript">
  $(function () {
    $("#page").bootstrapPaginator({
      currentPage: '{{.Page.PageNo}}',
      totalPages: '{{.Page.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        window.location.href = "/?p=" + page
      }
    });
  });
</script>