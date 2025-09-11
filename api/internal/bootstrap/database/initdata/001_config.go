package initdata

import (
	"gen_gin_tpl/internal/libs"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/logger/log"
)

// Config 初始化配置
func (r *InitData) Config() {
	modelList := libs.WebConfig.GetModelList()

	newModelList := insertIfNotExist[models.Config](modelList, func(model models.Config) (*models.Config, error) {
		return r.Q.Config.Where(r.Q.Config.Name.Eq(model.Name)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.Config.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[系统配置表]初始化失败！")
	} else {
		log.Info().Msgf("[系统配置表]初始化成功，共%d条数据！", len(newModelList))
	}
}
