package service

import (
	"context"

	"github.com/polapolo/omdbapp/internal/repository"
)

//go:generate mockgen -source=omdb.go -destination=../mock_service_provider/mock_omdb.go -package=mock_service_provider

// OMDBApiRepositoryInterface -> abstraction for OMDBApiRepository (Dependency injection, for unit test)
type OMDBApiRepositoryInterface interface {
	Search(ctx context.Context, keyword string, page int) (repository.OMDBSearchResponse, error)
	GetByID(ctx context.Context, imdbID string) (repository.OMDBGetByIDResponse, error)
}

// OMDBService -> OMDBService object
type OMDBService struct {
	omdbRepository OMDBApiRepositoryInterface
}

// NewOMDBService -> Create new OMDBService object
func NewOMDBService(omdbAPIRepository OMDBApiRepositoryInterface) OMDBService {
	return OMDBService{
		omdbRepository: omdbAPIRepository,
	}
}

// Search -> HTTP Client hit omdbapi to search movie based on keyword and pagination
func (s OMDBService) Search(ctx context.Context, keyword string, page int) (OMDBSearchResponse, error) {
	var result OMDBSearchResponse

	// hit api
	response, err := s.omdbRepository.Search(ctx, keyword, page)
	if err != nil {
		return result, err
	}

	// map response
	searchResult := make([]OMDBSearchResponseSearch, 0)
	for _, movie := range response.Search {
		searchResult = append(searchResult, OMDBSearchResponseSearch{
			Title:  movie.Title,
			Year:   movie.Year,
			ImdbID: movie.ImdbID,
			Type:   movie.Type,
			Poster: movie.Poster,
		})
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
	var result OMDBGetByIDResponse

	// hit api
	response, err := s.omdbRepository.GetByID(ctx, imdbID)
	if err != nil {
		return result, err
	}

	// map response
	ratings := make([]OMDBGetByIDResponseRating, 0)
	for _, rating := range response.Ratings {
		ratings = append(ratings, OMDBGetByIDResponseRating{
			Source: rating.Source,
			Value:  rating.Value,
		})
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
