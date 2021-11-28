package http

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/polapolo/omdbapp/internal/constant"
	"github.com/polapolo/omdbapp/internal/lib/distributedtracing"
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

// MovieSearch -> http handler to search for movies by keyword and page
func (h OMDBHttpHandler) MovieSearch(c *gin.Context) {
	ctxSegment, tracerSpan := distributedtracing.StartSegment(c.Request.Context(), constant.SegmentHTTPHandler+"MovieSearch")
	defer tracerSpan.End()

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
	response, err := h.omdbUsecase.Search(ctxSegment, keyword, page)
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

// MovieDetail -> http handler to search for movie by imdb id
func (h OMDBHttpHandler) MovieDetail(c *gin.Context) {
	ctxSegment, tracerSpan := distributedtracing.StartSegment(c.Request.Context(), constant.SegmentHTTPHandler+"MovieDetail")
	defer tracerSpan.End()

	// get search result
	response, err := h.omdbUsecase.GetByID(ctxSegment, c.Param("id"))
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
