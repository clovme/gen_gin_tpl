package utils

import (
	"gen_gin_tpl/pkg/snowflake"
	"github.com/rs/zerolog/log"
)

var node *snowflake.Node

// InitSnowflake 初始化雪花算法节点
// 参数：
//   - nodeID: 节点ID
//
// 返回值：
//   - error: 错误信息
func InitSnowflake(nodeID int64) {
	var err error
	node, err = snowflake.NewNode(nodeID)
	if err != nil {
		log.Panic().Err(err).Msg("雪花算法节点初始化失败")
	}
}

// GenerateID 生成唯一ID
// 返回值：
//   - int64: 唯一ID
func GenerateID() int64 {
	return node.Generate().Int64()
}
