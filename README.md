> 使用说明：请保留页面底部的 powered by 朋也社区

## 特性

- 不用session,选用cookie,为了集群方便
- 权限配置简单,轻松管理用户

## 依赖

- [github.com/astaxie/beego](https://github.com/astaxie/beego)
- [github.com/astaxie/beego/context](https://github.com/astaxie/beego/context)
- [github.com/astaxie/beego/orm](https://github.com/astaxie/beego/orm)
- [github.com/xeonx/timeago](https://github.com/xeonx/timeago)
- [github.com/russross/blackfriday](https://github.com/russross/blackfriday)

## 其他版本

- Java版朋也社区: [https://bbs.tomoya.cn](https://bbs.tomoya.cn)

## 如何开始

- 克隆代码到 $GOPATH/src 下
- 安装上面的依赖
- 安装 [bee](https://github.com/beego/bee) 工具
- 进入 pybbs-go 目录 
- 运行 bee run 
- 将pybbs-go.sql导入数据库
- 浏览器输入 http://localhost:8080

## 注意

- 如果访问地址不是localhost,需要修改conf/app.conf文件里的cookie.domain,否则登录后不会记录登录状态

## 碰到问题怎么办?

1. 到 https://bbs.tomoya.cn 上提问答
2. 在Github上提 Issues
3. 加QQ群：419343003

提问题的时候请将问题重现步骤描述清楚

## 贡献

欢迎大家提pr

## 开源协议

MIT