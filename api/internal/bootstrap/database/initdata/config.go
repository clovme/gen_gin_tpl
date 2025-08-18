package initdata

import (
	"gen_gin_tpl/internal/libs"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/logger/log"
)

// Config 初始化配置
func (d *InitData) Config() {
	modelList := libs.WebConfig.GetModelList()

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
