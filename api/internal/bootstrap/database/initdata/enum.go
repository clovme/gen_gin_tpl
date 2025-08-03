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

func (d *InitData) Enums() {
	var modelList []models.Enums

	for i, enum := range role.Values() {
		modelList = append(modelList, models.Enums{Category: role.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Enum(), ValueT: dtype.Int.Enum(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range gender.Values() {
		modelList = append(modelList, models.Enums{Category: gender.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Enum(), ValueT: dtype.Int.Enum(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range code.Values() {
		modelList = append(modelList, models.Enums{Category: code.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Enum(), ValueT: dtype.Int.Enum(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range status.Values() {
		modelList = append(modelList, models.Enums{Category: status.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Enum(), ValueT: dtype.Int.Enum(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range perm.Values() {
		modelList = append(modelList, models.Enums{Category: perm.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Enum(), ValueT: dtype.Int.Enum(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range dtype.Values() {
		modelList = append(modelList, models.Enums{Category: dtype.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Enum(), ValueT: dtype.Int.Enum(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range boolean.Values() {
		modelList = append(modelList, models.Enums{Category: boolean.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Enum(), ValueT: dtype.Int.Enum(), Sort: i + 1, Description: enum.Desc()})
	}

	newModelList := insertIfNotExist[models.Enums](modelList, func(model models.Enums) (*models.Enums, error) {
		return d.Q.Enums.Where(d.Q.Enums.Category.Eq(model.Category), d.Q.Enums.Value.Eq(model.Value), d.Q.Enums.Key.Eq(model.Key)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := d.Q.Enums.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msgf("[%s]初始化失败:", "系统枚举")
	} else {
		log.Info().Msgf("[%s]初始化成功，共%d条数据！", "系统枚举", len(newModelList))
	}
}
