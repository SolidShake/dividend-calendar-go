package handler

import (
	"github.com/SolidShake/dividend-calendar-go/pkg/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// gin-swagger middleware
// swagger embed files

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.signUp)
		auth.POST("/confirm", h.confirm)
		auth.POST("/login", h.signIn)
	}

	//api := router.Group("/api")
	return router
}
