package problemdetails

import "net/http"

// ProblemType represents a problem type as defined in RFC 9457. This struct encapsulates the details of different types of HTTP problems, such as their type URIs, status codes, and titles. This struct would help in organizing and reusing these details across the codebase.
type ProblemType struct {
	Code   int
	Type   string
	Title  string
	Detail string
}

var (
	// Continue (HTTP 100) indicates that the initial part of a request has been received and has not yet been rejected by the server.
	Continue = ProblemType{
		Code:   http.StatusContinue,
		Type:   "urn:problems:continue",
		Title:  "Continue",
		Detail: "The server has received the request headers and the client should proceed to send the request body.",
	}

	// SwitchingProtocols (HTTP 101) indicates that the server is willing to switch protocols as requested by the client.
	SwitchingProtocols = ProblemType{
		Code:   http.StatusSwitchingProtocols,
		Type:   "urn:problems:switching-protocols",
		Title:  "Switching Protocols",
		Detail: "The server is switching protocols according to the Upgrade header.",
	}

	// EarlyHints (HTTP 103) indicates that the server is likely to send a final response with the header fields included in the informational response.
	EarlyHints = ProblemType{
		Code:   http.StatusEarlyHints,
		Type:   "urn:problems:early-hints",
		Title:  "Early Hints",
		Detail: "The server is likely to send a final response with the header fields included in the informational response.",
	}

	// OK (HTTP 200) indicates that the request has succeeded.
	OK = ProblemType{
		Code:   http.StatusOK,
		Type:   "urn:problems:ok",
		Title:  "OK",
		Detail: "The request has succeeded.",
	}

	// Created (HTTP 201) indicates that the request has been fulfilled and has resulted in one or more new resources being created.
	Created = ProblemType{
		Code:   http.StatusCreated,
		Type:   "urn:problems:created",
		Title:  "Created",
		Detail: "The request has been fulfilled and has resulted in one or more new resources being created.",
	}

	// Accepted (HTTP 202) indicates that the request has been accepted for processing, but the processing has not been completed.
	Accepted = ProblemType{
		Code:   http.StatusAccepted,
		Type:   "urn:problems:accepted",
		Title:  "Accepted",
		Detail: "The request has been accepted for processing, but the processing has not been completed.",
	}

	// NonAuthoritativeInfo (HTTP 203) indicates that the request was successful but the enclosed payload has been modified from that of the origin server's 200 (OK) response by a transforming proxy.
	NonAuthoritativeInfo = ProblemType{
		Code:   http.StatusNonAuthoritativeInfo,
		Type:   "urn:problems:non-authoritative-info",
		Title:  "Non-Authoritative Information",
		Detail: "The request was successful but the enclosed payload has been modified from that of the origin server's 200 (OK) response by a transforming proxy.",
	}

	// NoContent (HTTP 204) indicates that the server has successfully fulfilled the request and that there is no additional content to send in the response payload body.
	NoContent = ProblemType{
		Code:   http.StatusNoContent,
		Type:   "urn:problems:no-content",
		Title:  "No Content",
		Detail: "The server has successfully fulfilled the request and that there is no additional content to send in the response payload body.",
	}

	// ResetContent (HTTP 205) indicates that the server has fulfilled the request and desires that the user agent reset the 'document view' which caused the request to be sent.
	ResetContent = ProblemType{
		Code:   http.StatusResetContent,
		Type:   "urn:problems:reset-content",
		Title:  "Reset Content",
		Detail: "The server has fulfilled the request and desires that the user agent reset the 'document view' which caused the request to be sent.",
	}

	// PartialContent (HTTP 206) indicates that the server is successfully fulfilling a range request for the target resource by transferring one or more parts of the selected representation.
	PartialContent = ProblemType{
		Code:   http.StatusPartialContent,
		Type:   "urn:problems:partial-content",
		Title:  "Partial Content",
		Detail: "The server is successfully fulfilling a range request for the target resource by transferring one or more parts of the selected representation.",
	}

	// IMUsed (HTTP 226) indicates that the server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.
	IMUsed = ProblemType{
		Code:   http.StatusIMUsed,
		Type:   "urn:problems:im-used",
		Title:  "IM Used",
		Detail: "The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.",
	}

	// MultipleChoices (HTTP 300) indicates that the request has more than one possible response.
	MultipleChoices = ProblemType{
		Code:   http.StatusMultipleChoices,
		Type:   "urn:problems:multiple-choices",
		Title:  "Multiple Choices",
		Detail: "The request has more than one possible response.",
	}

	// MovedPermanently (HTTP 301) indicates that the requested resource has been assigned a new permanent URI.
	MovedPermanently = ProblemType{
		Code:   http.StatusMovedPermanently,
		Type:   "urn:problems:moved-permanently",
		Title:  "Moved Permanently",
		Detail: "The requested resource has been assigned a new permanent URI.",
	}

	// Found (HTTP 302) indicates that the requested resource resides temporarily under a different URI.
	Found = ProblemType{
		Code:   http.StatusFound,
		Type:   "urn:problems:found",
		Title:  "Found",
		Detail: "The requested resource resides temporarily under a different URI.",
	}

	// SeeOther (HTTP 303) indicates that the response to the request can be found under another URI using a GET method.
	SeeOther = ProblemType{
		Code:   http.StatusSeeOther,
		Type:   "urn:problems:see-other",
		Title:  "See Other",
		Detail: "The response to the request can be found under another URI using a GET method.",
	}

	// NotModified (HTTP 304) indicates that the resource has not been modified since the version specified by the request headers.
	NotModified = ProblemType{
		Code:   http.StatusNotModified,
		Type:   "urn:problems:not-modified",
		Title:  "Not Modified",
		Detail: "The resource has not been modified since the version specified by the request headers.",
	}

	// UseProxy (HTTP 305) indicates that the requested resource is available only through a proxy.
	UseProxy = ProblemType{
		Code:   http.StatusUseProxy,
		Type:   "urn:problems:use-proxy",
		Title:  "Use Proxy",
		Detail: "The requested resource is available only through a proxy.",
	}

	// TemporaryRedirect (HTTP 307) indicates that the request should be repeated with another URI, but future requests can still use the original URI.
	TemporaryRedirect = ProblemType{
		Code:   http.StatusTemporaryRedirect,
		Type:   "urn:problems:temporary-redirect",
		Title:  "Temporary Redirect",
		Detail: "The request should be repeated with another URI, but future requests can still use the original URI.",
	}

	// PermanentRedirect (HTTP 308) indicates that the request and all future requests should be repeated using another URI.
	PermanentRedirect = ProblemType{
		Code:   http.StatusPermanentRedirect,
		Type:   "urn:problems:permanent-redirect",
		Title:  "Permanent Redirect",
		Detail: "The request and all future requests should be repeated using another URI.",
	}

	// BadRequest (HTTP 400) indicates that the provided request could not be processed because it is invalid.
	BadRequest = ProblemType{
		Code:   http.StatusBadRequest,
		Type:   "urn:problems:bad-request",
		Title:  "Bad Request",
		Detail: "The provided request could not be processed because it is invalid.",
	}

	// Unauthorized (HTTP 401) indicates that the request requires user authentication.
	Unauthorized = ProblemType{
		Code:   http.StatusUnauthorized,
		Type:   "urn:problems:unauthorized",
		Title:  "Unauthorized",
		Detail: "The request requires user authentication.",
	}

	// PaymentRequired (HTTP 402) indicates that the request requires payment before it can be processed.
	PaymentRequired = ProblemType{
		Code:   http.StatusPaymentRequired,
		Type:   "urn:problems:payment-required",
		Title:  "Payment Required",
		Detail: "The request requires payment before it can be processed.",
	}

	// Forbidden (HTTP 403) indicates that the server understood the request, but refuses to authorize it.
	Forbidden = ProblemType{
		Code:   http.StatusForbidden,
		Type:   "urn:problems:forbidden",
		Title:  "Forbidden",
		Detail: "The server understood the request, but refuses to authorize it.",
	}

	// NotFound (HTTP 404) indicates that the specified resource could not be found.
	NotFound = ProblemType{
		Code:   http.StatusNotFound,
		Type:   "urn:problems:not-found",
		Title:  "Not Found",
		Detail: "The specified resource could not be found.",
	}

	// MethodNotAllowed (HTTP 405) indicates that the method specified in the request is not allowed for the resource.
	MethodNotAllowed = ProblemType{
		Code:   http.StatusMethodNotAllowed,
		Type:   "urn:problems:method-not-allowed",
		Title:  "Method Not Allowed",
		Detail: "The method specified in the request is not allowed for the resource.",
	}

	// NotAcceptable (HTTP 406) indicates that the resource is not capable of generating content acceptable to the Accept headers.
	NotAcceptable = ProblemType{
		Code:   http.StatusNotAcceptable,
		Type:   "urn:problems:not-acceptable",
		Title:  "Not Acceptable",
		Detail: "The resource is not capable of generating content acceptable to the Accept headers.",
	}

	// ProxyAuthenticationRequired (HTTP 407) indicates that the client must first authenticate itself with the proxy.
	ProxyAuthenticationRequired = ProblemType{
		Code:   http.StatusProxyAuthRequired,
		Type:   "urn:problems:proxy-authentication-required",
		Title:  "Proxy Authentication Required",
		Detail: "The client must first authenticate itself with the proxy.",
	}

	// RequestTimeout (HTTP 408) indicates that the server timed out waiting for the request.
	RequestTimeout = ProblemType{
		Code:   http.StatusRequestTimeout,
		Type:   "urn:problems:request-timeout",
		Title:  "Request Timeout",
		Detail: "The server timed out waiting for the request.",
	}

	// Conflict (HTTP 409) indicates that the request could not be completed due to a conflict with the current state of the resource.
	Conflict = ProblemType{
		Code:   http.StatusConflict,
		Type:   "urn:problems:conflict",
		Title:  "Conflict",
		Detail: "The request could not be completed due to a conflict with the current state of the resource.",
	}

	// Gone (HTTP 410) indicates that the resource requested is no longer available and will not be available again.
	Gone = ProblemType{
		Code:   http.StatusGone,
		Type:   "urn:problems:gone",
		Title:  "Gone",
		Detail: "The resource requested is no longer available and will not be available again.",
	}

	// LengthRequired (HTTP 411) indicates that the request did not specify the length of its content, which is required by the requested resource.
	LengthRequired = ProblemType{
		Code:   http.StatusLengthRequired,
		Type:   "urn:problems:length-required",
		Title:  "Length Required",
		Detail: "The request did not specify the length of its content, which is required by the requested resource.",
	}

	// PreconditionFailed (HTTP 412) indicates that the server does not meet one of the preconditions that the requester put on the request.
	PreconditionFailed = ProblemType{
		Code:   http.StatusPreconditionFailed,
		Type:   "urn:problems:precondition-failed",
		Title:  "Precondition Failed",
		Detail: "The server does not meet one of the preconditions that the requester put on the request.",
	}

	// PayloadTooLarge (HTTP 413) indicates that the request is larger than the server is willing or able to process.
	PayloadTooLarge = ProblemType{
		Code:   http.StatusRequestEntityTooLarge,
		Type:   "urn:problems:payload-too-large",
		Title:  "Payload Too Large",
		Detail: "The request is larger than the server is willing or able to process.",
	}

	// URITooLong (HTTP 414) indicates that the URI provided was too long for the server to process.
	URITooLong = ProblemType{
		Code:   http.StatusRequestURITooLong,
		Type:   "urn:problems:uri-too-long",
		Title:  "URI Too Long",
		Detail: "The URI provided was too long for the server to process.",
	}

	// UnsupportedMediaType (HTTP 415) indicates that the request entity has a media type which the server or resource does not support.
	UnsupportedMediaType = ProblemType{
		Code:   http.StatusUnsupportedMediaType,
		Type:   "urn:problems:unsupported-media-type",
		Title:  "Unsupported Media Type",
		Detail: "The request entity has a media type which the server or resource does not support.",
	}

	// RangeNotSatisfiable (HTTP 416) indicates that the client has asked for a portion of the file, but the server cannot supply that portion.
	RangeNotSatisfiable = ProblemType{
		Code:   http.StatusRequestedRangeNotSatisfiable,
		Type:   "urn:problems:range-not-satisfiable",
		Title:  "Range Not Satisfiable",
		Detail: "The client has asked for a portion of the file, but the server cannot supply that portion.",
	}

	// ExpectationFailed (HTTP 417) indicates that the server cannot meet the requirements of the Expect request-header field.
	ExpectationFailed = ProblemType{
		Code:   http.StatusExpectationFailed,
		Type:   "urn:problems:expectation-failed",
		Title:  "Expectation Failed",
		Detail: "The server cannot meet the requirements of the Expect request-header field.",
	}

	// Teapot (HTTP 418) indicates that the server is a teapot and cannot brew coffee.
	Teapot = ProblemType{
		Code:   http.StatusTeapot,
		Type:   "urn:problems:teapot",
		Title:  "I'm a teapot",
		Detail: "I'm a teapot.",
	}

	// MisdirectedRequest (HTTP 421) indicates that the request was directed at a server that is not able to produce a response.
	MisdirectedRequest = ProblemType{
		Code:   http.StatusMisdirectedRequest,
		Type:   "urn:problems:misdirected-request",
		Title:  "Misdirected Request",
		Detail: "The request was directed at a server that is not able to produce a response.",
	}

	// TooEarly (HTTP 425) indicates that the server is unwilling to risk processing a request that might be replayed.
	TooEarly = ProblemType{
		Code:   http.StatusTooEarly,
		Type:   "urn:problems:too-early",
		Title:  "Too Early",
		Detail: "The server is unwilling to risk processing a request that might be replayed.",
	}

	// UpgradeRequired (HTTP 426) indicates that the client should switch to a different protocol.
	UpgradeRequired = ProblemType{
		Code:   http.StatusUpgradeRequired,
		Type:   "urn:problems:upgrade-required",
		Title:  "Upgrade Required",
		Detail: "The client should switch to a different protocol.",
	}

	// PreconditionRequired (HTTP 428) indicates that the server requires the request to be conditional.
	PreconditionRequired = ProblemType{
		Code:   http.StatusPreconditionRequired,
		Type:   "urn:problems:precondition-required",
		Title:  "Precondition Required",
		Detail: "The server requires the request to be conditional.",
	}

	// TooManyRequests (HTTP 429) indicates that the user has sent too many requests in a given amount of time.
	TooManyRequests = ProblemType{
		Code:   http.StatusTooManyRequests,
		Type:   "urn:problems:too-many-requests",
		Title:  "Too Many Requests",
		Detail: "The user has sent too many requests in a given amount of time.",
	}

	// RequestHeaderFieldsTooLarge (HTTP 431) indicates that the server is unwilling to process the request because its header fields are too large.
	RequestHeaderFieldsTooLarge = ProblemType{
		Code:   http.StatusRequestHeaderFieldsTooLarge,
		Type:   "urn:problems:request-header-fields-too-large",
		Title:  "Request Header Fields Too Large",
		Detail: "The server is unwilling to process the request because its header fields are too large.",
	}

	// UnavailableForLegalReasons (HTTP 451) indicates that the resource is unavailable for legal reasons.
	UnavailableForLegalReasons = ProblemType{
		Code:   http.StatusUnavailableForLegalReasons,
		Type:   "urn:problems:unavailable-for-legal-reasons",
		Title:  "Unavailable For Legal Reasons",
		Detail: "The resource is unavailable for legal reasons.",
	}

	// InternalServerError (HTTP 500) indicates that the server has encountered a situation it doesn't know how to handle.
	InternalServerError = ProblemType{
		Code:   http.StatusInternalServerError,
		Type:   "urn:problems:internal-server-error",
		Title:  "Internal Server Error",
		Detail: "The server has encountered a situation it doesn't know how to handle.",
	}

	// NotImplemented (HTTP 501) indicates that the request method is not supported by the server and cannot be handled.
	NotImplemented = ProblemType{
		Code:   http.StatusNotImplemented,
		Type:   "urn:problems:not-implemented",
		Title:  "Not Implemented",
		Detail: "The request method is not supported by the server and cannot be handled.",
	}

	// BadGateway (HTTP 502) indicates that the server, while acting as a gateway or proxy, received an invalid response from the upstream server.
	BadGateway = ProblemType{
		Code:   http.StatusBadGateway,
		Type:   "urn:problems:bad-gateway",
		Title:  "Bad Gateway",
		Detail: "The server, while acting as a gateway or proxy, received an invalid response from the upstream server.",
	}

	// ServiceUnavailable (HTTP 503) indicates that the server is not ready to handle the request, often due to maintenance or overload.
	ServiceUnavailable = ProblemType{
		Code:   http.StatusServiceUnavailable,
		Type:   "urn:problems:service-unavailable",
		Title:  "Service Unavailable",
		Detail: "The server is not ready to handle the request, often due to maintenance or overload.",
	}

	// GatewayTimeout (HTTP 504) indicates that the server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.
	GatewayTimeout = ProblemType{
		Code:   http.StatusGatewayTimeout,
		Type:   "urn:problems:gateway-timeout",
		Title:  "Gateway Timeout",
		Detail: "The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.",
	}

	// HTTPVersionNotSupported (HTTP 505) indicates that the server does not support the HTTP protocol version used in the request.
	HTTPVersionNotSupported = ProblemType{
		Code:   http.StatusHTTPVersionNotSupported,
		Type:   "urn:problems:http-version-not-supported",
		Title:  "HTTP Version Not Supported",
		Detail: "The server does not support the HTTP protocol version used in the request.",
	}

	// VariantAlsoNegotiates (HTTP 506) indicates that the server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.
	VariantAlsoNegotiates = ProblemType{
		Code:   http.StatusVariantAlsoNegotiates,
		Type:   "urn:problems:variant-also-negotiates",
		Title:  "Variant Also Negotiates",
		Detail: "The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.",
	}

	// NotExtended (HTTP 510) indicates that further extensions to the request are required for the server to fulfill it.
	NotExtended = ProblemType{
		Code:   http.StatusNotExtended,
		Type:   "urn:problems:not-extended",
		Title:  "Not Extended",
		Detail: "Further extensions to the request are required for the server to fulfill it.",
	}

	// NetworkAuthenticationRequired (HTTP 511) indicates that the client needs to authenticate to gain network access.
	NetworkAuthenticationRequired = ProblemType{
		Code:   http.StatusNetworkAuthenticationRequired,
		Type:   "urn:problems:network-authentication-required",
		Title:  "Network Authentication Required",
		Detail: "The client needs to authenticate to gain network access.",
	}
)
