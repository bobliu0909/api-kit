package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type MiddlewareHandlerFunc func() gin.HandlerFunc

var middlewares = map[string]MiddlewareHandlerFunc {
	"cors": gin_cors,
	"logger": gin_logger,
}

func gin_cors() gin.HandlerFunc {
	return cors.Default()
}

func gin_logger() gin.HandlerFunc {
	return logger.SetLogger()
}
