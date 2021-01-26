package main

//运行项目后访问swagger的路径
//http://localhost:8080/swagger/index.html
import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/razeencheng/demo-go/swaggo-gin/docs"
)

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.

// @host 127.0.0.1:8080
// @BasePath /api/v1/
func main() {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
