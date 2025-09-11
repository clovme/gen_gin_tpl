package initdata

import (
	"fmt"
	modelAuth "gen_gin_tpl/internal/models/auth"
	"gen_gin_tpl/pkg/crypto"
	"gen_gin_tpl/pkg/enums/perm"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"strings"
)

func (r *InitData) Permission() {
	modelList := make([]modelAuth.Permission, 0)

	// 遍历收集所有 URI
	for i, route := range r.Router {
		// 排除静态资源文件
		if strings.HasSuffix(route.Path, "*filepath") {
			continue
		}

		// 转换为小写
		name := strings.Split(utils.CamelToSnake(route.Name), "_")
		typ := name[len(name)-1]

		tempCode := fmt.Sprintf("%s-%s-%s-%s-%s", route.Method, route.Path, route.Name, typ, route.Description)
		tempCode = strings.ToLower(crypto.Encryption(tempCode))
		modelList = append(modelList, modelAuth.Permission{
			Name:        route.Name,
			Code:        tempCode,
			PID:         0,
			Type:        perm.Code(typ),
			Uri:         route.Path,
			Method:      route.Method,
			Sort:        i + 1,
			Description: route.Description,
		})
	}

	newModelList := insertIfNotExist[modelAuth.Permission](modelList, func(model modelAuth.Permission) (*modelAuth.Permission, error) {
		return r.Q.Permission.Where(r.Q.Permission.Method.Eq(model.Method), r.Q.Permission.Uri.Eq(model.Uri)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := r.Q.Permission.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[权限表]初始化失败！")
	} else {
		log.Info().Msgf("[权限表]初始化成功，共%d条数据！", len(newModelList))
	}
}
