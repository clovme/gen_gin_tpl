package initdata

import (
	"gen_gin_tpl/internal/models/auth"
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/logger/log"
)

func (r *InitData) UserRole() {
	var modelList []auth.UserRole
	for _, re := range role.Values() {
		if re != role.System && re != role.Admin {
			continue
		}
		modelList = append(modelList, auth.UserRole{
			UserID:      re.ID(),
			RoleID:      re.ID(),
			Status:      status.Enable,
			Description: re.Desc(),
		})
	}

	newModelList := insertIfNotExist[auth.UserRole](modelList, func(model auth.UserRole) (*auth.UserRole, error) {
		return r.Q.UserRole.Where(r.Q.UserRole.UserID.Eq(model.UserID), r.Q.UserRole.RoleID.Eq(model.RoleID)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.UserRole.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[用户角色表]初始化失败！")
	} else {
		log.Info().Msgf("[用户角色表]初始化成功，共%d条数据！", len(newModelList))
	}
}
