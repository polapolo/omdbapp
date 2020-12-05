package http

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	resty "github.com/go-resty/resty/v2"
)

// Docs: http://www.omdbapi.com/

// OMDBApiHTTPClient -> OMDBApiHTTPClient object
type OMDBApiHTTPClient struct {
	Host   string
	APIKey string
}

// NewOMDBApiHTTPClient -> Create new OMDBApiHTTPClient object
func NewOMDBApiHTTPClient(host string, apiKey string) OMDBApiHTTPClient {
	return OMDBApiHTTPClient{
		Host:   host,
		APIKey: apiKey,
	}
}

// Search -> HTTP Client hit omdbapi to search movie based on keyword and pagination
func (c OMDBApiHTTPClient) Search(ctx context.Context, keyword string, page int) (OMDBSearchResponse, error) {
	var result OMDBSearchResponse

	httpClient := resty.New()
	response, err := httpClient.R().
		SetQueryParams(map[string]string{
			"apikey": c.APIKey,
			"s":      keyword,
			"page":   strconv.Itoa(page),
		}).
		EnableTrace().
		Get(c.Host)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return result, err
	}

	if result.Response != "False" {
		return result, errors.New(result.Error)
	}

	return result, nil
}

// GetByID -> HTTP Client hit omdbapi to get movie detail based on IMDb ID
func (c OMDBApiHTTPClient) GetByID(ctx context.Context, imdbID string) (OMDBGetByIDResponse, error) {
	var result OMDBGetByIDResponse

	httpClient := resty.New()
	response, err := httpClient.R().
		SetQueryParams(map[string]string{
			"apikey": c.APIKey,
			"i":      imdbID,
		}).
		EnableTrace().
		Get(c.Host)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return result, err
	}

	if result.Response != "False" {
		return result, errors.New(result.Error)
	}

	return result, nil
}
