<div class="row">
  <div class="col-md-6">
    <div class="panel panel-default">
      <div class="panel-heading">登录</div>
      <div class="panel-body">
        {{template "components/flash_error.tpl" .}}
        <form action="/login" method="post">
          <div class="form-group">
            <label for="username">用户名</label>
            <input type="text" id="username" name="username" class="form-control" placeholder="用户名">
          </div>
          <div class="form-group">
            <label for="password">密码</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="密码">
          </div>
          <input type="submit" class="btn btn-default" value="登录"> <a href="/register">去注册</a>
        </form>
      </div>
    </div>
  </div>
</div>