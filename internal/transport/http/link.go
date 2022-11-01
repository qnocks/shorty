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

// @Summary Create short link
// @Tags links
// @Description Generate short url for the given one
// @Accept  json
// @Produce  json
// @Param input body createLinkInput true "origin url to short"
// @Success 201 {object} linkResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/links/ [post]
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

// @Summary Get link info
// @Tags links
// @Description Get short url and redirect counts for the url
// @Accept  json
// @Produce  json
// @Param url path string true "url"
// @Success 200 {object} linkResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/links/{url} [get]
func (h *Handler) getLinkInfo(c *gin.Context) {
	originalUrl := c.Params.ByName("url")
	if len(originalUrl) == 0 {
		buildErrorResponse(c, http.StatusBadRequest, "missing [url] parameter")
	}

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
