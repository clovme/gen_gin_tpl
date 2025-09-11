package initdata

import (
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/models/auth"
	"gen_gin_tpl/pkg/logger/log"
)

func (r *InitData) RolePermission() {
	var modelList []auth.RolePermission

	roleList, err := query.Q.Role.Select(query.Q.Role.ID).Where(query.Q.Role.RoleGroupID.Neq(0)).Find()
	if err != nil {
		log.Error().Err(err).Msg("[角色权限表]初始化失败,查询角色表失败！")
		return
	}

	permissionList, err := query.Q.Permission.Select(query.Q.Permission.ID).Find()
	if err != nil {
		log.Error().Err(err).Msg("[角色权限表]初始化失败,查询权限表失败！")
		return
	}

	for _, role := range roleList {
		for _, permission := range permissionList {
			modelList = append(modelList, auth.RolePermission{
				RoleID:       role.ID,
				PermissionID: permission.ID,
			})
		}
	}

	newModelList := insertIfNotExist[auth.RolePermission](modelList, func(model auth.RolePermission) (*auth.RolePermission, error) {
		return r.Q.RolePermission.Where(r.Q.RolePermission.RoleID.Eq(model.RoleID), r.Q.RolePermission.PermissionID.Eq(model.PermissionID)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.RolePermission.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[角色权限表]初始化失败!")
	} else {
		log.Info().Msgf("[角色权限表]初始化成功，共%d条数据！", len(newModelList))
	}
}
