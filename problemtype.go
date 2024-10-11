package problemdetails

import "net/http"

// ProblemType represents a problem type as defined in RFC 9457. This struct encapsulates the details of different types of HTTP problems, such as their type URIs, status codes, and titles. This struct would help in organizing and reusing these details across the codebase.
type ProblemType struct {
	code  int
	urn   string
	title string
}

var (

	// BadRequest (HTTP 400) indicates that the provided request could not be processed because it is invalid.
	BadRequest = ProblemType{
		code:  http.StatusBadRequest,
		urn:   "urn:problems:bad-request",
		title: "The provided request could not be processed because it is invalid.",
	}

	// Unauthorized (HTTP 401) indicates that the request requires user authentication.
	Unauthorized = ProblemType{
		code:  http.StatusUnauthorized,
		urn:   "urn:problems:unauthorized",
		title: "The request requires user authentication.",
	}

	// Forbidden (HTTP 403) indicates that the server understood the request, but refuses to authorize it.
	Forbidden = ProblemType{
		code:  http.StatusForbidden,
		urn:   "urn:problems:forbidden",
		title: "The server understood the request, but refuses to authorize it.",
	}

	// NotFound (HTTP 404) indicates that the specified resource could not be found.
	NotFound = ProblemType{
		code:  http.StatusNotFound,
		urn:   "urn:problems:not-found",
		title: "The specified resource could not be found.",
	}

	// Conflict (HTTP 409) indicates that the request could not be completed due to a conflict with the current state of the resource.
	Conflict = ProblemType{
		code:  http.StatusConflict,
		urn:   "urn:problems:conflict",
		title: "The request could not be completed due to a conflict with the current state of the resource.",
	}

	// TooManyRequests (HTTP 429) indicates that the user has sent too many requests in a given amount of time.
	TooManyRequests = ProblemType{
		code:  http.StatusTooManyRequests,
		urn:   "urn:problems:too-many-requests",
		title: "The user has sent too many requests in a given amount of time.",
	}

	// InternalServerError (HTTP 500) indicates that the server has encountered a situation it doesn't know how to handle.
	InternalServerError = ProblemType{
		code:  http.StatusInternalServerError,
		urn:   "urn:problems:internal-server-error",
		title: "The server has encountered a situation it doesn't know how to handle.",
	}

	// NotImplemented (HTTP 501) indicates that the request method is not supported by the server and cannot be handled.
	NotImplemented = ProblemType{
		code:  http.StatusNotImplemented,
		urn:   "urn:problems:not-implemented",
		title: "The request method is not supported by the server and cannot be handled.",
	}

	// BadGateway (HTTP 502) indicates that the server, while acting as a gateway or proxy, received an invalid response from the upstream server.
	BadGateway = ProblemType{
		code:  http.StatusBadGateway,
		urn:   "urn:problems:bad-gateway",
		title: "The server, while acting as a gateway or proxy, received an invalid response from the upstream server.",
	}

	// ServiceUnavailable (HTTP 503) indicates that the server is not ready to handle the request, often due to maintenance or overload.
	ServiceUnavailable = ProblemType{
		code:  http.StatusServiceUnavailable,
		urn:   "urn:problems:service-unavailable",
		title: "The server is not ready to handle the request, often due to maintenance or overload.",
	}

	// GatewayTimeout (HTTP 504) indicates that the server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.
	GatewayTimeout = ProblemType{
		code:  http.StatusGatewayTimeout,
		urn:   "urn:problems:gateway-timeout",
		title: "The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.",
	}
)
