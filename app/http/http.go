package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	clientHTTP "github.com/polapolo/omdbapp/internal/client/http"
	"github.com/polapolo/omdbapp/internal/constant"
	httpHandler "github.com/polapolo/omdbapp/internal/delivery/http"
	"github.com/polapolo/omdbapp/internal/repository"
	"github.com/polapolo/omdbapp/internal/service"
	"github.com/polapolo/omdbapp/internal/usecase"
)

// InitHTTPServer -> return http server to serve
func InitHTTPServer(db *sqlx.DB) *http.Server {
	// init http clients
	httpClient := &http.Client{}
	httpClient.Transport = newrelic.NewRoundTripper(httpClient.Transport)
	omdbAPIHttpClient := clientHTTP.NewOMDBApiHTTPClient(httpClient, "http://www.omdbapi.com/", "faf7e5bb")

	// init repositories
	omdbAPIRepository := repository.NewOMDBApiRepository(omdbAPIHttpClient)
	searchRepository := repository.NewSearchRepository(db)

	// init services
	omdbService := service.NewOMDBService(omdbAPIRepository, searchRepository)

	// init usecases
	omdbUsecase := usecase.NewOMDBUsecase(omdbService)

	// init handler / deliveries
	omdbHandler := httpHandler.NewOMDBHttpHandler(omdbUsecase)

	// init new relic
	newRelicApp, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("omdbapp-http"),
		newrelic.ConfigLicense(""),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	// init routes
	r := gin.Default()
	r.Use(nrgin.Middleware(newRelicApp))
	r.Use(middleware(newRelicApp))
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

func middleware(app *newrelic.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		if app != nil {
			// inject new relic transaction into request's context
			txnUncasted, txnExists := c.Get(constant.GinTransactionContextKey)
			if txnExists {
				txn, ok := txnUncasted.(*newrelic.Transaction)
				if ok {
					ctxNewRelic := newrelic.NewContext(c.Request.Context(), txn)
					c.Request = c.Request.WithContext(ctxNewRelic)
				}
			}
		}
		c.Next()
	}
}
