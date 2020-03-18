package v1

import (
	"github.com/suisai337/goblog/models"
	"github.com/suisai337/goblog/pkg/app"
	"github.com/suisai337/goblog/pkg/e"
	"github.com/suisai337/goblog/pkg/logging"
	"github.com/suisai337/goblog/pkg/setting"
	"github.com/suisai337/goblog/pkg/util"
	"github.com/suisai337/goblog/service/article_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 0, "id").Message("id必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{ID:id}
	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("state错误")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Range(tagId, 0, 1, "tag_id").Message("tag_id错误")
	}

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	} else {
		code = e.SUCCESS
		data["list"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})

}

//新增文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	coverImageUrl := c.Query("cover_image_url")

	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(coverImageUrl, "cover_image_url").Message("封面图不能为空")
	valid.MaxSize(coverImageUrl, 200, "cover_image_url").Message("封面图长度不能大于200")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	} else {
		if ! models.ExistTagByID(tagId) {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["cover_image_url"] = coverImageUrl
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})


}

//修改文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	coverImageUrl := c.Query("cover_image_url")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(state).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(coverImageUrl, "cover_image_url").Message("封面图不能为空")
	valid.MaxSize(coverImageUrl, 200, "cover_image_url").Message("封面图长度不能大于200")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	} else {
		if ! models.ExistArticleByID(id) {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
		if ! models.ExistTagByID(tagId) {
			code = e.ERROR_NOT_EXIST_TAG
		}
		data := make(map[string]interface{})
		if tagId > 0 {
			data["tag_id"] = tagId
		}
		if title != "" {
			data["title"] = title
		}
		if desc != "" {
			data["desc"] = desc
		}
		if content != "" {
			data["content"] =content
		}
		if coverImageUrl != "" {
			data["cover_image_url"] = coverImageUrl
		}
		data["modified_by"] = modifiedBy

		models.EditArticle(id, data)
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1 , "id").Message("")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	} else {
		if ! models.ExistArticleByID(id) {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			models.DeleteArticle(id)
			code = e.SUCCESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}