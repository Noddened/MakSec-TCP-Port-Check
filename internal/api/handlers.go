package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Status string `json:"status"`
}

type Request struct {
	IP   string `json:"ip" binding:"required"`
	Port int    `json:"port" binding:"required"`
}

func HandleCheck(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	addr := fmt.Sprintf("%s:%d", req.IP, req.Port)
	conn, err := d.DialContext(ctx, "tcp", addr)
	status := "closed"
	if err == nil {
		status = "open"
		conn.Close()
	}

	resp := Response{
		IP:     req.IP,
		Port:   req.Port,
		Status: status,
	}
	c.JSON(http.StatusOK, resp)
}
