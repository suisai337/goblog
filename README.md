# goblog
Go Gin Blog 同时引用beego种的gorm 验证规则等

## 安装
`$ go get github.com/suisai337/goblog`

## 运行条件
### 必须：
- mysql
- redis

### 配置：
你应该修改 conf/app.ini 配置文件

[database]|
--|
Type = mysql|
User = root|
Password = rootroot
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_`

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...`

## 运行：
` $ cd $GOPATH/src/gobolg `

$ go run main.go 
项目的运行信息和已存在的 API's

`[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/suisai337/goblog/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/suisai337/goblog/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/suisai337/goblog/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/suisai337/goblog/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/suisai337/goblog/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/suisai337/goblog/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/suisai337/goblog/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/suisai337/goblog/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/suisai337/goblog/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/suisai337/goblog/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/suisai337/goblog/routers/api/v1.DeleteArticle (4 handlers)

Listening port is 8000
Actual pid is 4393

## 特性
- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Cron
- Redis
