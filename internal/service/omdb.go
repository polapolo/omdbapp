package service

import (
	"context"

	"github.com/polapolo/omdbapp/internal/constant"
	"github.com/polapolo/omdbapp/internal/entity"
	"github.com/polapolo/omdbapp/internal/lib/distributedtracing"
	"github.com/polapolo/omdbapp/internal/repository"
)

//go:generate mockgen -source=omdb.go -destination=../mock_service_provider/mock_omdb.go -package=mock_service_provider

// OMDBApiRepositoryInterface -> abstraction for OMDBApiRepository (Dependency injection, for unit test)
type OMDBApiRepositoryInterface interface {
	Search(ctx context.Context, keyword string, page int) (repository.OMDBSearchResponse, error)
	GetByID(ctx context.Context, imdbID string) (repository.OMDBGetByIDResponse, error)
}

// SearchRepositoryInterface -> abstraction for SearchRepository (Dependency injection, for unit test)
type SearchRepositoryInterface interface {
	InsertSearchHistory(ctx context.Context, row entity.Search) error
}

// OMDBService -> OMDBService object
type OMDBService struct {
	omdbAPIRepository OMDBApiRepositoryInterface
	searchRepository  SearchRepositoryInterface
}

// NewOMDBService -> Create new OMDBService object
func NewOMDBService(omdbAPIRepository OMDBApiRepositoryInterface, searchRepository SearchRepositoryInterface) OMDBService {
	return OMDBService{
		omdbAPIRepository: omdbAPIRepository,
		searchRepository:  searchRepository,
	}
}

// Search -
// 1. HTTP Client hit omdbapi to search movie based on keyword and pagination
// 2. Save search history into database
func (s OMDBService) Search(ctx context.Context, keyword string, page int) (OMDBSearchResponse, error) {
	ctxSegment, tracerSpan := distributedtracing.StartSegment(ctx, constant.SegmentService+"Search")
	defer tracerSpan.End()

	var result OMDBSearchResponse

	// hit api
	response, err := s.omdbAPIRepository.Search(ctxSegment, keyword, page)
	if err != nil {
		return result, err
	}

	// insert search history to db
	err = s.searchRepository.InsertSearchHistory(ctxSegment, entity.Search{
		Keyword: keyword,
		Page:    page,
	})
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
func (s OMDBService) GetByID(ctx context.Context, imdbID string) (OMDBGetByIDResponse, error) {
	ctxSegment, tracerSpan := distributedtracing.StartSegment(ctx, constant.SegmentService+"GetByID")
	defer tracerSpan.End()

	var result OMDBGetByIDResponse

	// hit api
	response, err := s.omdbAPIRepository.GetByID(ctxSegment, imdbID)
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
