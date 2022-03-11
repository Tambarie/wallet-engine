package server

import (
	"fmt"
	"github.com/Tambarie/wallet-engine/application/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func initializeRouter() *gin.Engine {
	router := gin.Default()
	if os.Getenv("GIN_MODE") == "testing" {
		return router
	}

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	// setup cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	return router
}

func DefineRouter(router *gin.Engine, handler *handler.Handler) {
	apiRouter := router.Group("/api/v1")
	apiRouter.POST("/createWallet", handler.CreateWallet())
	apiRouter.POST("/creditWallet/:user-reference", handler.CreditWallet())
	apiRouter.POST("/debitWallet/:user-reference", handler.DebitWallet())
	apiRouter.PUT("/activate-deactivate/:user-reference", handler.ActivateWallet())
}
