package problemdetails

import "net/http"

// ProblemType represents a problem type as defined in RFC 9457. This struct encapsulates the details of different types of HTTP problems, such as their type URIs, status codes, and titles. This struct would help in organizing and reusing these details across the codebase.
type ProblemType struct {
	Code  int
	Type  string
	Title string
}

var (
	Continue = ProblemType{
		Code:  http.StatusContinue,
		Type:  "urn:problems:continue",
		Title: "The server has received the request headers and the client should proceed to send the request body.",
	}
	SwitchingProtocols = ProblemType{
		Code:  http.StatusSwitchingProtocols,
		Type:  "urn:problems:switching-protocols",
		Title: "The server is switching protocols according to the Upgrade header.",
	}
	OK = ProblemType{
		Code:  http.StatusOK,
		Type:  "urn:problems:ok",
		Title: "The request has succeeded.",
	}
	Created = ProblemType{
		Code:  http.StatusCreated,
		Type:  "urn:problems:created",
		Title: "The request has been fulfilled and has resulted in one or more new resources being created.",
	}
	Accepted = ProblemType{
		Code:  http.StatusAccepted,
		Type:  "urn:problems:accepted",
		Title: "The request has been accepted for processing, but the processing has not been completed.",
	}
	NonAuthoritativeInfo = ProblemType{
		Code:  http.StatusNonAuthoritativeInfo,
		Type:  "urn:problems:non-authoritative-info",
		Title: "The request was successful but the enclosed payload has been modified from that of the origin server's 200 (OK) response by a transforming proxy.",
	}
	NoContent = ProblemType{
		Code:  http.StatusNoContent,
		Type:  "urn:problems:no-content",
		Title: "The server has successfully fulfilled the request and that there is no additional content to send in the response payload body.",
	}
	ResetContent = ProblemType{
		Code:  http.StatusResetContent,
		Type:  "urn:problems:reset-content",
		Title: "The server has fulfilled the request and desires that the user agent reset the 'document view' which caused the request to be sent.",
	}
	PartialContent = ProblemType{
		Code:  http.StatusPartialContent,
		Type:  "urn:problems:partial-content",
		Title: "The server is successfully fulfilling a range request for the target resource by transferring one or more parts of the selected representation.",
	}
	MultipleChoices = ProblemType{
		Code:  http.StatusMultipleChoices,
		Type:  "urn:problems:multiple-choices",
		Title: "The request has more than one possible response.",
	}
	MovedPermanently = ProblemType{
		Code:  http.StatusMovedPermanently,
		Type:  "urn:problems:moved-permanently",
		Title: "The requested resource has been assigned a new permanent URI.",
	}
	Found = ProblemType{
		Code:  http.StatusFound,
		Type:  "urn:problems:found",
		Title: "The requested resource resides temporarily under a different URI.",
	}
	SeeOther = ProblemType{
		Code:  http.StatusSeeOther,
		Type:  "urn:problems:see-other",
		Title: "The response to the request can be found under another URI using a GET method.",
	}
	NotModified = ProblemType{
		Code:  http.StatusNotModified,
		Type:  "urn:problems:not-modified",
		Title: "The resource has not been modified since the version specified by the request headers.",
	}
	UseProxy = ProblemType{
		Code:  http.StatusUseProxy,
		Type:  "urn:problems:use-proxy",
		Title: "The requested resource is available only through a proxy.",
	}
	TemporaryRedirect = ProblemType{
		Code:  http.StatusTemporaryRedirect,
		Type:  "urn:problems:temporary-redirect",
		Title: "The request should be repeated with another URI, but future requests can still use the original URI.",
	}
	PermanentRedirect = ProblemType{
		Code:  http.StatusPermanentRedirect,
		Type:  "urn:problems:permanent-redirect",
		Title: "The request and all future requests should be repeated using another URI.",
	}
	BadRequest = ProblemType{
		Code:  http.StatusBadRequest,
		Type:  "urn:problems:bad-request",
		Title: "The provided request could not be processed because it is invalid.",
	}
	Unauthorized = ProblemType{
		Code:  http.StatusUnauthorized,
		Type:  "urn:problems:unauthorized",
		Title: "The request requires user authentication.",
	}
	Forbidden = ProblemType{
		Code:  http.StatusForbidden,
		Type:  "urn:problems:forbidden",
		Title: "The server understood the request, but refuses to authorize it.",
	}
	NotFound = ProblemType{
		Code:  http.StatusNotFound,
		Type:  "urn:problems:not-found",
		Title: "The specified resource could not be found.",
	}
	MethodNotAllowed = ProblemType{
		Code:  http.StatusMethodNotAllowed,
		Type:  "urn:problems:method-not-allowed",
		Title: "The method specified in the request is not allowed for the resource.",
	}
	NotAcceptable = ProblemType{
		Code:  http.StatusNotAcceptable,
		Type:  "urn:problems:not-acceptable",
		Title: "The resource is not capable of generating content acceptable to the Accept headers.",
	}
	Conflict = ProblemType{
		Code:  http.StatusConflict,
		Type:  "urn:problems:conflict",
		Title: "The request could not be completed due to a conflict with the current state of the resource.",
	}
	Gone = ProblemType{
		Code:  http.StatusGone,
		Type:  "urn:problems:gone",
		Title: "The resource requested is no longer available and will not be available again.",
	}
	UnprocessableEntity = ProblemType{
		Code:  http.StatusUnprocessableEntity,
		Type:  "urn:problems:unprocessable-entity",
		Title: "The server understands the content type of the request entity, but was unable to process the contained instructions.",
	}
	InternalServerError = ProblemType{
		Code:  http.StatusInternalServerError,
		Type:  "urn:problems:internal-server-error",
		Title: "The server has encountered a situation it doesn't know how to handle.",
	}
	NotImplemented = ProblemType{
		Code:  http.StatusNotImplemented,
		Type:  "urn:problems:not-implemented",
		Title: "The request method is not supported by the server and cannot be handled.",
	}
	BadGateway = ProblemType{
		Code:  http.StatusBadGateway,
		Type:  "urn:problems:bad-gateway",
		Title: "The server, while acting as a gateway or proxy, received an invalid response from the upstream server.",
	}
	ServiceUnavailable = ProblemType{
		Code:  http.StatusServiceUnavailable,
		Type:  "urn:problems:service-unavailable",
		Title: "The server is not ready to handle the request, often due to maintenance or overload.",
	}
	GatewayTimeout = ProblemType{
		Code:  http.StatusGatewayTimeout,
		Type:  "urn:problems:gateway-timeout",
		Title: "The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.",
	}
	HTTPVersionNotSupported = ProblemType{
		Code:  http.StatusHTTPVersionNotSupported,
		Type:  "urn:problems:http-version-not-supported",
		Title: "The server does not support the HTTP protocol version used in the request.",
	}
)
