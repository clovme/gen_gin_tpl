package initdata

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/enums/boolean"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/enums/dtype"
	"gen_gin_tpl/pkg/enums/gender"
	"gen_gin_tpl/pkg/enums/perm"
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/logger/log"
)

func (r *InitData) Enums() {
	var modelList []models.Enums

	for i, enum := range role.Values() {
		modelList = append(modelList, models.Enums{EID: enum.ID(), Category: role.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: dtype.Int, Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range gender.Values() {
		modelList = append(modelList, models.Enums{EID: enum.ID(), Category: gender.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: dtype.Int, Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range code.Values() {
		modelList = append(modelList, models.Enums{EID: enum.ID(), Category: code.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: dtype.Int, Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range status.Values() {
		modelList = append(modelList, models.Enums{EID: enum.ID(), Category: status.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: dtype.Int, Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range perm.Values() {
		modelList = append(modelList, models.Enums{EID: enum.ID(), Category: perm.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: dtype.Int, Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range dtype.Values() {
		modelList = append(modelList, models.Enums{EID: enum.ID(), Category: dtype.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: dtype.Int, Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range boolean.Values() {
		modelList = append(modelList, models.Enums{EID: enum.ID(), Category: boolean.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: dtype.Int, Sort: i + 1, Description: enum.Desc()})
	}

	newModelList := insertIfNotExist[models.Enums](modelList, func(model models.Enums) (*models.Enums, error) {
		return r.Q.Enums.Where(r.Q.Enums.Category.Eq(model.Category), r.Q.Enums.Value.Eq(model.Value), r.Q.Enums.Key.Eq(model.Key)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.Enums.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[系统枚举]初始化失败！")
	} else {
		log.Info().Msgf("[系统枚举]初始化成功，共%d条数据！", len(newModelList))
	}
}
