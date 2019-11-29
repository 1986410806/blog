package v1

import (
	"blog/app/common/jwt"
	"blog/app/repositories"
	"blog/app/web/responses/admin"
	"blog/database"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type ArticleController struct {
	Ctx               iris.Context
	ArticleRepository *repositories.ArticleRepository
	ArticleResponse   admin.ArticleResponse
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		ArticleRepository: repositories.NewArticleRepository(
			database.DB())}
}

/**
 * 新增文章列表
 */
func (this ArticleController) PostCreate() *simple.JsonResult {
	tokenClaims := jwt.GetTokenClaim()
	var (
		tags            = simple.FormValueStringArray(this.Ctx, "tags")
		title           = this.Ctx.PostValueTrim("title")
		summary         = this.Ctx.PostValue("summary")
		content         = this.Ctx.PostValue("content")
		categoryId, err = simple.FormValueInt(this.Ctx, "category_id")
	)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	// 手动 validator
	switch {
	case err != nil:
		println("请选择栏目～")
		return simple.JsonErrorMsg("请选择栏目～")
	case tags == nil:
		println("请选择标签～")
		return simple.JsonErrorMsg("请选择标签～")
	case title == "":
		println("请填写标题～")
		return simple.JsonErrorMsg("请填写标题～")
	case summary == "":
		println("请填写摘要～")
		return simple.JsonErrorMsg("请填写摘要～")
	case content == "":
		println("内容不能为空～")
		return simple.JsonErrorMsg("内容不能为空～")
	}

	article, err := this.ArticleRepository.Create(tokenClaims.UserId, title, summary, content, categoryId, tags)

	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		this.ArticleResponse.Create(article))
}

/**
 * 获取文章列表
 */
func (c ArticleController) GetList() *simple.JsonResult {
	return simple.JsonData("")
}

/**
 * 编辑文章
 */
func (this ArticleController) PostUpdate() *simple.JsonResult {
	var (
		tags                      = simple.FormValueStringArray(this.Ctx, "tags")
		title                     = this.Ctx.PostValueTrim("title")
		summary                   = this.Ctx.PostValue("summary")
		content                   = this.Ctx.PostValue("content")
		articleId, articleIdErr   = simple.FormValueInt(this.Ctx, "article_id")
		categoryId, categoryIdErr = simple.FormValueInt(this.Ctx, "category_id")
	)

	// 手动 validator
	switch {
	case articleIdErr != nil:
		println("文章id不能为空～")
		return simple.JsonErrorMsg("文章id不能为空～")
	case categoryIdErr != nil:
		println("请选择栏目～")
		return simple.JsonErrorMsg("请选择栏目～")
	case tags == nil:
		println("请选择标签～")
		return simple.JsonErrorMsg("请选择标签～")
	case title == "":
		println("请填写标题～")
		return simple.JsonErrorMsg("请填写标题～")
	case summary == "":
		println("请填写摘要～")
		return simple.JsonErrorMsg("请填写摘要～")
	case content == "":
		println("内容不能为空～")
		return simple.JsonErrorMsg("内容不能为空～")
	}

	article, err := this.ArticleRepository.UpdateById(uint(articleId), title, summary, content, categoryId, tags)

	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		this.ArticleResponse.Create(article))
}

/**
 * 删除文章
 */
func (this ArticleController) PostDel() *simple.JsonResult {
	return simple.JsonSuccess()
}
