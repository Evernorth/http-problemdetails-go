package problemdetails

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

const (
	internalServerErrorType  = "urn:problems:internal-server-error"
	internalServerErrorTitle = "An unexpected error occurred."
	badRequestType           = "urn:problems:bad-request"
	badRequestTitle          = "The provided request could not be processed because it is invalid."
	notFoundType             = "urn:problems:not-found"
	notFoundTitle            = "The specified resource could not be found."
	uuidUrnPrefix            = "urn:uuid:"
)

// ProblemDetails represents a HTTP Problem Details structure from https://www.rfc-editor.org/rfc/rfc9457.html
type ProblemDetails struct {
	Type       string         `json:"type"`
	Status     int            `json:"status"`
	Title      string         `json:"title,omitempty"`
	Detail     string         `json:"detail,omitempty"`
	Instance   string         `json:"instance"`
	Created    string         `json:"created"`
	Extensions map[string]any `json:"extensions,omitempty"`
}

// WithTitle sets a custom Title value on the ProblemDetails object. The ProblemDetails pointer is
// returned only for developer convenience.
func (p *ProblemDetails) WithTitle(title string) *ProblemDetails {
	p.Title = title
	return p
}

// WithDetail sets a custom Detail value on the ProblemDetails object. The ProblemDetails pointer is
// returned only for developer convenience.
func (p *ProblemDetails) WithDetail(detail string) *ProblemDetails {
	p.Detail = detail
	return p
}

// WithExtension adds a custom Extension value on the ProblemDetails object. The ProblemDetails pointer is
// returned only for developer convenience.
func (p *ProblemDetails) WithExtension(name string, value any) *ProblemDetails {
	if p.Extensions == nil {
		p.Extensions = make(map[string]any)
	}
	p.Extensions[name] = value
	return p //&p recreates pointer?
}

// New creates a new ProblemDetails instance. Do not use this function unless you are creating a
// custom ProblemDetails instance. Prefer using the functions that create standard ProblemDetails instances.
func New(typeUri string, status int) *ProblemDetails {
	return &ProblemDetails{
		Type:     typeUri,
		Status:   status,
		Instance: uuidUrnPrefix + uuid.NewString(),
		Created:  time.Now().UTC().Format(time.RFC3339Nano),
	}
}

// NewInternalServerError creates a new ProblemDetails instance that represents an HTTP 500 Internal Server Error.
func NewInternalServerError() *ProblemDetails {
	return New(internalServerErrorType,
		http.StatusInternalServerError).WithTitle(internalServerErrorTitle)
}

// NewBadRequest creates a new ProblemDetails instance that represents an HTTP 400 Bad Request.
func NewBadRequest() *ProblemDetails {
	return New(badRequestType,
		http.StatusBadRequest).WithTitle(badRequestTitle)
}

// NewNotFound creates a new ProblemDetails instance that represents an HTTP 404 Not Found.
func NewNotFound() *ProblemDetails {
	return New(notFoundType,
		http.StatusNotFound).WithTitle(notFoundTitle)
}
