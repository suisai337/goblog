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

[database]  
Type = mysql  
User = root  
Password = rootroot  
Host = 127.0.0.1:3306  
Name = blog  
TablePrefix = blog_  

[redis]  
Host = 127.0.0.1:6379  
Password =  
MaxIdle = 30  
MaxActive = 30  
IdleTimeout = 200  
...  

## 运行：
` $ cd $GOPATH/src/gobolg `

$ go run main.go 
项目的运行信息和已存在的 API's

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.  
 - using env:	export GIN_MODE=release  
 - using code:	gin.SetMode(gin.ReleaseMode)  

[GIN-debug]GET&emsp;&emsp;/auth&emsp;&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api.GetAuth  
[GIN-debug]GET&emsp;&emsp;/api/v1/tags&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.GetTags  
[GIN-debug]POST&emsp;&emsp;/api/v1/tags &emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.AddTag  
[GIN-debug]PUT&emsp;&emsp;/api/v1/tags/:id &emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.EditTag  
[GIN-debug]DELETE&emsp;&emsp;/api/v1/tags/:id&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.DeleteTag  
[GIN-debug]GET&emsp;&emsp;/api/v1/articles&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.GetArticles  
[GIN-debug]GET&emsp;&emsp;/api/v1/articles/:id&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.GetArticle    
[GIN-debug]POST&emsp;&emsp;/api/v1/articles&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.AddArticle  
[GIN-debug]PUT&emsp;&emsp;/api/v1/articles/:id&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.EditArticle    
[GIN-debug]DELETE&emsp;&emsp;/api/v1/articles/:id&emsp;&emsp;&emsp;--> github.com/suisai337/goblog/routers/api/v1.DeleteArticle  

Listening port is 8000  

## 特性
- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Cron
- Redis
