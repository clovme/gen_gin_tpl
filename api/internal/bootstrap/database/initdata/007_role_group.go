package initdata

import (
	"gen_gin_tpl/internal/models/auth"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/logger/log"
)

func (r *InitData) RoleGroup() {
	modelList := []auth.RoleGroup{
		{ID: 1000000000000000000, Name: "系统管理员组", Description: "拥有系统最高权限，管理平台所有配置与用户操作。", Status: status.Enable},
		{ID: 1000000000000000001, Name: "超级管理员组", Description: "具备高级管理权限，负责日常运营与用户管理。", Status: status.Enable},
	}

	newModelList := insertIfNotExist[auth.RoleGroup](modelList, func(model auth.RoleGroup) (*auth.RoleGroup, error) {
		return r.Q.RoleGroup.Where(r.Q.RoleGroup.ID.Eq(model.ID)).Where(r.Q.RoleGroup.Name.Eq(model.Name)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.RoleGroup.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[角色组表]初始化失败！")
	} else {
		log.Info().Msgf("[角色组表]初始化成功，共%d条数据！", len(newModelList))
	}
}
