package initdata

import (
	"gen_gin_tpl/internal/domain/do_enums"
	"gen_gin_tpl/pkg/enums/em_bool"
	"gen_gin_tpl/pkg/enums/em_gender"
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/enums/em_perm"
	"gen_gin_tpl/pkg/enums/em_role"
	"gen_gin_tpl/pkg/enums/em_status"
	"gen_gin_tpl/pkg/enums/em_type"
	"gen_gin_tpl/pkg/logger/log"
)

func (d *InitData) Enums() {
	var modelList []do_enums.Enums

	for i, enum := range em_role.Values() {
		modelList = append(modelList, do_enums.Enums{Category: em_role.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: em_type.Int.Int(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range em_gender.Values() {
		modelList = append(modelList, do_enums.Enums{Category: em_gender.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: em_type.Int.Int(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range em_http.Values() {
		modelList = append(modelList, do_enums.Enums{Category: em_http.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: em_type.Int.Int(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range em_status.Values() {
		modelList = append(modelList, do_enums.Enums{Category: em_status.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: em_type.Int.Int(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range em_perm.Values() {
		modelList = append(modelList, do_enums.Enums{Category: em_perm.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: em_type.Int.Int(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range em_type.Values() {
		modelList = append(modelList, do_enums.Enums{Category: em_type.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: em_type.Int.Int(), Sort: i + 1, Description: enum.Desc()})
	}

	for i, enum := range em_bool.Values() {
		modelList = append(modelList, do_enums.Enums{Category: em_bool.Name, Key: enum.Key(), Name: enum.Name(), Value: enum.Int(), ValueT: em_type.Int.Int(), Sort: i + 1, Description: enum.Desc()})
	}

	newModelList := insertIfNotExist[do_enums.Enums](modelList, func(model do_enums.Enums) (*do_enums.Enums, error) {
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
