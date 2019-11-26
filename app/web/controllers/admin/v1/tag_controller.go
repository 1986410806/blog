package v1

import (
	"blog/app/models"
	"blog/app/repositories"
	"blog/app/web/responses"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type TagController struct {
	Ctx           iris.Context
	TagRepository *repositories.TagRepositories
	TagResponse   responses.TagResponse
}

func NewTagController() *TagController {
	return &TagController{
		TagRepository: repositories.NewTagRepositories()}
}

/**
 * 新增标签列表
 */
func (c TagController) PostCreate() *simple.JsonResult {
	var (
		name        = c.Ctx.FormValue("name")
		description = c.Ctx.FormValue("description")
	)
	if name == "" {
		return simple.JsonErrorMsg("参数不能为空")
	}
	tag, err := c.TagRepository.Create(&models.Tag{
		Name:        name,
		Description: description,
		Status:      0,
	})
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.TagResponse.Tag(tag))
}

/**
 * 获取标签列表
 */
func (c TagController) GetList() *simple.JsonResult {
	list, err := c.TagRepository.TagList()
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.TagResponse.Tags(list))
}

/**
 * 编辑标签
 */
func (c TagController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt(c.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	data := repositories.UpdateData{
		"name":        c.Ctx.FormValue("name"),
		"description": c.Ctx.FormValue("description"),
	}
	tag, err := c.TagRepository.UpdateById(id, data)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.TagResponse.Tag(tag))
}

/**
 * 删除标签列表
 */
func (c TagController) PostDel() *simple.JsonResult {
	id, err := simple.FormValueInt(c.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	err = c.TagRepository.DelById(id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}
