package problemdetails

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestNewInternalServerError(t *testing.T) {
	problemDetails := NewInternalServerError()
	assert.Equal(t, internalServerErrorType, problemDetails.Type)
	assert.Equal(t, http.StatusInternalServerError, problemDetails.Status)
	assert.Equal(t, internalServerErrorTitle, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotFound(t *testing.T) {
	problemDetails := NewNotFound()
	assert.Equal(t, notFoundType, problemDetails.Type)
	assert.Equal(t, http.StatusNotFound, problemDetails.Status)
	assert.Equal(t, notFoundTitle, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadRequest(t *testing.T) {
	problemDetails := NewBadRequest()
	assert.Equal(t, badRequestType, problemDetails.Type)
	assert.Equal(t, http.StatusBadRequest, problemDetails.Status)
	assert.Equal(t, badRequestTitle, problemDetails.Title)
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
	assert.Equal(t, badRequestType, problemDetails.Type)
	assert.Equal(t, http.StatusBadRequest, problemDetails.Status)
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
