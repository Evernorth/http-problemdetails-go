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

// NewEarlyHints creates a new ProblemDetails instance that represents an HTTP 103 Early Hints.
func NewEarlyHints() *ProblemDetails {
	return New(EarlyHints.Type,
		EarlyHints.Code).WithTitle(EarlyHints.Title)
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

// NewNonAuthoritativeInfo creates a new ProblemDetails instance that represents an HTTP 203 Non-Authoritative Information.
func NewNonAuthoritativeInfo() *ProblemDetails {
	return New(NonAuthoritativeInfo.Type,
		NonAuthoritativeInfo.Code).WithTitle(NonAuthoritativeInfo.Title)
}

// NewNoContent creates a new ProblemDetails instance that represents an HTTP 204 No Content.
func NewNoContent() *ProblemDetails {
	return New(NoContent.Type,
		NoContent.Code).WithTitle(NoContent.Title)
}

// NewResetContent creates a new ProblemDetails instance that represents an HTTP 205 Reset Content.
func NewResetContent() *ProblemDetails {
	return New(ResetContent.Type,
		ResetContent.Code).WithTitle(ResetContent.Title)
}

// NewPartialContent creates a new ProblemDetails instance that represents an HTTP 206 Partial Content.
func NewPartialContent() *ProblemDetails {
	return New(PartialContent.Type,
		PartialContent.Code).WithTitle(PartialContent.Title)
}

// NewIMUsed creates a new ProblemDetails instance that represents an HTTP 226 IM Used.
func NewIMUsed() *ProblemDetails {
	return New(IMUsed.Type,
		IMUsed.Code).WithTitle(IMUsed.Title)
}

// NewMultipleChoices creates a new ProblemDetails instance that represents an HTTP 300 Multiple Choices.
func NewMultipleChoices() *ProblemDetails {
	return New(MultipleChoices.Type,
		MultipleChoices.Code).WithTitle(MultipleChoices.Title)
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

// NewSeeOther creates a new ProblemDetails instance that represents an HTTP 303 See Other.
func NewSeeOther() *ProblemDetails {
	return New(SeeOther.Type,
		SeeOther.Code).WithTitle(SeeOther.Title)
}

// NewNotModified creates a new ProblemDetails instance that represents an HTTP 304 Not Modified.
func NewNotModified() *ProblemDetails {
	return New(NotModified.Type,
		NotModified.Code).WithTitle(NotModified.Title)
}

// NewUseProxy creates a new ProblemDetails instance that represents an HTTP 305 Use Proxy.
func NewUseProxy() *ProblemDetails {
	return New(UseProxy.Type,
		UseProxy.Code).WithTitle(UseProxy.Title)
}

// NewTemporaryRedirect creates a new ProblemDetails instance that represents an HTTP 307 Temporary Redirect.
func NewTemporaryRedirect() *ProblemDetails {
	return New(TemporaryRedirect.Type,
		TemporaryRedirect.Code).WithTitle(TemporaryRedirect.Title)
}

// NewPermanentRedirect creates a new ProblemDetails instance that represents an HTTP 308 Permanent Redirect.
func NewPermanentRedirect() *ProblemDetails {
	return New(PermanentRedirect.Type,
		PermanentRedirect.Code).WithTitle(PermanentRedirect.Title)
}

// NewBadRequest creates a new ProblemDetails instance that represents an HTTP 400 Bad Request.
func NewBadRequest() *ProblemDetails {
	return New(BadRequest.Type,
		BadRequest.Code).WithTitle(BadRequest.Title)
}

// NewUnauthorized creates a new ProblemDetails instance that represents an HTTP 401 Unauthorized.
func NewUnauthorized() *ProblemDetails {
	return New(Unauthorized.Type,
		Unauthorized.Code).WithTitle(Unauthorized.Title)
}

// NewPaymentRequired creates a new ProblemDetails instance that represents an HTTP 402 Payment Required.
func NewPaymentRequired() *ProblemDetails {
	return New(PaymentRequired.Type,
		PaymentRequired.Code).WithTitle(PaymentRequired.Title)
}

// NewForbidden creates a new ProblemDetails instance that represents an HTTP 403 Forbidden.
func NewForbidden() *ProblemDetails {
	return New(Forbidden.Type,
		Forbidden.Code).WithTitle(Forbidden.Title)
}

// NewNotFound creates a new ProblemDetails instance that represents an HTTP 404 Not Found.
func NewNotFound() *ProblemDetails {
	return New(NotFound.Type,
		NotFound.Code).WithTitle(NotFound.Title)
}

// NewMethodNotAllowed creates a new ProblemDetails instance that represents an HTTP 405 Method Not Allowed.
func NewMethodNotAllowed() *ProblemDetails {
	return New(MethodNotAllowed.Type,
		MethodNotAllowed.Code).WithTitle(MethodNotAllowed.Title)
}

// NewNotAcceptable creates a new ProblemDetails instance that represents an HTTP 406 Not Acceptable.
func NewNotAcceptable() *ProblemDetails {
	return New(NotAcceptable.Type,
		NotAcceptable.Code).WithTitle(NotAcceptable.Title)
}

// NewProxyAuthenticationRequired creates a new ProblemDetails instance that represents an HTTP 407 Proxy Authentication Required.
func NewProxyAuthenticationRequired() *ProblemDetails {
	return New(ProxyAuthenticationRequired.Type,
		ProxyAuthenticationRequired.Code).WithTitle(ProxyAuthenticationRequired.Title)
}

// NewRequestTimeout creates a new ProblemDetails instance that represents an HTTP 408 Request Timeout.
func NewRequestTimeout() *ProblemDetails {
	return New(RequestTimeout.Type,
		RequestTimeout.Code).WithTitle(RequestTimeout.Title)
}

// NewConflict creates a new ProblemDetails instance that represents an HTTP 409 Conflict.
func NewConflict() *ProblemDetails {
	return New(Conflict.Type,
		Conflict.Code).WithTitle(Conflict.Title)
}

// NewGone creates a new ProblemDetails instance that represents an HTTP 410 Gone.
func NewGone() *ProblemDetails {
	return New(Gone.Type,
		Gone.Code).WithTitle(Gone.Title)
}

// NewLengthRequired creates a new ProblemDetails instance that represents an HTTP 411 Length Required.
func NewLengthRequired() *ProblemDetails {
	return New(LengthRequired.Type,
		LengthRequired.Code).WithTitle(LengthRequired.Title)
}

// NewPreconditionFailed creates a new ProblemDetails instance that represents an HTTP 412 Precondition Failed.
func NewPreconditionFailed() *ProblemDetails {
	return New(PreconditionFailed.Type,
		PreconditionFailed.Code).WithTitle(PreconditionFailed.Title)
}

// NewPayloadTooLarge creates a new ProblemDetails instance that represents an HTTP 413 Payload Too Large.
func NewPayloadTooLarge() *ProblemDetails {
	return New(PayloadTooLarge.Type,
		PayloadTooLarge.Code).WithTitle(PayloadTooLarge.Title)
}

// NewURITooLong creates a new ProblemDetails instance that represents an HTTP 414 URI Too Long.
func NewURITooLong() *ProblemDetails {
	return New(URITooLong.Type,
		URITooLong.Code).WithTitle(URITooLong.Title)
}

// NewUnsupportedMediaType creates a new ProblemDetails instance that represents an HTTP 415 Unsupported Media Type.
func NewUnsupportedMediaType() *ProblemDetails {
	return New(UnsupportedMediaType.Type,
		UnsupportedMediaType.Code).WithTitle(UnsupportedMediaType.Title)
}

// NewRangeNotSatisfiable creates a new ProblemDetails instance that represents an HTTP 416 Range Not Satisfiable.
func NewRangeNotSatisfiable() *ProblemDetails {
	return New(RangeNotSatisfiable.Type,
		RangeNotSatisfiable.Code).WithTitle(RangeNotSatisfiable.Title)
}

// NewExpectationFailed creates a new ProblemDetails instance that represents an HTTP 417 Expectation Failed.
func NewExpectationFailed() *ProblemDetails {
	return New(ExpectationFailed.Type,
		ExpectationFailed.Code).WithTitle(ExpectationFailed.Title)
}

// NewMisdirectedRequest creates a new ProblemDetails instance that represents an HTTP 421 Misdirected Request.
func NewMisdirectedRequest() *ProblemDetails {
	return New(MisdirectedRequest.Type,
		MisdirectedRequest.Code).WithTitle(MisdirectedRequest.Title)
}

// NewUpgradeRequired creates a new ProblemDetails instance that represents an HTTP 426 Upgrade Required.
func NewUpgradeRequired() *ProblemDetails {
	return New(UpgradeRequired.Type,
		UpgradeRequired.Code).WithTitle(UpgradeRequired.Title)
}

// NewPreconditionRequired creates a new ProblemDetails instance that represents an HTTP 428 Precondition Required.
func NewPreconditionRequired() *ProblemDetails {
	return New(PreconditionRequired.Type,
		PreconditionRequired.Code).WithTitle(PreconditionRequired.Title)
}

// NewTooManyRequests creates a new ProblemDetails instance that represents an HTTP 429 Too Many Requests.
func NewTooManyRequests() *ProblemDetails {
	return New(TooManyRequests.Type,
		TooManyRequests.Code).WithTitle(TooManyRequests.Title)
}

// NewRequestHeaderFieldsTooLarge creates a new ProblemDetails instance that represents an HTTP 431 Request Header Fields Too Large.
func NewRequestHeaderFieldsTooLarge() *ProblemDetails {
	return New(RequestHeaderFieldsTooLarge.Type,
		RequestHeaderFieldsTooLarge.Code).WithTitle(RequestHeaderFieldsTooLarge.Title)
}

// NewUnavailableForLegalReasons creates a new ProblemDetails instance that represents an HTTP 451 Unavailable For Legal Reasons.
func NewUnavailableForLegalReasons() *ProblemDetails {
	return New(UnavailableForLegalReasons.Type,
		UnavailableForLegalReasons.Code).WithTitle(UnavailableForLegalReasons.Title)
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

// NewGatewayTimeout creates a new ProblemDetails instance that represents an HTTP 504 Gateway Timeout.
func NewGatewayTimeout() *ProblemDetails {
	return New(GatewayTimeout.Type,
		GatewayTimeout.Code).WithTitle(GatewayTimeout.Title)
}

// NewHTTPVersionNotSupported creates a new ProblemDetails instance that represents an HTTP 505 HTTP Version Not Supported.
func NewHTTPVersionNotSupported() *ProblemDetails {
	return New(HTTPVersionNotSupported.Type,
		HTTPVersionNotSupported.Code).WithTitle(HTTPVersionNotSupported.Title)
}

// NewVariantAlsoNegotiates creates a new ProblemDetails instance that represents an HTTP 506 Variant Also Negotiates.
func NewVariantAlsoNegotiates() *ProblemDetails {
	return New(VariantAlsoNegotiates.Type,
		VariantAlsoNegotiates.Code).WithTitle(VariantAlsoNegotiates.Title)
}

// NewNotExtended creates a new ProblemDetails instance that represents an HTTP 510 Not Extended.
func NewNotExtended() *ProblemDetails {
	return New(NotExtended.Type,
		NotExtended.Code).WithTitle(NotExtended.Title)
}

// NewNetworkAuthenticationRequired creates a new ProblemDetails instance that represents an HTTP 511 Network Authentication Required.
func NewNetworkAuthenticationRequired() *ProblemDetails {
	return New(NetworkAuthenticationRequired.Type,
		NetworkAuthenticationRequired.Code).WithTitle(NetworkAuthenticationRequired.Title)
}

// NewTeapot creates a new ProblemDetails instance that represents an HTTP 418 I'm a teapot.
func NewTeapot() *ProblemDetails {
	return New(Teapot.Type,
		Teapot.Code).WithTitle(Teapot.Title)
}

// NewTooEarly creates a new ProblemDetails instance that represents an HTTP 425 Too Early.
func NewTooEarly() *ProblemDetails {
	return New(TooEarly.Type,
		TooEarly.Code).WithTitle(TooEarly.Title)
}
