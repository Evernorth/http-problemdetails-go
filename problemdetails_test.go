package problemdetails

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestNewInternalServerError(t *testing.T) {
	problemDetails := NewInternalServerError()
	assert.Equal(t, InternalServerError.urn, problemDetails.Type)
	assert.Equal(t, InternalServerError.code, problemDetails.Status)
	assert.Equal(t, InternalServerError.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotFound(t *testing.T) {
	problemDetails := NewNotFound()
	assert.Equal(t, NotFound.urn, problemDetails.Type)
	assert.Equal(t, NotFound.code, problemDetails.Status)
	assert.Equal(t, NotFound.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadRequest(t *testing.T) {
	problemDetails := NewBadRequest()
	assert.Equal(t, BadRequest.urn, problemDetails.Type)
	assert.Equal(t, BadRequest.code, problemDetails.Status)
	assert.Equal(t, BadRequest.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUnauthorized(t *testing.T) {
	problemDetails := NewUnauthorized()
	assert.Equal(t, Unauthorized.urn, problemDetails.Type)
	assert.Equal(t, Unauthorized.code, problemDetails.Status)
	assert.Equal(t, Unauthorized.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewForbidden(t *testing.T) {
	problemDetails := NewForbidden()
	assert.Equal(t, Forbidden.urn, problemDetails.Type)
	assert.Equal(t, Forbidden.code, problemDetails.Status)
	assert.Equal(t, Forbidden.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewConflict(t *testing.T) {
	problemDetails := NewConflict()
	assert.Equal(t, Conflict.urn, problemDetails.Type)
	assert.Equal(t, Conflict.code, problemDetails.Status)
	assert.Equal(t, Conflict.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotImplemented(t *testing.T) {
	problemDetails := NewNotImplemented()
	assert.Equal(t, NotImplemented.urn, problemDetails.Type)
	assert.Equal(t, NotImplemented.code, problemDetails.Status)
	assert.Equal(t, NotImplemented.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadGateway(t *testing.T) {
	problemDetails := NewBadGateway()
	assert.Equal(t, BadGateway.urn, problemDetails.Type)
	assert.Equal(t, BadGateway.code, problemDetails.Status)
	assert.Equal(t, BadGateway.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewServiceUnavailable(t *testing.T) {
	problemDetails := NewServiceUnavailable()
	assert.Equal(t, ServiceUnavailable.urn, problemDetails.Type)
	assert.Equal(t, ServiceUnavailable.code, problemDetails.Status)
	assert.Equal(t, ServiceUnavailable.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewGatewayTimeout(t *testing.T) {
	problemDetails := NewGatewayTimeout()
	assert.Equal(t, GatewayTimeout.urn, problemDetails.Type)
	assert.Equal(t, GatewayTimeout.code, problemDetails.Status)
	assert.Equal(t, GatewayTimeout.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewTooManyRequests(t *testing.T) {
	problemDetails := NewTooManyRequests()
	assert.Equal(t, TooManyRequests.urn, problemDetails.Type)
	assert.Equal(t, TooManyRequests.code, problemDetails.Status)
	assert.Equal(t, TooManyRequests.title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
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
