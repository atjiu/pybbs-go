<div class="row">
  <div class="col-md-12">
    <div class="panel panel-default">
      <div class="panel-heading">
        <a href="/">主页</a> / 个人设置
      </div>
      <div class="panel-body">
        {{if .flash.success}}
        <div class="alert alert-success alert-dismissible" role="alert">
          <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span
            aria-hidden="true">&times;</span></button>
          {{.flash.success}}
        </div>
        {{end}}
        {{template "../components/flash_error.tpl" .}}
        <form action="/user/setting" method="post">
          <div class="form-group">
            <label for="username">昵称</label>
            <input type="text" disabled="" class="form-control" id="username" value="{{.UserInfo.Username}}">
          </div>
          <div class="form-group">
            <label for="email">邮箱</label>
            <input type="text" class="form-control" name="email" id="email" value="{{.UserInfo.Email}}">
          </div>
          <div class="form-group">
            <label for="url">个人主页</label>
            <input type="text" class="form-control" name="url" id="url" name="url" value="{{.UserInfo.Url}}">
          </div>
          <div class="form-group">
            <label for="signature">个性签名</label>
            <textarea class="form-control" name="signature" id="signature">{{.UserInfo.Signature}}</textarea>
          </div>
          <button type="submit" class="btn btn-default">保存设置</button>
        </form>
      </div>
    </div>
    <div class="panel panel-default">
      <div class="panel-heading">
        修改头像
      </div>
      <div class="panel-body">
        <form action="/user/updateavatar" method="post" enctype="multipart/form-data">
          <div class="form-group">
            <label for="oldpwd">选择头像</label>
            <input type="file" class="form-control" name="avatar">
          </div>
          <button type="submit" class="btn btn-default">上传</button>
        </form>
      </div>
    </div>
    <div class="panel panel-default">
      <div class="panel-heading">
        修改密码
      </div>
      <div class="panel-body">
        <form action="/user/updatepwd" method="post">
          <div class="form-group">
            <label for="oldpwd">旧密码</label>
            <input type="password" class="form-control" name="oldpwd" id="oldpwd" value="">
          </div>
          <div class="form-group">
            <label for="newpwd">新密码</label>
            <input type="password" class="form-control" name="newpwd" id="newpwd" value="">
          </div>
          <button type="submit" class="btn btn-default">修改密码</button>
        </form>
      </div>
    </div>
  </div>
</div>