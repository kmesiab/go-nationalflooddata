package go_nationalflooddata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/kmesiab/go-nationalflooddata/client"
)

// Service is the main client for interacting with the National Flood Data API.
type Service struct {
	// BaseURL is the base endpoint for the v3 API.
	BaseURL string

	// APIKey is your National Flood Data API key.
	APIKey string

	// HTTPClient is the underlying HTTP client used to make requests.
	// You can override this with a custom client (e.g., with custom timeouts).
	HTTPClient *http.Client
}

// NewService returns a new NFD service client initialized with the given API key.
// By default, it uses https://api.nationalflooddata.com/v3 as the BaseURL and
// http.DefaultClient for the HTTP client.
func NewService(apiKey string) *Service {
	return &Service{
		BaseURL:    "https://api.nationalflooddata.com/v3",
		APIKey:     apiKey,
		HTTPClient: http.DefaultClient,
	}
}

// DoRequest is a helper to build and execute an HTTP request, returning the raw response body
// and the *http.Response so the caller can handle status codes if necessary.
func (s *Service) DoRequest(
	ctx context.Context,
	method, path string,
	queryParams url.Values,
	body []byte,
) ([]byte, *http.Response, error) {
	// Build the full URL
	endpoint := fmt.Sprintf("%s%s", strings.TrimSuffix(s.BaseURL, "/"), path)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	// Attach query parameters
	if queryParams != nil {
		u.RawQuery = queryParams.Encode()
	}

	// Create request
	var reqBody io.Reader
	if len(body) > 0 {
		reqBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("creating request: %w", err)
	}

	// Always attach the x-api-key header
	req.Header.Set("x-api-key", s.APIKey)
	req.Header.Set("Content-Type", "application/json")

	// Execute
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("request error: %w", err)
	}

	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		// Attempt to parse the response as an error payload
		apiErr := &client.ErrorResponse{
			Response: resp,
			Status:   resp.StatusCode,
			Message:  http.StatusText(resp.StatusCode),
		}
		// Attempt to unmarshal body into that structure (to get the "message" field from JSON)
		_ = json.Unmarshal(raw, apiErr)
		return nil, resp, ParseError(apiErr)
	}

	return raw, resp, nil
}

// ParseError looks at the status code in ErrorResponse and returns the correct typed error.
func ParseError(e *client.ErrorResponse) error {

	if e == nil {
		return e
	}

	switch e.Status {
	case 400:
		return &client.InvalidRequestError{ErrorResponse: e}
	case 401:
		return &client.AuthenticationError{ErrorResponse: e}
	case 402:
		return &client.NoDataAvailableError{ErrorResponse: e}
	case 404:
		return &client.LocationNotFoundError{ErrorResponse: e}
	case 405:
		return &client.ParcelNotFoundError{ErrorResponse: e}
	case 500:
		return &client.InternalServerError{ErrorResponse: e}
	default:
		return e // fallback: return the generic *ErrorResponse
	}
}

// -----------------------------------------------------------------------------
//  GetFloodData
// -----------------------------------------------------------------------------

// GetFloodData queries the /data endpoint for FEMA Flood Data. It returns a FloodData struct.
func (s *Service) GetFloodData(ctx context.Context, opts client.FloodDataOptions) (*client.Response, error) {
	// Build query parameters from FloodDataOptions
	q := url.Values{}
	q.Set("searchtype", string(opts.SearchType))

	// Conditionally add parameters
	if opts.Address != "" {
		q.Set("address", opts.Address)
	}
	if opts.Lat != 0 {
		q.Set("lat", strconv.FormatFloat(opts.Lat, 'f', -1, 64))
	}
	if opts.Lng != 0 {
		q.Set("lng", strconv.FormatFloat(opts.Lng, 'f', -1, 64))
	}
	if opts.Polygon != "" {
		q.Set("polygon", opts.Polygon)
	}
	if opts.LOMA {
		q.Set("loma", "true")
	}
	if opts.Elevation {
		q.Set("elevation", "true")
	}
	if opts.Property {
		q.Set("property", "true")
	}
	if opts.Parcel {
		q.Set("parcel", "true")
	}

	raw, _, err := s.DoRequest(ctx, http.MethodGet, "/data", q, nil)
	if err != nil {
		return nil, err
	}

	// Clean up this garbage response
	sanitizedResponse, err := sanitizeResponse(string(raw))
	if err != nil {
		return nil, err
	}

	var fd client.Response
	if err := json.Unmarshal([]byte(sanitizedResponse), &fd); err != nil {
		return nil, fmt.Errorf("json unmarshal FloodData: %w", err)
	}

	if fd.Status == "" && fd.MatchType == nil && fd.RequestID == "" {
		return nil, fmt.Errorf("invalid response from API: no status, no matchType, or no request: %s", raw)
	}
	return &fd, nil
}

