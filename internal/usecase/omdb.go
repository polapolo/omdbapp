package usecase

import (
	"context"

	"github.com/polapolo/omdbapp/internal/constant"
	"github.com/polapolo/omdbapp/internal/lib/distributedtracing"
	"github.com/polapolo/omdbapp/internal/service"
)

//go:generate mockgen -source=omdb.go -destination=../mock_usecase_provider/mock_omdb.go -package=mock_usecase_provider

// OMDBServiceInterface -> abstraction for OMDBService (Dependency injection, for unit test)
type OMDBServiceInterface interface {
	Search(ctx context.Context, keyword string, page int) (service.OMDBSearchResponse, error)
	GetByID(ctx context.Context, imdbID string) (service.OMDBGetByIDResponse, error)
}

// OMDBUsecase -> OMDBUsecase object
type OMDBUsecase struct {
	omdbService OMDBServiceInterface
}

// NewOMDBUsecase -> Create new OMDBUsecase object
func NewOMDBUsecase(omdbService OMDBServiceInterface) OMDBUsecase {
	return OMDBUsecase{
		omdbService: omdbService,
	}
}

// Search -> HTTP Client hit omdbapi to search movie based on keyword and pagination
func (u OMDBUsecase) Search(ctx context.Context, keyword string, page int) (OMDBSearchResponse, error) {
	ctxSegment, tracerSpan := distributedtracing.StartSegment(ctx, constant.SegmentUsecase+"Search")
	defer tracerSpan.End()

	var result OMDBSearchResponse

	response, err := u.omdbService.Search(ctxSegment, keyword, page)
	if err != nil {
		return result, err
	}

	// map response
	searchResult := make([]OMDBSearchResponseSearch, 0)
	for _, movie := range response.Search {
		searchResult = append(searchResult, OMDBSearchResponseSearch(movie))
	}
	result = OMDBSearchResponse{
		Search:       searchResult,
		TotalResults: response.TotalResults,
		Response:     response.Response,
		Error:        response.Error,
	}

	return result, nil
}

// GetByID -> HTTP Client hit omdbapi to get movie detail based on IMDb ID
func (u OMDBUsecase) GetByID(ctx context.Context, imdbID string) (OMDBGetByIDResponse, error) {
	ctxSegment, tracerSpan := distributedtracing.StartSegment(ctx, constant.SegmentUsecase+"GetByID")
	defer tracerSpan.End()

	var result OMDBGetByIDResponse

	response, err := u.omdbService.GetByID(ctxSegment, imdbID)
	if err != nil {
		return result, err
	}

	// map response
	ratings := make([]OMDBGetByIDResponseRating, 0)
	for _, rating := range response.Ratings {
		ratings = append(ratings, OMDBGetByIDResponseRating(rating))
	}
	result = OMDBGetByIDResponse{
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
		Ratings:    ratings,
		Metascore:  response.Metascore,
		ImdbRating: response.ImdbRating,
		ImdbVotes:  response.ImdbVotes,
		ImdbID:     response.ImdbID,
		Type:       response.Type,
		DVD:        response.DVD,
		BoxOffice:  response.BoxOffice,
		Production: response.Production,
		Website:    response.Website,
		Response:   response.Response,
		Error:      response.Error,
	}

	return result, nil
}
