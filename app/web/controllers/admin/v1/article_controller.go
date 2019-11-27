package v1

import (
	"blog/app/common/jwt"
	"blog/app/repositories"
	"blog/app/web/responses/admin"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type ArticleController struct {
	Ctx               iris.Context
	ArticleRepository *repositories.ArticleRepository
	ArticleResponse   admin.TagResponse
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		ArticleRepository: repositories.NewArticleRepository()}
}

/**
 * 新增文章列表
 */
func (this ArticleController) ArticleController() *simple.JsonResult {
	tokenClaims := jwt.GetTokenClaim(this.Ctx)
	var (
		tags             = simple.FormValueStringArray(this.Ctx, "tags")
		title            = this.Ctx.PostValueTrim("title")
		summary          = this.Ctx.PostValue("summary")
		content          = this.Ctx.PostValue("content")
		category_id, err = simple.FormValueInt(this.Ctx, "category_id")
	)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	// 手动 validator
	switch {
	case err != nil:
		return simple.JsonErrorMsg("请选择栏目～")
	case tags == nil:
		return simple.JsonErrorMsg("请选择标签～")
	case title == "":
		return simple.JsonErrorMsg("请填写标题～")
	case summary == "":
		return simple.JsonErrorMsg("请填写摘要～")
	case content == "":
		return simple.JsonErrorMsg("内容不能为空～")
	}

	article, err := this.ArticleRepository.Create(tokenClaims.UserId, title, summary, content, category_id, tags)

	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(article)
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
func (c ArticleController) PostUpdate() *simple.JsonResult {
	return simple.JsonData("")
}

/**
 * 删除文章
 */
func (c ArticleController) PostDel() *simple.JsonResult {
	return simple.JsonSuccess()
}
