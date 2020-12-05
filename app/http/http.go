package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	httpClient "github.com/polapolo/omdbapp/internal/client/http"
	httpHandler "github.com/polapolo/omdbapp/internal/delivery/http"
	"github.com/polapolo/omdbapp/internal/repository"
	"github.com/polapolo/omdbapp/internal/service"
	"github.com/polapolo/omdbapp/internal/usecase"
)

// InitHTTPServer -> return http server to serve
func InitHTTPServer(db *sqlx.DB) *http.Server {
	// init http clients
	omdbAPIHttpClient := httpClient.NewOMDBApiHTTPClient("http://www.omdbapi.com/", "faf7e5bb")

	// init repositories
	omdbAPIRepository := repository.NewOMDBApiRepository(omdbAPIHttpClient)

	// init services
	omdbService := service.NewOMDBService(omdbAPIRepository)

	// init usecases
	omdbUsecase := usecase.NewOMDBUsecase(omdbService)

	// init handler / deliveries
	omdbHandler := httpHandler.NewOMDBHttpHandler(omdbUsecase)

	// init routes
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) { omdbHandler.MovieSearch(c) })
	r.GET("/detail/:id", func(c *gin.Context) { omdbHandler.MovieDetail(c) })

	// server
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
