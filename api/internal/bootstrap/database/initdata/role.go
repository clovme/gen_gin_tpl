package initdata

import (
	modelAuth "gen_gin_tpl/internal/models/auth"
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/logger/log"
)

func (d *InitData) Role() {
	var modelList []modelAuth.Role

	for _, enum := range role.Values() {
		modelList = append(modelList, modelAuth.Role{Name: enum.Name(), Type: enum, Code: enum.Key(), CreatedBy: int64(role.System), Description: enum.Desc()})
	}

	newModelList := insertIfNotExist[modelAuth.Role](modelList, func(model modelAuth.Role) (*modelAuth.Role, error) {
		return d.Q.Role.Where(d.Q.Role.Type.Eq(int(model.Type)), d.Q.Role.CreatedBy.Eq(int64(role.System))).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := d.Q.Role.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msgf("[%s]初始化失败:", "角色表")
	} else {
		log.Info().Msgf("[%s]初始化成功，共%d条数据！", "角色表", len(newModelList))
	}
}
