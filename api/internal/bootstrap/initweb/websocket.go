package initweb

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/logger"
	"github.com/gorilla/websocket"
	"net/http"
)

// WebSocketHandler 处理 WebSocket 连接。
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有跨域，正式环境记得收紧
	},
}

// LogWebSocketHandler 处理 WebSocket 连接。
// 参数：
//   - c *gin.Context: Gin 上下文对象
//
// 返回值：
//   - 无
func LogWebSocketHandler(c *core.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// 添加 client
	client := logger.LogHub.AddClient(conn)

	// 保持连接，读防断线
	for {
		if _, _, err := conn.NextReader(); err != nil {
			logger.LogHub.RemoveClient(client)
			conn.Close()
			break
		}
	}
}
