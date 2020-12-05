package http

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/polapolo/omdbapp/internal/usecase"
)

//go:generate mockgen -source=omdb.go -destination=../../mock_deliver_http_provider/mock_omdb.go -package=mock_deliver_http_provider

// OMDBUsecaseInterface -> abstraction for OMDBUsecase (Dependency injection, for unit test)
type OMDBUsecaseInterface interface {
	Search(ctx context.Context, keyword string, page int) (usecase.OMDBSearchResponse, error)
	GetByID(ctx context.Context, imdbID string) (usecase.OMDBGetByIDResponse, error)
}

// OMDBHttpHandler -> OMDBHttpHandler object
type OMDBHttpHandler struct {
	omdbUsecase OMDBUsecaseInterface
}

// NewOMDBHttpHandler -> Create new OMDBHttpHandler object
func NewOMDBHttpHandler(omdbUsecase OMDBUsecaseInterface) OMDBHttpHandler {
	return OMDBHttpHandler{
		omdbUsecase: omdbUsecase,
	}
}

func (h OMDBHttpHandler) Search(c *gin.Context) {
	// set default payload if empty
	keyword := c.DefaultQuery("keyword", "")
	pageStr := c.DefaultQuery("page", "1")

	// validate payload
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(200, gin.H{
			"data":  []struct{}{},
			"error": "Invalid parameter",
		})
		return
	}

	// get search result
	response, err := h.omdbUsecase.Search(c.Request.Context(), keyword, page)
	if err != nil {
		c.JSON(200, gin.H{
			"data":  []struct{}{},
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data":  response,
		"error": "",
	})
	return
}

func (h OMDBHttpHandler) GetByID(c *gin.Context) {
	// set default payload if empty
	id := c.DefaultQuery("id", "")

	// get search result
	response, err := h.omdbUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(200, gin.H{
			"data":  []struct{}{},
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data":  response,
		"error": "",
	})
	return
}
