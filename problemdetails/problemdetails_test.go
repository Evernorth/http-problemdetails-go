package problemdetails

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestNewInternalServerError(t *testing.T) {
	problemDetails := NewInternalServerError()
	assert.Equal(t, internalServerError.urn, problemDetails.Type)
	assert.Equal(t, internalServerError.code, problemDetails.Status)
	assert.Equal(t, internalServerError.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotFound(t *testing.T) {
	problemDetails := NewNotFound()
	assert.Equal(t, notFound.urn, problemDetails.Type)
	assert.Equal(t, notFound.code, problemDetails.Status)
	assert.Equal(t, notFound.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewMethodNotAllowed(t *testing.T) {
	problemDetails := NewMethodNotAllowed()
	assert.Equal(t, methodNotAllowed.urn, problemDetails.Type)
	assert.Equal(t, methodNotAllowed.code, problemDetails.Status)
	assert.Equal(t, methodNotAllowed.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadRequest(t *testing.T) {
	problemDetails := NewBadRequest()
	assert.Equal(t, badRequest.urn, problemDetails.Type)
	assert.Equal(t, badRequest.code, problemDetails.Status)
	assert.Equal(t, badRequest.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUnauthorized(t *testing.T) {
	problemDetails := NewUnauthorized()
	assert.Equal(t, unauthorized.urn, problemDetails.Type)
	assert.Equal(t, unauthorized.code, problemDetails.Status)
	assert.Equal(t, unauthorized.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewForbidden(t *testing.T) {
	problemDetails := NewForbidden()
	assert.Equal(t, forbidden.urn, problemDetails.Type)
	assert.Equal(t, forbidden.code, problemDetails.Status)
	assert.Equal(t, forbidden.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewConflict(t *testing.T) {
	problemDetails := NewConflict()
	assert.Equal(t, conflict.urn, problemDetails.Type)
	assert.Equal(t, conflict.code, problemDetails.Status)
	assert.Equal(t, conflict.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadGateway(t *testing.T) {
	problemDetails := NewBadGateway()
	assert.Equal(t, badGateway.urn, problemDetails.Type)
	assert.Equal(t, badGateway.code, problemDetails.Status)
	assert.Equal(t, badGateway.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewServiceUnavailable(t *testing.T) {
	problemDetails := NewServiceUnavailable()
	assert.Equal(t, serviceUnavailable.urn, problemDetails.Type)
	assert.Equal(t, serviceUnavailable.code, problemDetails.Status)
	assert.Equal(t, serviceUnavailable.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewGatewayTimeout(t *testing.T) {
	problemDetails := NewGatewayTimeout()
	assert.Equal(t, gatewayTimeout.urn, problemDetails.Type)
	assert.Equal(t, gatewayTimeout.code, problemDetails.Status)
	assert.Equal(t, gatewayTimeout.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewTooManyRequests(t *testing.T) {
	problemDetails := NewTooManyRequests()
	assert.Equal(t, tooManyRequests.urn, problemDetails.Type)
	assert.Equal(t, tooManyRequests.code, problemDetails.Status)
	assert.Equal(t, tooManyRequests.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewWithDetail(t *testing.T) {
	const detail = "This is a detailed error message"

	problemDetails := NewInternalServerError().WithDetail(detail)
	assert.Equal(t, internalServerError.urn, problemDetails.Type)
	assert.Equal(t, internalServerError.code, problemDetails.Status)
	assert.Equal(t, internalServerError.title, problemDetails.Title)

}
func TestNewWithExtension(t *testing.T) {
	const key = "example1"
	const value = "test"

	problemDetails := NewInternalServerError().WithExtension(key, value)
	assert.Equal(t, internalServerError.urn, problemDetails.Type)
	assert.Equal(t, internalServerError.code, problemDetails.Status)
	assert.Equal(t, internalServerError.title, problemDetails.Title)
	assert.Equal(t, 1, len(problemDetails.Extensions))
	assert.Equal(t, value, problemDetails.Extensions[key])
	doCommonAssertions(t, problemDetails)
}

func doCommonAssertions(t *testing.T, problemDetails *ProblemDetails) {
	assert.Equal(t, true, strings.Contains(problemDetails.Instance, uuidUrnPrefix))

	var createdTime time.Time
	var err error
	createdTime, err = time.Parse(time.RFC3339Nano, problemDetails.Created)
	assert.Nil(t, err)
	assert.NotNil(t, createdTime)
}
