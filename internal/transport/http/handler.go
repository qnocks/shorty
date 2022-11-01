package http

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	_ "shorty/docs"
	"shorty/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	h.initAPI(router)
	h.healthcheck(router)
	h.initSwagger(router)

	return router
}

func (h *Handler) initSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (h *Handler) initAPI(router *gin.Engine) {
	root := router.Group("/")
	{
		root.GET("/:url", h.redirect)
	}

	link := router.Group("/api/links")
	{
		link.POST("/", h.createShortURL)
		link.GET("/:url", h.getLinkInfo)
	}
}

func (h *Handler) healthcheck(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
}