// sanitizeResponse is a helper function to sanitize the raw response from the API.
// Its very existence is a signal that you have written a mess of an API and your
// integrators are very unhappy.
func sanitizeResponse(rawResponse string) (string, error) {
	var data map[string]interface{}

	// Unmarshal the JSON string into a map
	if err := json.Unmarshal([]byte(rawResponse), &data); err != nil {

		return rawResponse, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Traverse and sanitize the map
	sanitizeMap(data)

	// Marshal the sanitized map back to a JSON string
	sanitizedJSON, err := json.Marshal(data)

	if err != nil {
		return rawResponse, fmt.Errorf("error marshalling sanitized JSON: %w", err)
	}

	return string(sanitizedJSON), nil
}

// sanitizeMap recursively traverses the map and sanitizes values
func sanitizeMap(data map[string]interface{}) {

	for key, value := range data {
		switch v := value.(type) {
		case string:

			// They literally have tons of trailing spaces on their output!
			v = strings.TrimSpace(v)

			// If you don't have certain privileges on your API key, you can't access certain
			// sections, instead of excluding it, they send a string "Access Denied".
			// God knows why.

			if v == "Access Denied" {
				// nil the original value
				data[key] = nil // Replace "Access Denied" with nil
			} else {
				data[key] = v
			}

		case map[string]interface{}:
			sanitizeMap(v) // Recursively sanitize nested maps
		case []interface{}:
			sanitizeSlice(v) // Sanitize slices
		}
	}
}

// sanitizeSlice traverses and sanitizes slices
func sanitizeSlice(data []interface{}) {
	for i, value := range data {
		switch v := value.(type) {
		case string:
			// They literally have tons of trailing spaces on their output!
			v = strings.TrimSpace(v)

			if v == "Access Denied" {
				log.Printf("Access Denied in slice at index %d\n", i)
				data[i] = nil // Replace "Access Denied" with nil
			} else {
				data[i] = v
			}

		case map[string]interface{}:
			sanitizeMap(v) // Recursively sanitize nested maps
		case []interface{}:
			sanitizeSlice(v) // Recursively sanitize nested slices
		}
	}
}

// -----------------------------------------------------------------------------
//  GetFloodMapRaw
// -----------------------------------------------------------------------------

// GetFloodMapRaw queries the /floodmapraw endpoint for the raw FEMA Flood Map polygons.
// This often returns large geojson content. The structure is defined by FloodMapContent.
func (s *Service) GetFloodMapRaw(ctx context.Context, opts client.FloodMapRawOptions) (*client.FloodMapContent, error) {
	q := url.Values{}
	q.Set("lat", strconv.FormatFloat(opts.Lat, 'f', -1, 64))
	q.Set("lng", strconv.FormatFloat(opts.Lng, 'f', -1, 64))

	// The size parameter can be 0.04, 0.06, or 0.08. Default 0.08 if omitted.
	if opts.Size != 0.0 {
		q.Set("size", strconv.FormatFloat(opts.Size, 'f', 2, 64))
	}
	q.Set("geojson", strconv.FormatBool(opts.GeoJSON))
	q.Set("excludex", strconv.FormatBool(opts.ExcludeX))
	q.Set("elevation", strconv.FormatBool(opts.Elevation))

	raw, _, err := s.DoRequest(ctx, http.MethodGet, "/floodmapraw", q, nil)
	if err != nil {
		return nil, err
	}

	var content client.FloodMapContent
	if err := json.Unmarshal(raw, &content); err != nil {
		return nil, fmt.Errorf("json unmarshal FloodMapContent: %w", err)
	}

	return &content, nil
}

// -----------------------------------------------------------------------------
//  GetFloodDataBatch
// -----------------------------------------------------------------------------

// GetFloodDataBatch posts a batch request to /databatch. It returns immediately with a
// FloodDataBatch that contains a batch_id and a URL in `Result` which you can poll.
func (s *Service) GetFloodDataBatch(ctx context.Context, batch client.BatchDataRequest) (*client.FloodDataBatch, error) {
	// The batch JSON must contain "apiKey" at the top level as the spec indicates,
	// but we also set the X-API-KEY header. Usually, these match.
	if batch.APIKey == "" {
		batch.APIKey = s.APIKey
	}

	body, err := json.Marshal(batch)
	if err != nil {
		return nil, fmt.Errorf("json marshal batch request: %w", err)
	}

	raw, _, reqErr := s.DoRequest(ctx, http.MethodPost, "/databatch", nil, body)
	if reqErr != nil {
		return nil, reqErr
	}

	var resp client.FloodDataBatch
	if err := json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal FloodDataBatch: %w", err)
	}
	return &resp, nil
}

