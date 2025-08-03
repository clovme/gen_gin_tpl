package initdata

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/boolean"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/variable"
	"gen_gin_tpl/public"
)

// Config 初始化配置
func (d *InitData) Config() {
	modelList := []models.Config{
		{Name: constants.ContextIsEncrypted, Value: boolean.False.Key(), Default: boolean.False.Key(), Show: boolean.True, Description: "是否开启加密模式"},
		{Name: constants.WebTitle, Value: variable.WebTitle, Default: variable.WebTitle, Show: boolean.True, Description: "站点标题"},
		{Name: constants.PublicPEM, Value: string(public.PublicPEM), Default: string(public.PublicPEM), Show: boolean.True, Description: "加密公钥"},
		{Name: constants.PrivatePEM, Value: string(public.PrivatePEM), Default: string(public.PrivatePEM), Show: boolean.True, Description: "加密私钥"},
		{Name: constants.Countdown, Value: "60", Default: "60", Show: boolean.True, Description: "统一倒计时时间，单位秒"},
	}

	newModelList := insertIfNotExist[models.Config](modelList, func(model models.Config) (*models.Config, error) {
		return d.Q.Config.Where(d.Q.Config.Name.Eq(model.Name)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := d.Q.Config.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msgf("[%s]初始化失败:", "系统配置表")
	} else {
		log.Info().Msgf("[%s]初始化成功，共%d条数据！", "系统配置表", len(newModelList))
	}
}
