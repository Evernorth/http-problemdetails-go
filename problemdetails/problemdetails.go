package problemdetails

import (
	"github.com/google/uuid"
	"time"
)

const (
	uuidUrnPrefix = "urn:uuid:"
	MIMETypeJSON  = "application/problem+json"
)

// Problem represents a HTTP Problem Details structure from https://www.rfc-editor.org/rfc/rfc9457.html
type Problem struct {
	Type       string         `json:"type"`
	Status     int            `json:"status"`
	Title      string         `json:"title,omitempty"`
	Detail     string         `json:"detail,omitempty"`
	Instance   string         `json:"instance"`
	Created    string         `json:"created"`
	Extensions map[string]any `json:"extensions,omitempty"`
}

// WithTitle sets a custom Title value on the Problem object. The Problem pointer is
// returned only for developer convenience.
func (p *Problem) WithTitle(title string) *Problem {
	p.Title = title
	return p
}

// WithDetail sets a custom Detail value on the Problem object. The Problem pointer is
// returned only for developer convenience.
func (p *Problem) WithDetail(detail string) *Problem {
	p.Detail = detail
	return p
}

// WithExtension adds a custom Extension value on the Problem object. The Problem pointer is
// returned only for developer convenience.
func (p *Problem) WithExtension(name string, value any) *Problem {
	if p.Extensions == nil {
		p.Extensions = make(map[string]any)
	}
	p.Extensions[name] = value
	return p //&p recreates pointer?
}

// New creates a new Problem instance. Do not use this function unless you are creating a
// custom Problem instance. Prefer using the functions that create standard Problem instances.
func New(typeUri string, status int) *Problem {
	return &Problem{
		Type:     typeUri,
		Status:   status,
		Instance: uuidUrnPrefix + uuid.NewString(),
		Created:  time.Now().UTC().Format(time.RFC3339Nano),
	}
}

// NewBadRequest creates a new Problem instance that represents an HTTP 400 Bad Request.
func NewBadRequest() *Problem {
	return New(badRequest.urn, badRequest.code).WithTitle(badRequest.title)
}

// NewUnauthorized creates a new Problem instance that represents an HTTP 401 Unauthorized.
func NewUnauthorized() *Problem {
	return New(unauthorized.urn, unauthorized.code).WithTitle(unauthorized.title)
}

// NewForbidden creates a new Problem instance that represents an HTTP 403 Forbidden.
func NewForbidden() *Problem {
	return New(forbidden.urn, forbidden.code).WithTitle(forbidden.title)
}

// NewNotFound creates a new Problem instance that represents an HTTP 404 Not Found.
func NewNotFound() *Problem {
	return New(notFound.urn, notFound.code).WithTitle(notFound.title)
}

// NewMethodNotAllowed creates a new Problem instance that represents an HTTP 405 Method Not Allowed.
func NewMethodNotAllowed() *Problem {
	return New(methodNotAllowed.urn, methodNotAllowed.code).WithTitle(methodNotAllowed.title)
}

// NewConflict creates a new Problem instance that represents an HTTP 409 Conflict.
func NewConflict() *Problem {
	return New(conflict.urn, conflict.code).WithTitle(conflict.title)
}

// NewTooManyRequests creates a new Problem instance that represents an HTTP 429 Too Many Requests.
func NewTooManyRequests() *Problem {
	return New(tooManyRequests.urn, tooManyRequests.code).WithTitle(tooManyRequests.title)
}

// NewInternalServerError creates a new Problem instance that represents an HTTP 500 Internal Server Error.
func NewInternalServerError() *Problem {
	return New(internalServerError.urn, internalServerError.code).WithTitle(internalServerError.title)
}

// NewBadGateway creates a new Problem instance that represents an HTTP 502 Bad Gateway.
func NewBadGateway() *Problem {
	return New(badGateway.urn, badGateway.code).WithTitle(badGateway.title)
}

// NewServiceUnavailable creates a new Problem instance that represents an HTTP 503 Service Unavailable.
func NewServiceUnavailable() *Problem {
	return New(serviceUnavailable.urn, serviceUnavailable.code).WithTitle(serviceUnavailable.title)
}

// NewGatewayTimeout creates a new Problem instance that represents an HTTP 504 Gateway Timeout.
func NewGatewayTimeout() *Problem {
	return New(gatewayTimeout.urn, gatewayTimeout.code).WithTitle(gatewayTimeout.title)
}
