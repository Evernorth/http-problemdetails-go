package problemdetails

import (
	"github.com/google/uuid"
	"time"
)

const (
	uuidUrnPrefix = "urn:uuid:"
	MimeType      = "application/problem+json"
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

// NewBadRequest creates a new ProblemDetails instance that represents an HTTP 400 Bad Request.
func NewBadRequest() *ProblemDetails {
	return New(BadRequest.urn,
		BadRequest.code).WithTitle(BadRequest.title)
}

// NewUnauthorized creates a new ProblemDetails instance that represents an HTTP 401 Unauthorized.
func NewUnauthorized() *ProblemDetails {
	return New(Unauthorized.urn,
		Unauthorized.code).WithTitle(Unauthorized.title)
}

// NewForbidden creates a new ProblemDetails instance that represents an HTTP 403 Forbidden.
func NewForbidden() *ProblemDetails {
	return New(Forbidden.urn,
		Forbidden.code).WithTitle(Forbidden.title)
}

// NewNotFound creates a new ProblemDetails instance that represents an HTTP 404 Not Found.
func NewNotFound() *ProblemDetails {
	return New(NotFound.urn,
		NotFound.code).WithTitle(NotFound.title)
}

// NewConflict creates a new ProblemDetails instance that represents an HTTP 409 Conflict.
func NewConflict() *ProblemDetails {
	return New(Conflict.urn,
		Conflict.code).WithTitle(Conflict.title)
}

// NewTooManyRequests creates a new ProblemDetails instance that represents an HTTP 429 Too Many Requests.
func NewTooManyRequests() *ProblemDetails {
	return New(TooManyRequests.urn,
		TooManyRequests.code).WithTitle(TooManyRequests.title)
}

// NewInternalServerError creates a new ProblemDetails instance that represents an HTTP 500 Internal Server Error.
func NewInternalServerError() *ProblemDetails {
	return New(InternalServerError.urn,
		InternalServerError.code).WithTitle(InternalServerError.title)
}

// NewNotImplemented creates a new ProblemDetails instance that represents an HTTP 501 Not Implemented.
func NewNotImplemented() *ProblemDetails {
	return New(NotImplemented.urn,
		NotImplemented.code).WithTitle(NotImplemented.title)
}

// NewBadGateway creates a new ProblemDetails instance that represents an HTTP 502 Bad Gateway.
func NewBadGateway() *ProblemDetails {
	return New(BadGateway.urn,
		BadGateway.code).WithTitle(BadGateway.title)
}

// NewServiceUnavailable creates a new ProblemDetails instance that represents an HTTP 503 Service Unavailable.
func NewServiceUnavailable() *ProblemDetails {
	return New(ServiceUnavailable.urn,
		ServiceUnavailable.code).WithTitle(ServiceUnavailable.title)
}

// NewGatewayTimeout creates a new ProblemDetails instance that represents an HTTP 504 Gateway Timeout.
func NewGatewayTimeout() *ProblemDetails {
	return New(GatewayTimeout.urn,
		GatewayTimeout.code).WithTitle(GatewayTimeout.title)
}
