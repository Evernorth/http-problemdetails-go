package problemdetails

import "net/http"

// problemType represents a problem type as defined in RFC 9457. This struct encapsulates the details of different
// types of HTTP problems, such as their type URIs, status codes, and titles. This struct would help in organizing
// and reusing these details across the codebase.
type problemType struct {
	code  int
	urn   string
	title string
}

var (

	// badRequest (HTTP 400)
	badRequest = problemType{
		code:  http.StatusBadRequest,
		urn:   "urn:problems:bad-request",
		title: "Request could not be processed because it is invalid.",
	}

	// unauthorized (HTTP 401)
	unauthorized = problemType{
		code:  http.StatusUnauthorized,
		urn:   "urn:problems:unauthorized",
		title: "Authentication required.",
	}

	// forbidden (HTTP 403)
	forbidden = problemType{
		code:  http.StatusForbidden,
		urn:   "urn:problems:forbidden",
		title: "User is not authorized to perform the requested operation.",
	}

	// notFound (HTTP 404)
	notFound = problemType{
		code:  http.StatusNotFound,
		urn:   "urn:problems:not-found",
		title: "The specified resource could not be found.",
	}

	// notFound (HTTP 405)
	methodNotAllowed = problemType{
		code:  http.StatusMethodNotAllowed,
		urn:   "urn:problems:method-not-allowed",
		title: "The specified method is not allowed.",
	}

	// conflict (HTTP 409)
	conflict = problemType{
		code:  http.StatusConflict,
		urn:   "urn:problems:conflict",
		title: "Request could not be completed due to a conflict with the current state of the resource.",
	}

	// tooManyRequests (HTTP 429)
	tooManyRequests = problemType{
		code:  http.StatusTooManyRequests,
		urn:   "urn:problems:too-many-requests",
		title: "User has sent too many requests.",
	}

	// internalServerError (HTTP 500)
	internalServerError = problemType{
		code:  http.StatusInternalServerError,
		urn:   "urn:problems:internal-server-error",
		title: "An unexpected error occurred.",
	}

	// badGateway (HTTP 502)
	badGateway = problemType{
		code:  http.StatusBadGateway,
		urn:   "urn:problems:bad-gateway",
		title: "Invalid response from upstream server.",
	}

	// serviceUnavailable (HTTP 503)
	serviceUnavailable = problemType{
		code:  http.StatusServiceUnavailable,
		urn:   "urn:problems:service-unavailable",
		title: "Service is temporarily unavailable.",
	}

	// gatewayTimeout (HTTP 504)
	gatewayTimeout = problemType{
		code:  http.StatusGatewayTimeout,
		urn:   "urn:problems:gateway-timeout",
		title: "Timeout invoking upstream server.",
	}
)
