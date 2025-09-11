package initdata

import (
	"fmt"
	modelUser "gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/gender"
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/logger/log"
	"strings"
	"time"
)

func (r *InitData) User() {
	var modelList []modelUser.User

	for i, re := range role.Values() {
		if re != role.System && re != role.Admin {
			continue
		}

		now := time.Now()
		name := strings.ToLower(re.Key())

		data := modelUser.User{
			ID:          re.ID(),
			Username:    name,
			Email:       fmt.Sprintf("%s@%s.com", name, constants.ProjectName),
			Phone:       fmt.Sprintf("1380000000%d", i),
			Password:    name,
			Nickname:    re.Name(),
			Avatar:      "/assets/images/avatar.png",
			Gender:      gender.Male,
			Birthday:    &now,
			Status:      status.Enable,
			Description: re.Desc(),
		}
		modelList = append(modelList, data)
	}

	newModelList := insertIfNotExist[modelUser.User](modelList, func(model modelUser.User) (*modelUser.User, error) {
		return r.Q.User.Where(r.Q.User.ID.Eq(model.ID)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.User.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[用户表]初始化失败！")
	} else {
		log.Info().Msgf("[用户表]初始化成功，共%d条数据！", len(newModelList))
	}
}
