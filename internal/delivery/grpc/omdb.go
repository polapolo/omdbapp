package grpc

import (
	"context"

	"github.com/polapolo/omdbapp/internal/delivery/grpc/proto"
	"github.com/polapolo/omdbapp/internal/usecase"
)

//go:generate protoc --proto_path=C:\Projects\omdbapp\internal\delivery\grpc\proto --go_out=plugins=grpc:C:\Projects\omdbapp\internal\delivery\grpc\proto C:\Projects\omdbapp\internal\delivery\grpc\proto\omdb.proto
//go:generate mockgen -source=omdb.go -destination=../../mock_deliver_grpc_provider/mock_omdb.go -package=mock_deliver_grpc_provider

// OMDBUsecaseInterface -> abstraction for OMDBUsecase (Dependency injection, for unit test)
type OMDBUsecaseInterface interface {
	Search(ctx context.Context, keyword string, page int) (usecase.OMDBSearchResponse, error)
	GetByID(ctx context.Context, imdbID string) (usecase.OMDBGetByIDResponse, error)
}

// OMDBGRPCService -> implement proto.OMDBServer interface
type OMDBGRPCService struct {
	omdbUsecase OMDBUsecaseInterface
}

// NewOMDBGRPCService -> Create new OMDBGRPCService object
func NewOMDBGRPCService(omdbUsecase OMDBUsecaseInterface) OMDBGRPCService {
	return OMDBGRPCService{
		omdbUsecase: omdbUsecase,
	}
}

// MovieSearch -> GRPC handler to search for movies by keyword and page
func (h OMDBGRPCService) MovieSearch(ctx context.Context, request *proto.MovieSearchRequest) (*proto.MovieSearchResponse, error) {
	// get search result
	response, err := h.omdbUsecase.Search(ctx, request.Keyword, int(request.Page))
	if err != nil {
		return nil, err
	}

	// map result
	resultDataSearch := make([]*proto.MovieSearchResponseDataSearch, 0)
	for _, search := range response.Search {
		resultDataSearch = append(resultDataSearch, &proto.MovieSearchResponseDataSearch{
			Title:  search.Title,
			Year:   search.Year,
			ImdbId: search.ImdbID,
			Type:   search.Type,
			Poster: search.Poster,
		})
	}
	resultData := &proto.MovieSearchResponseData{
		Search:       resultDataSearch,
		TotalResults: response.TotalResults,
		Response:     response.Response,
		Error:        response.Error,
	}
	return &proto.MovieSearchResponse{
		Data:  resultData,
		Error: response.Error,
	}, nil
}

// MovieDetail -> http handler to search for movie by imdb id
func (h OMDBGRPCService) MovieDetail(ctx context.Context, request *proto.MovieDetailRequest) (*proto.MovieDetailResponse, error) {
	// get search result
	response, err := h.omdbUsecase.GetByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	// map result
	resultDataRatings := make([]*proto.MovieDetailResponseDataRating, 0)
	for _, rating := range response.Ratings {
		resultDataRatings = append(resultDataRatings, &proto.MovieDetailResponseDataRating{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}
	resultData := &proto.MovieDetailResponseData{
		Title:      response.Title,
		Year:       response.Year,
		Rated:      response.Rated,
		Released:   response.Released,
		Runtime:    response.Runtime,
		Genre:      response.Genre,
		Director:   response.Director,
		Writer:     response.Writer,
		Actors:     response.Actors,
		Plot:       response.Plot,
		Language:   response.Language,
		Country:    response.Country,
		Awards:     response.Awards,
		Poster:     response.Poster,
		Ratings:    resultDataRatings,
		Metascore:  response.Metascore,
		ImdbRating: response.ImdbRating,
		ImdbVotes:  response.ImdbVotes,
		ImdbId:     response.ImdbID,
		Type:       response.Type,
		Dvd:        response.DVD,
		BoxOffice:  response.BoxOffice,
		Production: response.Production,
		Website:    response.Website,
		Response:   response.Response,
		Error:      response.Error,
	}
	return &proto.MovieDetailResponse{
		Data:  resultData,
		Error: response.Error,
	}, nil
}
