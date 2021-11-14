package app

import (
	"fmt"
	handler "mini-merchant-service/handler"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(handler.CORSMiddleware())
	RegisterApi(router)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
