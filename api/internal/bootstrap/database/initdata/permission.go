package initdata

import (
	"fmt"
	modelAuth "gen_gin_tpl/internal/models/auth"
	"gen_gin_tpl/pkg/crypto"
	"gen_gin_tpl/pkg/enums/perm"
	"gen_gin_tpl/pkg/logger/log"
	"strings"
)

func (d *InitData) Permission() {
	modelList := make([]modelAuth.Permission, 0)

	// 遍历收集所有 URI
	for i, route := range d.Router {
		if strings.HasSuffix(route.Path, "*filepath") {
			continue
		}
		tempCode := fmt.Sprintf("%s-%s-%s-%s-%s", route.Method, route.Path, route.Name, route.Type, route.Description)
		tempCode = strings.ToLower(crypto.Encryption(tempCode))
		modelList = append(modelList, modelAuth.Permission{Name: route.Name, Code: tempCode, PID: 0, Type: perm.Code(route.Type), Uri: route.Path, Method: route.Method, Sort: i + 1, Description: route.Description})
	}

	newModelList := insertIfNotExist[modelAuth.Permission](modelList, func(model modelAuth.Permission) (*modelAuth.Permission, error) {
		return d.Q.Permission.Where(d.Q.Permission.Method.Eq(model.Method), d.Q.Permission.Uri.Eq(model.Uri)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := d.Q.Permission.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msgf("[%s]初始化失败:", "权限表")
	} else {
		log.Info().Msgf("[%s]初始化成功，共%d条数据！", "权限表", len(newModelList))
	}
}
