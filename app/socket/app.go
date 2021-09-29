package socket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Data struct {
	Method string `json:"method"`
	Params json.RawMessage
}

func Run(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("ws链接错误")
		return
	}

	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(ws)

	for {
		data := &Data{}
		err := ws.ReadJSON(&data)

		if err != nil {
			fmt.Printf("消息格式错误：%v\n", err.Error())
			continue
		}

		switch data.Method {
		case "ping":
			Ping(ws)
			break
		default:
			fmt.Printf("不存在接口：%v\n", data.Method)
			break
		}
	}
}

// Ping 检测心跳
func Ping(ws *websocket.Conn) {
	if err := ws.WriteJSON("ping"); err != nil {
		fmt.Println(err.Error())
	}
}
