package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createLinkInput struct {
	OriginURL string `json:"originUrl" binding:"required"`
}

type linkResponse struct {
	OriginURL     string `json:"originUrl"`
	ShortURL      string `json:"shortUrl"`
	RedirectCount int    `json:"redirectCount"`
}

func (h *Handler) createShortURL(c *gin.Context) {
	var requestBody createLinkInput
	if err := c.BindJSON(&requestBody); err != nil {
		buildErrorResponse(c, http.StatusBadRequest, "error parsing request body")
		return
	}

	link, err := h.service.Link.CreateShortLink(requestBody.OriginURL)
	if err != nil {
		buildErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, linkResponse{
		OriginURL:     link.OriginURL,
		ShortURL:      link.ShortURL,
		RedirectCount: link.RedirectCount,
	})
}

func (h *Handler) getLinkInfo(c *gin.Context) {
	originalUrl := c.Params.ByName("url")

	link, err := h.service.Link.GetByShortURL(originalUrl)
	if err != nil {
		buildErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, linkResponse{
		OriginURL:     link.OriginURL,
		ShortURL:      link.ShortURL,
		RedirectCount: link.RedirectCount,
	})
}
