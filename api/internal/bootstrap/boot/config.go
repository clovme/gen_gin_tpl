package boot

import (
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/libs"
)

// 初始化系统配置
// 参数：
//   - query: 查询对象
//
// 返回值：
//   - 无
func initializationSystemConfig(query *query.Query) {
	configs, err := query.Config.Find()
	if err != nil {
		return
	}

	libs.UpdateWebConfig(configs)

	//if err := crypto.ParseRsaKeys(variable.PublicPEM, variable.PrivatePEM); err != nil {
	//	log.Error().Err(err).Msg("密钥初始化失败")
	//	os.Exit(-1)
	//}
}
