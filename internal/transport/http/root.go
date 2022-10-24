package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) redirect(c *gin.Context) {
	shortURL := c.Params.ByName("url")

	link, err := h.service.Link.ProcessURLClick(shortURL)
	if err != nil {
		buildErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, link.OriginURL)
}
