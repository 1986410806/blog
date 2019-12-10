package v1

import (
	"blog/app/models"
	"blog/app/repositories"
	"blog/app/web/responses/admin"
	"blog/app/web/services"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type SystemConfigController struct {
	Ctx                    iris.Context
	SystemConfigRepository *repositories.SystemConfigRepository
	SystemService          services.SystemService
	SystemConfigResponse   admin.SystemConfigResponse
}

func NewSystemConfigController() *SystemConfigController {
	return &SystemConfigController{
		SystemConfigRepository: repositories.NewSystemConfigRepository(),
	}
}

/**
 *	获取系统配置
 * @return json map[string]interface{}
 */
func (this *SystemConfigController) GetConfig() *simple.JsonResult {
	config := this.SystemService.GetSystemConfigs()
	return simple.JsonData(config)
}

/**
 *	获取系统配置列表
 * @return json []models.SysConfig{}
 */
func (c *SystemConfigController) GetList() *simple.JsonResult {

	var page = &simple.Paging{
		Page:  simple.FormValueIntDefault(c.Ctx, "page", 1),
		Limit: simple.FormValueIntDefault(c.Ctx, "limit", 10),
		Total: 0,
	}
	list := c.SystemConfigRepository.ConfigList(page)

	return simple.JsonData(
		simple.PageResult{
			Page:    page,
			Results: c.SystemConfigResponse.List(list),
		})
}

/**
 *	新增系统配置
 * @return json models.SysConfig{}
 */
func (c *SystemConfigController) PostCreate() *simple.JsonResult {
	var (
		key         = c.Ctx.FormValue("key")
		value       = c.Ctx.FormValue("value")
		name        = c.Ctx.FormValue("name")
		description = c.Ctx.FormValue("description")
	)
	if name == "" {
		return simple.JsonErrorMsg("参数不能为空")
	}
	page, err := c.SystemConfigRepository.Create(&models.SysConfig{
		Key:         key,
		Value:       value,
		Name:        name,
		Description: description,
	})
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.SystemConfigResponse.Item(page))
}

/**
 * 编辑系统配置
 * @return json models.SysConfig{}
 */
func (c *SystemConfigController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt(c.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	data := repositories.UpdateData{
		"name":        c.Ctx.FormValue("name"),
		"description": c.Ctx.FormValue("description"),
	}
	page, err := c.SystemConfigRepository.UpdateById(id, data)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.SystemConfigResponse.Item(page))
}

/**
 * 删除系统配置
 * @return json models.SysConfig{}
 */
func (c *SystemConfigController) PostDel() *simple.JsonResult {
	id, err := simple.FormValueInt(c.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	err = c.SystemConfigRepository.DelById(id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}
