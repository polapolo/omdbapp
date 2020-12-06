package http

import (
	"github.com/jmoiron/sqlx"
	httpClient "github.com/polapolo/omdbapp/internal/client/http"
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
	omdbAPIHttpClient := httpClient.NewOMDBApiHTTPClient("http://www.omdbapi.com/", "faf7e5bb")

	// init repositories
	omdbAPIRepository := repository.NewOMDBApiRepository(omdbAPIHttpClient)
	searchRepository := repository.NewSearchRepository(db)

	// init services
	omdbService := service.NewOMDBService(omdbAPIRepository, searchRepository)

	// init usecases
	omdbUsecase := usecase.NewOMDBUsecase(omdbService)

	// init handler / deliveries
	omdbGRPCService := grpcHandler.NewOMDBGRPCService(omdbUsecase)

	// init server
	grpcServer := grpc.NewServer()
	// register services to server
	proto.RegisterOMDBServer(grpcServer, &omdbGRPCService)

	return grpcServer
}
