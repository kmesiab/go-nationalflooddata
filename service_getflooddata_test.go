package go_nationalflooddata_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	go_nationalflooddata "github.com/kmesiab/go-nationalflooddata"
)

func TestGetFloodData_ShouldReturnErrorWhenDoRequestFails(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	opts := go_nationalflooddata.FloodDataOptions{
		SearchType: "address",
		Address:    "123 Test St",
	}

	// Mock the HTTP client to simulate an error
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return nil // Simulate an error by returning nil response
		}),
	}

	_, err := service.GetFloodData(ctx, opts)
	if err == nil {
		t.Error("expected an error when DoRequest fails, but got none")
	}

	expectedErrorMsg := "request error"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestGetFloodData_ShouldReturnErrorWhenSanitizeResponseFails(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	opts := go_nationalflooddata.FloodDataOptions{
		SearchType: "address",
		Address:    "123 Test St",
	}

	// Mock the HTTP client to return a malformed JSON response
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"invalid_json":`)), // Malformed JSON
				Header:     make(http.Header),
			}
		}),
	}

	_, err := service.GetFloodData(ctx, opts)
	if err == nil {
		t.Error("expected an error when sanitizeResponse fails, but got none")
	}

	expectedErrorMsg := "error unmarshalling JSON"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestGetFloodData_ShouldReturnErrorWhenJSONUnmarshalFails(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	ctx := context.Background()
	opts := go_nationalflooddata.FloodDataOptions{
		SearchType: "address",
		Address:    "123 Test St",
	}

	// Mock the HTTP client to return a valid JSON that doesn't match the expected structure
	service.HTTPClient = &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"unexpected_field": "unexpected_value"}`)),
				Header:     make(http.Header),
			}
		}),
	}

	data, err := service.GetFloodData(ctx, opts)
	if err == nil {
		t.Error("expected an error when JSON unmarshal fails, but got none")
	}

	assert.Nil(t, data)

	expectedErrorMsg := "invalid response from API: no status, no matchType, or no request: {\"unexpected_field\": \"unexpected_value\"}"
	if !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain %q, got %q", expectedErrorMsg, err.Error())
	}
}

type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip executes a single HTTP transaction and returns a Response
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
