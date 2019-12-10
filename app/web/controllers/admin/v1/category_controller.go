package v1

import (
	"blog/app/models"
	"blog/app/repositories"
	"blog/app/web/responses/admin"
	"blog/database"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type CategoryController struct {
	Ctx                iris.Context
	CategoryRepository *repositories.CategoryRepository
	CategoryResponse   admin.CategoryResponse
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		CategoryRepository: repositories.NewCategoryRepository(
			database.DB())}
}

/**
 * 新增栏目列表
 */
func (c CategoryController) PostCreate() *simple.JsonResult {
	var (
		name        = c.Ctx.FormValue("name")
		description = c.Ctx.FormValue("description")
	)
	if name == "" {
		return simple.JsonErrorMsg("参数不能为空")
	}
	Category, err := c.CategoryRepository.Create(&models.Category{
		Name:        name,
		Description: description,
		Status:      0,
	})
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.CategoryResponse.Category(Category))
}

/**
 * 获取栏目列表
 */
func (c CategoryController) GetList() *simple.JsonResult {

	var page = &simple.Paging{
		Page:  simple.FormValueIntDefault(c.Ctx, "page", 1),
		Limit: simple.FormValueIntDefault(c.Ctx, "limit", 10),
		Total: 0,
	}
	list := c.CategoryRepository.CategoryList(page)
	return simple.JsonData(
		simple.PageResult{
			Page:    page,
			Results: c.CategoryResponse.Categorys(list),
		})
}

/**
 * 编辑栏目
 */
func (c CategoryController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt(c.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	data := repositories.UpdateData{
		"name":        c.Ctx.FormValue("name"),
		"description": c.Ctx.FormValue("description"),
	}
	Category, err := c.CategoryRepository.UpdateById(id, data)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.CategoryResponse.Category(Category))
}

/**
 * 删除栏目列表
 */
func (c CategoryController) PostDel() *simple.JsonResult {
	id, err := simple.FormValueInt(c.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	err = c.CategoryRepository.DelById(id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}