// -----------------------------------------------------------------------------
//  GetFloodVectorTile
// -----------------------------------------------------------------------------

// GetFloodVectorTile queries the /tiles/flood-vector/{z}/{x}/{y}.mvt endpoint for flood vector tiles.
// It returns the raw tile data as a byte slice.
func (s *Service) GetFloodVectorTile(ctx context.Context, z, x, y int) ([]byte, error) {
	path := fmt.Sprintf("/tiles/flood-vector/%d/%d/%d.mvt", z, x, y)
	raw, _, err := s.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

// -----------------------------------------------------------------------------
//  GetStormSurgeTile
// -----------------------------------------------------------------------------

// GetStormSurgeTile queries the /tiles/stormsurge/{category}/{z}/{x}/{y}.png endpoint for storm surge tiles.
// It returns the raw tile data as a byte slice.
func (s *Service) GetStormSurgeTile(ctx context.Context, category string, z, x, y int) ([]byte, error) {
	path := fmt.Sprintf("/tiles/stormsurge/%s/%d/%d/%d.png", category, z, x, y)
	raw, _, err := s.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

// -----------------------------------------------------------------------------
//  GetDynamicFloodMap
// -----------------------------------------------------------------------------

// GetDynamicFloodMap queries the /dynamic.html endpoint for the dynamic flood map.
// It returns the HTML content as a string.
func (s *Service) GetDynamicFloodMap(ctx context.Context, key string, lat, lng float64, zoom int, showLegend bool) (string, error) {
	q := url.Values{}
	q.Set("key", key)
	q.Set("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	q.Set("lng", strconv.FormatFloat(lng, 'f', -1, 64))
	q.Set("zoom", strconv.Itoa(zoom))
	q.Set("showLegend", strconv.FormatBool(showLegend))

	raw, _, err := s.DoRequest(ctx, http.MethodGet, "/dynamic.html", q, nil)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// -----------------------------------------------------------------------------
//  GetStaticFloodMap
// -----------------------------------------------------------------------------

// GetStaticFloodMap queries the /staticmap endpoint for the static flood map.
// It returns the image data as a byte slice.
func (s *Service) GetStaticFloodMap(ctx context.Context, opts client.StaticMapOptions) ([]byte, error) {
	q := url.Values{}
	q.Set("lat", strconv.FormatFloat(opts.Lat, 'f', -1, 64))
	q.Set("lng", strconv.FormatFloat(opts.Lng, 'f', -1, 64))
	q.Set("height", strconv.Itoa(opts.Height))
	q.Set("width", strconv.Itoa(opts.Width))
	q.Set("showMarker", strconv.FormatBool(opts.ShowMarker))
	q.Set("showLegend", strconv.FormatBool(opts.ShowLegend))
	q.Set("zoom", strconv.Itoa(opts.Zoom))

	raw, _, err := s.DoRequest(ctx, http.MethodGet, "/staticmap", q, nil)
	if err != nil {
		return nil, err
	}
	return raw, nil
}
