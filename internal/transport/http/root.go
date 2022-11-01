package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Redirect
// @Tags redirect
// @Description Redirect to original url by the short one
// @Accept  json
// @Produce  json
// @Param url path string true "url to redirect"
// @Success 307
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /{url} [get]
func (h *Handler) redirect(c *gin.Context) {
	shortURL := c.Params.ByName("url")
	if len(shortURL) == 0 {
		buildErrorResponse(c, http.StatusBadRequest, "missing [url] parameter")
	}

	link, err := h.service.Link.ProcessURLClick(shortURL)
	if err != nil {
		buildErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, link.OriginURL)
}
