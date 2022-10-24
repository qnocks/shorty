package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	// TODO: swagger init

	return router
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
