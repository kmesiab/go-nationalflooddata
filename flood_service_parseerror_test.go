package go_nationalflooddata_test

import (
	"errors"
	"testing"

	"github.com/kmesiab/go-nationalflooddata"
	"github.com/kmesiab/go-nationalflooddata/client"
)

func TestParseError_ShouldReturnInvalidRequestErrorWhenStatusIs400(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status: 400,
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	var invalidRequestErr *client.InvalidRequestError
	if !errors.As(err, &invalidRequestErr) {
		t.Errorf("expected error to be of type InvalidRequestError, got %T", err)
	}
}

func TestParseError_ShouldReturnAuthenticationErrorWhenStatusIs401(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status: 401,
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	var authErr *client.AuthenticationError
	if !errors.As(err, &authErr) {
		t.Errorf("expected error to be of type AuthenticationError, got %T", err)
	}
}

func TestParseError_ShouldReturnNoDataAvailableErrorWhenStatusIs402(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status: 402,
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	var noDataErr *client.NoDataAvailableError
	if !errors.As(err, &noDataErr) {
		t.Errorf("expected error to be of type NoDataAvailableError, got %T", err)
	}
}

func TestParseError_ShouldReturnLocationNotFoundErrorWhenStatusIs404(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status: 404,
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	var locationNotFoundErr *client.LocationNotFoundError
	if !errors.As(err, &locationNotFoundErr) {
		t.Errorf("expected error to be of type LocationNotFoundError, got %T", err)
	}
}

func TestParseError_ShouldReturnParcelNotFoundErrorWhenStatusIs405(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status: 405,
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	var parcelNotFoundErr *client.ParcelNotFoundError
	if !errors.As(err, &parcelNotFoundErr) {
		t.Errorf("expected error to be of type ParcelNotFoundError, got %T", err)
	}
}

func TestParseError_ShouldReturnInternalServerErrorWhenStatusIs500(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status: 500,
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	var internalServerErr *client.InternalServerError
	if !errors.As(err, &internalServerErr) {
		t.Errorf("expected error to be of type InternalServerError, got %T", err)
	}
}

func TestParseError_ShouldReturnOriginalErrorResponseWhenStatusIs403(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status:  403,
		Message: "Forbidden",
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	if err != errorResponse {
		t.Errorf("expected original ErrorResponse to be returned, got %T", err)
	}
}

func TestParseError_ShouldReturnOriginalErrorResponseWhenStatusIs418(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status:  418,
		Message: "I'm a teapot",
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	if err != errorResponse {
		t.Errorf("expected original ErrorResponse to be returned, got %T", err)
	}
}

func TestParseError_ShouldReturnOriginalErrorResponseWhenStatusIs503(t *testing.T) {
	errorResponse := &client.ErrorResponse{
		Status:  503,
		Message: "Service Unavailable",
	}

	err := go_nationalflooddata.ParseError(errorResponse)

	if err != errorResponse {
		t.Errorf("expected original ErrorResponse to be returned, got %T", err)
	}
}
