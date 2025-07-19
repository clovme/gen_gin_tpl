package utils

import (
	"context"
	"fmt"
	"gen_gin_tpl/pkg/logger/log"
	"github.com/redis/go-redis/v9"
	"time"
)

// CheckRedisConn 检查Redis连接是否正常
//
// 参数:
//   - host: Redis主机地址
//   - port: Redis端口
//   - username: Redis用户名
//   - password: Redis密码
//   - db: Redis数据库索引
//
// 返回值:
//   - error: 连接错误，如果为nil则表示连接正常
func CheckRedisConn(host string, port int, username, password string, db int) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Username: username,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return err // 直接返回ping失败的错误
	}

	if err := rdb.Close(); err != nil {
		log.Warn().Err(err).Msg("Redis 检测用连接关闭失败") // 这只是警告，不影响连接检测结果
	}

	return nil
}
