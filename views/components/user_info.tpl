<div class="panel panel-default">
  <div class="panel-heading">
    个人信息
  </div>
  <div class="panel-body">
    <div class="media">
      <div class="media-left">
        <a href="/user/{{.UserInfo.Username}}">
          <img src="{{.UserInfo.Avatar}}" title="{{.UserInfo.Username}}" class="avatar">
        </a>
      </div>
      <div class="media-body">
        <div class="media-heading">
          <a href="/user/{{.UserInfo.Username}}">{{.UserInfo.Username}}</a>
        </div>
        {{if .UserInfo.Url}}<a href="{{.UserInfo.Url}}" target="_blank">{{.UserInfo.Url}}</a>{{end}}
      </div>
      {{if .UserInfo.Signature}}
      <div style="color: #7A7A7A; font-size: 12px; margin-top:5px;">
        <i>“ {{.UserInfo.Signature}} ” </i>
      </div>
      {{end}}
    </div>
  </div>
</div>