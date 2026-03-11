package handler

import (
	"bufio"
	"log/slog"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunCommand(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Error("websocket upgrade failed", "error", err)
		return
	}
	defer conn.Close()

	cmd := exec.Command("docker", "ps")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		slog.Error("stdout pipe failed", "error", err)
		return
	}
	if err := cmd.Start(); err != nil {
		slog.Error("cmd start failed", "error", err)
		return
	}

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		if err := conn.WriteMessage(websocket.TextMessage, scanner.Bytes()); err != nil {
			return
		}
	}

	cmd.Wait()
}
