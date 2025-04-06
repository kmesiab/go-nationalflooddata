package go_nationalflooddata_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/kmesiab/go-nationalflooddata"
	"github.com/kmesiab/go-nationalflooddata/client"
)

// Do Request Tests
func TestDoRequest_ShouldReturnErrorForInvalidEndpointURL(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)
	service.BaseURL = "://invalid-url"

	ctx := context.Background()
	_, _, err := service.DoRequest(ctx, http.MethodGet, "/test-path", nil, nil)

	if err == nil {
		t.Error("expected an error for invalid endpoint URL, but got none")
	}

	expectedErrorMsg := "invalid endpoint URL"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestDoRequest_ShouldCorrectlyAttachQueryParametersToRequestURL(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	queryParams := url.Values{}
	queryParams.Set("param1", "value1")
	queryParams.Set("param2", "value2")

	// Mock the HTTP client to prevent real endpoint calls
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			expectedURL := "https://api.nationalflooddata.com/v3/test-path?param1=value1&param2=value2"
			if req.URL.String() != expectedURL {
				t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Header:     make(http.Header),
			}
		}),
	}

	_, _, err := service.DoRequest(ctx, http.MethodGet, "/test-path", queryParams, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDoRequest_ShouldHandleEmptyBodyWithoutErrors(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	queryParams := url.Values{}
	path := "/test-path"

	// Mock the HTTP client to prevent real endpoint calls
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			if req.Body != nil {
				t.Error("expected request body to be nil for empty body, but got non-nil")
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Header:     make(http.Header),
			}
		}),
	}

	_, _, err := service.DoRequest(ctx, http.MethodGet, path, queryParams, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDoRequest_ShouldReturnErrorWhenCreatingRequestFails(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	// Invalid method to trigger request creation failure
	invalidMethod := "INVALID_METHOD"
	ctx := context.Background()

	_, _, err := service.DoRequest(ctx, invalidMethod, "/test-path", nil, nil)
	if err == nil {
		t.Error("expected an error when creating request with invalid method, but got none")
	}

	expectedErrorMsg := "INVALID_METHOD"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestDoRequest_ShouldSetXApiKeyHeader(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	path := "/test-path"

	// Mock the HTTP client to prevent real endpoint calls
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			if req.Header.Get("x-api-key") != apiKey {
				t.Errorf("expected x-api-key header to be %s, got %s", apiKey, req.Header.Get("x-api-key"))
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Header:     make(http.Header),
			}
		}),
	}

	_, _, err := service.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDoRequest_ShouldSetContentTypeHeaderToApplicationJSON(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	path := "/test-path"

	// Mock the HTTP client to prevent real endpoint calls
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			if req.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type header to be application/json, got %s", req.Header.Get("Content-Type"))
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{}`)),
				Header:     make(http.Header),
			}
		}),
	}

	_, _, err := service.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDoRequest_ShouldReturnErrorWhenHTTPClientFails(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	path := "/test-path"

	// Mock the HTTP client to simulate an error
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return nil // Simulate an error by returning nil response
		}),
	}

	_, _, err := service.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err == nil {
		t.Error("expected an error when HTTP client returns an error, but got none")
	}

	expectedErrorMsg := "request error"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestDoRequest_ShouldCorrectlyReadAndReturnResponseBodyForSuccessfulRequests(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	path := "/test-path"
	expectedResponseBody := `{"key":"value"}`

	// Mock the HTTP client to prevent real endpoint calls
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(expectedResponseBody)),
				Header:     make(http.Header),
			}
		}),
	}

	raw, _, err := service.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(raw) != expectedResponseBody {
		t.Errorf("expected response body to be %s, got %s", expectedResponseBody, string(raw))
	}
}

func TestDoRequest_ShouldReturnErrorWhenReadingResponseBodyFails(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	path := "/test-path"

	// Mock the HTTP client to simulate an error when reading the response body
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(&errorReader{}), // Simulate error on read
				Header:     make(http.Header),
			}
		}),
	}

	_, _, err := service.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err == nil {
		t.Error("expected an error when reading response body fails, but got none")
	}

	expectedErrorMsg := "reading response body"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestDoRequest_ShouldParseAndReturnErrorResponseForHTTPStatusCodes400AndAbove(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	path := "/test-path"

	// Mock the HTTP client to return a 400 error response
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader(`{"message": "Bad Request"}`)),
				Header:     make(http.Header),
			}
		}),
	}

	_, resp, err := service.DoRequest(ctx, http.MethodGet, path, nil, nil)
	if err == nil {
		t.Error("expected an error for HTTP status code 400, but got none")
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code to be 400, got %d", resp.StatusCode)
	}

	var invalidRequestErr *client.InvalidRequestError
	if !errors.As(err, &invalidRequestErr) {
		t.Errorf("expected error to be of type InvalidRequestError, got %T", err)
	}
}

// errorReader is a helper type to simulate an error when reading from the response body
type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("simulated read error")
}
