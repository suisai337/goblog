module github.com/suisai337/goblog

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.1
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.52.0
	github.com/go-openapi/spec v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/robfig/cron v1.2.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	golang.org/x/tools v0.0.0-20200305224536-de023d59a5d1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/suisai337/goblog/conf => /User/yanfan/go-application/go-gin-example/pkg/conf
	github.com/suisai337/goblog/docs => /User/yanfan/go-application/go-gin-example/pkg/docs
	github.com/suisai337/goblog/middleware => /User/yanfan/go-application/go-gin-example/middleware
	github.com/suisai337/goblog/middleware/jwt => /User/yanfan/go-application/go-gin-example/middleware/jwt
	github.com/suisai337/goblog/models => /User/yanfan/go-application/go-gin-example/models
	github.com/suisai337/goblog/pkg/app => /User/yanfan/go-application/go-gin-example/pkg/app
	github.com/suisai337/goblog/pkg/e => /User/yanfan/go-application/go-gin-example/pkg/e
	github.com/suisai337/goblog/pkg/file => /User/yanfan/go-application/go-gin-example/pkg/file
	github.com/suisai337/goblog/pkg/gredis => /User/yanfan/go-application/go-gin-example/pkg/gredis
	github.com/suisai337/goblog/pkg/logging => /User/yanfan/go-application/go-gin-example/pkg/logging
	github.com/suisai337/goblog/pkg/setting => /User/yanfan/go-application/go-gin-example/pkg/setting
	github.com/suisai337/goblog/pkg/upload => /User/yanfan/go-application/go-gin-example/pkg/upload
	github.com/suisai337/goblog/pkg/util => /User/yanfan/go-application/go-gin-example/pkg/util
	github.com/suisai337/goblog/routers => /User/yanfan/go-application/go-gin-example/routers
	github.com/suisai337/goblog/routers/api => /User/yanfan/go-application/go-gin-example/routers/api
	github.com/suisai337/goblog/service => /User/yanfan/go-application/go-gin-example/service
	github.com/suisai337/goblog/service/article_service => /User/yanfan/go-application/go-gin-example/service/article_service
	github.com/suisai337/goblog/service/cache_service => /User/yanfan/go-application/go-gin-example/service/cache_service
)
