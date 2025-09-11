package initdata

import (
	modelAuth "gen_gin_tpl/internal/models/auth"
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/logger/log"
)

func (r *InitData) Role() {
	var modelList []modelAuth.Role

	for _, enum := range role.Values() {
		if enum == role.System {
			modelList = append(modelList, modelAuth.Role{ID: enum.ID(), Name: enum.Name(), Type: enum, Code: enum.Key(), Status: status.Enable, RoleGroupID: 1000000000000000000, CreatedBy: int64(role.System), Description: enum.Desc()})
		} else if enum == role.Admin {
			modelList = append(modelList, modelAuth.Role{ID: enum.ID(), Name: enum.Name(), Type: enum, Code: enum.Key(), Status: status.Enable, RoleGroupID: 1000000000000000001, CreatedBy: int64(role.System), Description: enum.Desc()})
		} else {
			modelList = append(modelList, modelAuth.Role{ID: enum.ID(), Name: enum.Name(), Type: enum, Code: enum.Key(), Status: status.Enable, CreatedBy: int64(role.System), Description: enum.Desc()})
		}
	}

	newModelList := insertIfNotExist[modelAuth.Role](modelList, func(model modelAuth.Role) (*modelAuth.Role, error) {
		return r.Q.Role.Where(r.Q.Role.Type.Eq(int(model.Type)), r.Q.Role.CreatedBy.Eq(int64(role.System))).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.Role.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[角色表]初始化失败！")
	} else {
		log.Info().Msgf("[角色表]初始化成功，共%d条数据！", len(newModelList))
	}
}
