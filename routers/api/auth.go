package api

import (
	"github.com/suisai337/goblog/models"
	"github.com/suisai337/goblog/pkg/e"
	"github.com/suisai337/goblog/pkg/logging"
	"github.com/suisai337/goblog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth{Username: username, Password: password})

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		if ! models.CheckAuth(username, password) {
			code = e.ERROR_AUTH
		} else {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				code = e.SUCCESS
				data["token"] = token
			}
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
