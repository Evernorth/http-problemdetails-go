package problemdetails

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestNewInternalServerError(t *testing.T) {
	problemDetails := NewInternalServerError()
	assert.Equal(t, InternalServerError.Type, problemDetails.Type)
	assert.Equal(t, InternalServerError.Code, problemDetails.Status)
	assert.Equal(t, InternalServerError.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotFound(t *testing.T) {
	problemDetails := NewNotFound()
	assert.Equal(t, NotFound.Type, problemDetails.Type)
	assert.Equal(t, NotFound.Code, problemDetails.Status)
	assert.Equal(t, NotFound.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadRequest(t *testing.T) {
	problemDetails := NewBadRequest()
	assert.Equal(t, BadRequest.Type, problemDetails.Type)
	assert.Equal(t, BadRequest.Code, problemDetails.Status)
	assert.Equal(t, BadRequest.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewContinue(t *testing.T) {
	problemDetails := NewContinue()
	assert.Equal(t, Continue.Type, problemDetails.Type)
	assert.Equal(t, Continue.Code, problemDetails.Status)
	assert.Equal(t, Continue.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewSwitchingProtocols(t *testing.T) {
	problemDetails := NewSwitchingProtocols()
	assert.Equal(t, SwitchingProtocols.Type, problemDetails.Type)
	assert.Equal(t, SwitchingProtocols.Code, problemDetails.Status)
	assert.Equal(t, SwitchingProtocols.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewOK(t *testing.T) {
	problemDetails := NewOK()
	assert.Equal(t, OK.Type, problemDetails.Type)
	assert.Equal(t, OK.Code, problemDetails.Status)
	assert.Equal(t, OK.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewCreated(t *testing.T) {
	problemDetails := NewCreated()
	assert.Equal(t, Created.Type, problemDetails.Type)
	assert.Equal(t, Created.Code, problemDetails.Status)
	assert.Equal(t, Created.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewAccepted(t *testing.T) {
	problemDetails := NewAccepted()
	assert.Equal(t, Accepted.Type, problemDetails.Type)
	assert.Equal(t, Accepted.Code, problemDetails.Status)
	assert.Equal(t, Accepted.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNoContent(t *testing.T) {
	problemDetails := NewNoContent()
	assert.Equal(t, NoContent.Type, problemDetails.Type)
	assert.Equal(t, NoContent.Code, problemDetails.Status)
	assert.Equal(t, NoContent.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewMovedPermanently(t *testing.T) {
	problemDetails := NewMovedPermanently()
	assert.Equal(t, MovedPermanently.Type, problemDetails.Type)
	assert.Equal(t, MovedPermanently.Code, problemDetails.Status)
	assert.Equal(t, MovedPermanently.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewFound(t *testing.T) {
	problemDetails := NewFound()
	assert.Equal(t, Found.Type, problemDetails.Type)
	assert.Equal(t, Found.Code, problemDetails.Status)
	assert.Equal(t, Found.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotModified(t *testing.T) {
	problemDetails := NewNotModified()
	assert.Equal(t, NotModified.Type, problemDetails.Type)
	assert.Equal(t, NotModified.Code, problemDetails.Status)
	assert.Equal(t, NotModified.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUnauthorized(t *testing.T) {
	problemDetails := NewUnauthorized()
	assert.Equal(t, Unauthorized.Type, problemDetails.Type)
	assert.Equal(t, Unauthorized.Code, problemDetails.Status)
	assert.Equal(t, Unauthorized.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewForbidden(t *testing.T) {
	problemDetails := NewForbidden()
	assert.Equal(t, Forbidden.Type, problemDetails.Type)
	assert.Equal(t, Forbidden.Code, problemDetails.Status)
	assert.Equal(t, Forbidden.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewMethodNotAllowed(t *testing.T) {
	problemDetails := NewMethodNotAllowed()
	assert.Equal(t, MethodNotAllowed.Type, problemDetails.Type)
	assert.Equal(t, MethodNotAllowed.Code, problemDetails.Status)
	assert.Equal(t, MethodNotAllowed.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewConflict(t *testing.T) {
	problemDetails := NewConflict()
	assert.Equal(t, Conflict.Type, problemDetails.Type)
	assert.Equal(t, Conflict.Code, problemDetails.Status)
	assert.Equal(t, Conflict.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUnprocessableEntity(t *testing.T) {
	problemDetails := NewUnprocessableEntity()
	assert.Equal(t, UnprocessableEntity.Type, problemDetails.Type)
	assert.Equal(t, UnprocessableEntity.Code, problemDetails.Status)
	assert.Equal(t, UnprocessableEntity.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotImplemented(t *testing.T) {
	problemDetails := NewNotImplemented()
	assert.Equal(t, NotImplemented.Type, problemDetails.Type)
	assert.Equal(t, NotImplemented.Code, problemDetails.Status)
	assert.Equal(t, NotImplemented.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadGateway(t *testing.T) {
	problemDetails := NewBadGateway()
	assert.Equal(t, BadGateway.Type, problemDetails.Type)
	assert.Equal(t, BadGateway.Code, problemDetails.Status)
	assert.Equal(t, BadGateway.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewServiceUnavailable(t *testing.T) {
	problemDetails := NewServiceUnavailable()
	assert.Equal(t, ServiceUnavailable.Type, problemDetails.Type)
	assert.Equal(t, ServiceUnavailable.Code, problemDetails.Status)
	assert.Equal(t, ServiceUnavailable.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadRequestCustomized(t *testing.T) {
	title := "kaboom"
	detail := "the unthinkable has occurred"
	extKey1 := "key1"
	extValue1 := "value1"
	extKey2 := "key2"
	extValue2 := "value2"

	problemDetails := NewBadRequest().WithTitle(title).WithDetail(detail).WithExtension(
		extKey1, extValue1).WithExtension(extKey2, extValue2)
	assert.Equal(t, BadRequest.Type, problemDetails.Type)
	assert.Equal(t, BadRequest.Code, problemDetails.Status)
	assert.Equal(t, title, problemDetails.Title)
	assert.Equal(t, detail, problemDetails.Detail)
	assert.Equal(t, 2, len(problemDetails.Extensions))
	assert.Equal(t, extValue1, problemDetails.Extensions[extKey1])
	assert.Equal(t, extValue2, problemDetails.Extensions[extKey2])
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
