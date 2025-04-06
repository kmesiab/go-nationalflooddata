package go_nationalflooddata_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/kmesiab/go-nationalflooddata"
)

func TestNewService_ShouldCreateServiceWithProvidedAPIKey(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	if service.APIKey != apiKey {
		t.Errorf("expected APIKey to be %s, got %s", apiKey, service.APIKey)
	}
	if service.BaseURL != "https://api.nationalflooddata.com/v3" {
		t.Errorf("expected BaseURL to be https://api.nationalflooddata.com/v3, got %s", service.BaseURL)
	}
	if service.HTTPClient != http.DefaultClient {
		t.Error("expected HTTPClient to be http.DefaultClient")
	}
}

func TestNewService_ShouldSetBaseURLToDefaultAPIEndpoint(t *testing.T) {
	apiKey := "dummy-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	expectedBaseURL := "https://api.nationalflooddata.com/v3"
	if service.BaseURL != expectedBaseURL {
		t.Errorf("expected BaseURL to be %s, got %s", expectedBaseURL, service.BaseURL)
	}
}

func TestNewService_ShouldUseDefaultHTTPClient(t *testing.T) {
	apiKey := "another-test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	if service.HTTPClient != http.DefaultClient {
		t.Error("expected HTTPClient to be http.DefaultClient")
	}
}

func TestNewService_ShouldHandleEmptyAPIKey(t *testing.T) {
	apiKey := ""
	service := go_nationalflooddata.NewService(apiKey)

	if service.APIKey != apiKey {
		t.Errorf("expected APIKey to be an empty string, got %s", service.APIKey)
	}
	if service.BaseURL != "https://api.nationalflooddata.com/v3" {
		t.Errorf("expected BaseURL to be https://api.nationalflooddata.com/v3, got %s", service.BaseURL)
	}
	if service.HTTPClient != http.DefaultClient {
		t.Error("expected HTTPClient to be http.DefaultClient")
	}
}

func TestNewService_ShouldHandleVeryLongAPIKey(t *testing.T) {
	longAPIKey := "a" + strings.Repeat("b", 1000) + "c"
	service := go_nationalflooddata.NewService(longAPIKey)

	if service.APIKey != longAPIKey {
		t.Errorf("expected APIKey to be %s, got %s", longAPIKey, service.APIKey)
	}
	if service.BaseURL != "https://api.nationalflooddata.com/v3" {
		t.Errorf("expected BaseURL to be https://api.nationalflooddata.com/v3, got %s", service.BaseURL)
	}
	if service.HTTPClient != http.DefaultClient {
		t.Error("expected HTTPClient to be http.DefaultClient")
	}
}

func TestNewService_ShouldCreateNewInstanceEachTime(t *testing.T) {
	apiKey1 := "api-key-1"
	apiKey2 := "api-key-2"

	service1 := go_nationalflooddata.NewService(apiKey1)
	service2 := go_nationalflooddata.NewService(apiKey2)

	if service1 == service2 {
		t.Error("expected NewService to create a new instance each time, but got the same instance")
	}

	if service1.APIKey == service2.APIKey {
		t.Errorf("expected different API keys, got %s and %s", service1.APIKey, service2.APIKey)
	}
}

func TestNewService_ShouldNotModifyDefaultHTTPClient(t *testing.T) {
	originalClient := http.DefaultClient
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	if service.HTTPClient != originalClient {
		t.Error("expected HTTPClient to be the same as the original http.DefaultClient")
	}
}

func TestNewService_ShouldEnsureBaseURLIsNotNil(t *testing.T) {
	apiKey := "test-api-key"
	service := go_nationalflooddata.NewService(apiKey)

	if service.BaseURL == "" {
		t.Error("expected BaseURL to be non-nil, got an empty string")
	}
}
