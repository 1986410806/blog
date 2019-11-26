package v1

import (
	"blog/app/models"
	"blog/app/repositories"
	"blog/app/web/responses"
	"blog/app/web/services"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type SystemConfigController struct {
	Ctx                      iris.Context
	SystemConfigRepositories *repositories.SystemConfigRepositories
	SystemService            services.SystemService
	SystemConfigResponse     responses.SystemConfigResponse
}

func NewSystemConfigController() *SystemConfigController {
	return &SystemConfigController{
		SystemConfigRepositories: repositories.NewSystemConfigRepositories(),
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
	list, err := c.SystemConfigRepositories.ConfigList()
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(
		c.SystemConfigResponse.List(list))
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
	page, err := c.SystemConfigRepositories.Create(&models.SysConfig{
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
	page, err := c.SystemConfigRepositories.UpdateById(id, data)
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
	err = c.SystemConfigRepositories.DelById(id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}
