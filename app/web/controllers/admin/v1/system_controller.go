package v1

import (
	"blog/app/web/services"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
)

type SystemController struct {
	Ctx           iris.Context
	SystemService services.SystemService
}

func NewSystemController() *SystemController {
	return &SystemController{}
}

/**
 *	获取系统配置
 * @return json map[string]interface{}
 */
func (this *SystemController) GetConfig() *simple.JsonResult {
	config := this.SystemService.GetSystemConfigs()
	return simple.JsonData(config)
}