package utilities

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

func NewConsoleLogger() hclog.Logger {
	logger := hclog.New(&hclog.LoggerOptions{
		Level: hclog.LevelFromString("DEBUG"),
	})
	return logger
}

func GinLogger() {
	fileName := strings.Split(time.Now().String(), " ")[0] + ".log"
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(file)
}
