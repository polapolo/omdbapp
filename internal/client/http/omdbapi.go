package http

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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

// patching
var (
	ioutilReadAll = ioutil.ReadAll
	httpGet       = http.Get
	urlParse      = url.Parse
)

// Search -> HTTP Client hit omdbapi to search movie based on keyword and pagination
func (c OMDBApiHTTPClient) Search(ctx context.Context, keyword string, page int) (OMDBSearchResponse, error) {
	var result OMDBSearchResponse

	// prepare url
	u, err := urlParse(c.Host)
	if err != nil {
		return result, err
	}
	// add query params on url
	q := u.Query()
	q.Set("apikey", c.APIKey)
	q.Set("s", keyword)
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()

	// hit
	response, err := httpGet(u.String())
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	// read response
	body, err := ioutilReadAll(response.Body)
	if err != nil {
		return result, err
	}

	// change to struct
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	// handle error
	if result.Response == "False" {
		return result, errors.New(result.Error)
	}

	return result, nil
}

// GetByID -> HTTP Client hit omdbapi to get movie detail based on IMDb ID
func (c OMDBApiHTTPClient) GetByID(ctx context.Context, imdbID string) (OMDBGetByIDResponse, error) {
	var result OMDBGetByIDResponse

	// prepare url
	u, err := urlParse(c.Host)
	if err != nil {
		return result, err
	}
	// add query params on url
	q := u.Query()
	q.Set("apikey", c.APIKey)
	q.Set("i", imdbID)
	u.RawQuery = q.Encode()

	// hit
	response, err := httpGet(u.String())
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	// read response
	body, err := ioutilReadAll(response.Body)
	if err != nil {
		return result, err
	}

	// change to struct
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	// handle error
	if result.Response == "False" {
		return result, errors.New(result.Error)
	}

	return result, nil
}
