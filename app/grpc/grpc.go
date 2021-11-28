package http

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"

	clientHTTP "github.com/polapolo/omdbapp/internal/client/http"
	grpcHandler "github.com/polapolo/omdbapp/internal/delivery/grpc"
	"github.com/polapolo/omdbapp/internal/delivery/grpc/proto"
	"github.com/polapolo/omdbapp/internal/repository"
	"github.com/polapolo/omdbapp/internal/service"
	"github.com/polapolo/omdbapp/internal/usecase"
	"google.golang.org/grpc"
)

// InitGRPCServer -> return grpc server to serve
func InitGRPCServer(db *sqlx.DB) *grpc.Server {
	// init http clients
	httpClient := &http.Client{}
	httpClient.Transport = newrelic.NewRoundTripper(httpClient.Transport) // https://github.com/newrelic/go-agent/blob/master/GUIDE.md#external-segments
	omdbAPIHttpClient := clientHTTP.NewOMDBApiHTTPClient(httpClient, "http://www.omdbapi.com/", "faf7e5bb")

	// init repositories
	omdbAPIRepository := repository.NewOMDBApiRepository(omdbAPIHttpClient)
	searchRepository := repository.NewSearchRepository(db)

	// init services
	omdbService := service.NewOMDBService(omdbAPIRepository, searchRepository)

	// init usecases
	omdbUsecase := usecase.NewOMDBUsecase(omdbService)

	// init handler / deliveries
	omdbGRPCService := grpcHandler.NewOMDBGRPCService(omdbUsecase)

	// init new relic
	newRelicApp, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("omdbapp-grpc"),
		newrelic.ConfigLicense(""),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	// init server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(nrgrpc.UnaryServerInterceptor(newRelicApp)),
		grpc.StreamInterceptor(nrgrpc.StreamServerInterceptor(newRelicApp)),
	)
	// register services to server
	proto.RegisterOMDBServer(grpcServer, &omdbGRPCService)

	return grpcServer
}
