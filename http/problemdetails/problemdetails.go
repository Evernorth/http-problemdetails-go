package problemdetails

import (
	"github.com/google/uuid"
	"time"
)

const (
	uuidUrnPrefix = "urn:uuid:"
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

// NewContinue creates a new ProblemDetails instance that represents an HTTP 100 Continue.
func NewContinue() *ProblemDetails {
	return New(Continue.Type,
		Continue.Code).WithTitle(Continue.Title)
}

// NewSwitchingProtocols creates a new ProblemDetails instance that represents an HTTP 101 Switching Protocols.
func NewSwitchingProtocols() *ProblemDetails {
	return New(SwitchingProtocols.Type,
		SwitchingProtocols.Code).WithTitle(SwitchingProtocols.Title)
}

// NewOK creates a new ProblemDetails instance that represents an HTTP 200 OK.
func NewOK() *ProblemDetails {
	return New(OK.Type,
		OK.Code).WithTitle(OK.Title)
}

// NewCreated creates a new ProblemDetails instance that represents an HTTP 201 Created.
func NewCreated() *ProblemDetails {
	return New(Created.Type,
		Created.Code).WithTitle(Created.Title)
}

// NewAccepted creates a new ProblemDetails instance that represents an HTTP 202 Accepted.
func NewAccepted() *ProblemDetails {
	return New(Accepted.Type,
		Accepted.Code).WithTitle(Accepted.Title)
}

// NewNoContent creates a new ProblemDetails instance that represents an HTTP 204 No Content.
func NewNoContent() *ProblemDetails {
	return New(NoContent.Type,
		NoContent.Code).WithTitle(NoContent.Title)
}

// NewMovedPermanently creates a new ProblemDetails instance that represents an HTTP 301 Moved Permanently.
func NewMovedPermanently() *ProblemDetails {
	return New(MovedPermanently.Type,
		MovedPermanently.Code).WithTitle(MovedPermanently.Title)
}

// NewFound creates a new ProblemDetails instance that represents an HTTP 302 Found.
func NewFound() *ProblemDetails {
	return New(Found.Type,
		Found.Code).WithTitle(Found.Title)
}

// NewNotModified creates a new ProblemDetails instance that represents an HTTP 304 Not Modified.
func NewNotModified() *ProblemDetails {
	return New(NotModified.Type,
		NotModified.Code).WithTitle(NotModified.Title)
}

// NewBadRequest creates a new ProblemDetails instance that represents an HTTP 400 Bad Request.
func NewBadRequest() *ProblemDetails {
	return New(BadRequest.Type,
		BadRequest.Code).WithTitle(BadRequest.Title)
}

// NewNotFound creates a new ProblemDetails instance that represents an HTTP 404 Not Found.
func NewNotFound() *ProblemDetails {
	return New(NotFound.Type,
		NotFound.Code).WithTitle(NotFound.Title)
}

// NewUnauthorized creates a new ProblemDetails instance that represents an HTTP 401 Unauthorized.
func NewUnauthorized() *ProblemDetails {
	return New(Unauthorized.Type,
		Unauthorized.Code).WithTitle(Unauthorized.Title)
}

// NewForbidden creates a new ProblemDetails instance that represents an HTTP 403 Forbidden.
func NewForbidden() *ProblemDetails {
	return New(Forbidden.Type,
		Forbidden.Code).WithTitle(Forbidden.Title)
}

// NewMethodNotAllowed creates a new ProblemDetails instance that represents an HTTP 405 Method Not Allowed.
func NewMethodNotAllowed() *ProblemDetails {
	return New(MethodNotAllowed.Type,
		MethodNotAllowed.Code).WithTitle(MethodNotAllowed.Title)
}

// NewConflict creates a new ProblemDetails instance that represents an HTTP 409 Conflict.
func NewConflict() *ProblemDetails {
	return New(Conflict.Type,
		Conflict.Code).WithTitle(Conflict.Title)
}

// NewUnprocessableEntity creates a new ProblemDetails instance that represents an HTTP 422 Unprocessable Entity.
func NewUnprocessableEntity() *ProblemDetails {
	return New(UnprocessableEntity.Type,
		UnprocessableEntity.Code).WithTitle(UnprocessableEntity.Title)
}

// NewInternalServerError creates a new ProblemDetails instance that represents an HTTP 500 Internal Server Error.
func NewInternalServerError() *ProblemDetails {
	return New(InternalServerError.Type,
		InternalServerError.Code).WithTitle(InternalServerError.Title)
}

// NewNotImplemented creates a new ProblemDetails instance that represents an HTTP 501 Not Implemented.
func NewNotImplemented() *ProblemDetails {
	return New(NotImplemented.Type,
		NotImplemented.Code).WithTitle(NotImplemented.Title)
}

// NewBadGateway creates a new ProblemDetails instance that represents an HTTP 502 Bad Gateway.
func NewBadGateway() *ProblemDetails {
	return New(BadGateway.Type,
		BadGateway.Code).WithTitle(BadGateway.Title)
}

// NewServiceUnavailable creates a new ProblemDetails instance that represents an HTTP 503 Service Unavailable.
func NewServiceUnavailable() *ProblemDetails {
	return New(ServiceUnavailable.Type,
		ServiceUnavailable.Code).WithTitle(ServiceUnavailable.Title)
}
