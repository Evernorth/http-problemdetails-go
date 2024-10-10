package problemdetails

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestNewInternalServerError(t *testing.T) {
	problemDetails := NewInternalServerError()
	assert.Equal(t, InternalServerError.Type, problemDetails.Type)
	assert.Equal(t, InternalServerError.Code, problemDetails.Status)
	assert.Equal(t, InternalServerError.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotFound(t *testing.T) {
	problemDetails := NewNotFound()
	assert.Equal(t, NotFound.Type, problemDetails.Type)
	assert.Equal(t, NotFound.Code, problemDetails.Status)
	assert.Equal(t, NotFound.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadRequest(t *testing.T) {
	problemDetails := NewBadRequest()
	assert.Equal(t, BadRequest.Type, problemDetails.Type)
	assert.Equal(t, BadRequest.Code, problemDetails.Status)
	assert.Equal(t, BadRequest.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewContinue(t *testing.T) {
	problemDetails := NewContinue()
	assert.Equal(t, Continue.Type, problemDetails.Type)
	assert.Equal(t, Continue.Code, problemDetails.Status)
	assert.Equal(t, Continue.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewSwitchingProtocols(t *testing.T) {
	problemDetails := NewSwitchingProtocols()
	assert.Equal(t, SwitchingProtocols.Type, problemDetails.Type)
	assert.Equal(t, SwitchingProtocols.Code, problemDetails.Status)
	assert.Equal(t, SwitchingProtocols.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewOK(t *testing.T) {
	problemDetails := NewOK()
	assert.Equal(t, OK.Type, problemDetails.Type)
	assert.Equal(t, OK.Code, problemDetails.Status)
	assert.Equal(t, OK.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewCreated(t *testing.T) {
	problemDetails := NewCreated()
	assert.Equal(t, Created.Type, problemDetails.Type)
	assert.Equal(t, Created.Code, problemDetails.Status)
	assert.Equal(t, Created.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewAccepted(t *testing.T) {
	problemDetails := NewAccepted()
	assert.Equal(t, Accepted.Type, problemDetails.Type)
	assert.Equal(t, Accepted.Code, problemDetails.Status)
	assert.Equal(t, Accepted.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNoContent(t *testing.T) {
	problemDetails := NewNoContent()
	assert.Equal(t, NoContent.Type, problemDetails.Type)
	assert.Equal(t, NoContent.Code, problemDetails.Status)
	assert.Equal(t, NoContent.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewMovedPermanently(t *testing.T) {
	problemDetails := NewMovedPermanently()
	assert.Equal(t, MovedPermanently.Type, problemDetails.Type)
	assert.Equal(t, MovedPermanently.Code, problemDetails.Status)
	assert.Equal(t, MovedPermanently.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewFound(t *testing.T) {
	problemDetails := NewFound()
	assert.Equal(t, Found.Type, problemDetails.Type)
	assert.Equal(t, Found.Code, problemDetails.Status)
	assert.Equal(t, Found.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotModified(t *testing.T) {
	problemDetails := NewNotModified()
	assert.Equal(t, NotModified.Type, problemDetails.Type)
	assert.Equal(t, NotModified.Code, problemDetails.Status)
	assert.Equal(t, NotModified.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUnauthorized(t *testing.T) {
	problemDetails := NewUnauthorized()
	assert.Equal(t, Unauthorized.Type, problemDetails.Type)
	assert.Equal(t, Unauthorized.Code, problemDetails.Status)
	assert.Equal(t, Unauthorized.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewForbidden(t *testing.T) {
	problemDetails := NewForbidden()
	assert.Equal(t, Forbidden.Type, problemDetails.Type)
	assert.Equal(t, Forbidden.Code, problemDetails.Status)
	assert.Equal(t, Forbidden.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewMethodNotAllowed(t *testing.T) {
	problemDetails := NewMethodNotAllowed()
	assert.Equal(t, MethodNotAllowed.Type, problemDetails.Type)
	assert.Equal(t, MethodNotAllowed.Code, problemDetails.Status)
	assert.Equal(t, MethodNotAllowed.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewConflict(t *testing.T) {
	problemDetails := NewConflict()
	assert.Equal(t, Conflict.Type, problemDetails.Type)
	assert.Equal(t, Conflict.Code, problemDetails.Status)
	assert.Equal(t, Conflict.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotImplemented(t *testing.T) {
	problemDetails := NewNotImplemented()
	assert.Equal(t, NotImplemented.Type, problemDetails.Type)
	assert.Equal(t, NotImplemented.Code, problemDetails.Status)
	assert.Equal(t, NotImplemented.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewBadGateway(t *testing.T) {
	problemDetails := NewBadGateway()
	assert.Equal(t, BadGateway.Type, problemDetails.Type)
	assert.Equal(t, BadGateway.Code, problemDetails.Status)
	assert.Equal(t, BadGateway.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewServiceUnavailable(t *testing.T) {
	problemDetails := NewServiceUnavailable()
	assert.Equal(t, ServiceUnavailable.Type, problemDetails.Type)
	assert.Equal(t, ServiceUnavailable.Code, problemDetails.Status)
	assert.Equal(t, ServiceUnavailable.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewGatewayTimeout(t *testing.T) {
	problemDetails := NewGatewayTimeout()
	assert.Equal(t, GatewayTimeout.Type, problemDetails.Type)
	assert.Equal(t, GatewayTimeout.Code, problemDetails.Status)
	assert.Equal(t, GatewayTimeout.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewTeapot(t *testing.T) {
	problemDetails := NewTeapot()
	assert.Equal(t, Teapot.Type, problemDetails.Type)
	assert.Equal(t, Teapot.Code, problemDetails.Status)
	assert.Equal(t, Teapot.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewEarlyHints(t *testing.T) {
	problemDetails := NewEarlyHints()
	assert.Equal(t, EarlyHints.Type, problemDetails.Type)
	assert.Equal(t, EarlyHints.Code, problemDetails.Status)
	assert.Equal(t, EarlyHints.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNonAuthoritativeInfo(t *testing.T) {
	problemDetails := NewNonAuthoritativeInfo()
	assert.Equal(t, NonAuthoritativeInfo.Type, problemDetails.Type)
	assert.Equal(t, NonAuthoritativeInfo.Code, problemDetails.Status)
	assert.Equal(t, NonAuthoritativeInfo.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewResetContent(t *testing.T) {
	problemDetails := NewResetContent()
	assert.Equal(t, ResetContent.Type, problemDetails.Type)
	assert.Equal(t, ResetContent.Code, problemDetails.Status)
	assert.Equal(t, ResetContent.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewPartialContent(t *testing.T) {
	problemDetails := NewPartialContent()
	assert.Equal(t, PartialContent.Type, problemDetails.Type)
	assert.Equal(t, PartialContent.Code, problemDetails.Status)
	assert.Equal(t, PartialContent.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewIMUsed(t *testing.T) {
	problemDetails := NewIMUsed()
	assert.Equal(t, IMUsed.Type, problemDetails.Type)
	assert.Equal(t, IMUsed.Code, problemDetails.Status)
	assert.Equal(t, IMUsed.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewMultipleChoices(t *testing.T) {
	problemDetails := NewMultipleChoices()
	assert.Equal(t, MultipleChoices.Type, problemDetails.Type)
	assert.Equal(t, MultipleChoices.Code, problemDetails.Status)
	assert.Equal(t, MultipleChoices.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewSeeOther(t *testing.T) {
	problemDetails := NewSeeOther()
	assert.Equal(t, SeeOther.Type, problemDetails.Type)
	assert.Equal(t, SeeOther.Code, problemDetails.Status)
	assert.Equal(t, SeeOther.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUseProxy(t *testing.T) {
	problemDetails := NewUseProxy()
	assert.Equal(t, UseProxy.Type, problemDetails.Type)
	assert.Equal(t, UseProxy.Code, problemDetails.Status)
	assert.Equal(t, UseProxy.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewTemporaryRedirect(t *testing.T) {
	problemDetails := NewTemporaryRedirect()
	assert.Equal(t, TemporaryRedirect.Type, problemDetails.Type)
	assert.Equal(t, TemporaryRedirect.Code, problemDetails.Status)
	assert.Equal(t, TemporaryRedirect.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewPermanentRedirect(t *testing.T) {
	problemDetails := NewPermanentRedirect()
	assert.Equal(t, PermanentRedirect.Type, problemDetails.Type)
	assert.Equal(t, PermanentRedirect.Code, problemDetails.Status)
	assert.Equal(t, PermanentRedirect.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewPaymentRequired(t *testing.T) {
	problemDetails := NewPaymentRequired()
	assert.Equal(t, PaymentRequired.Type, problemDetails.Type)
	assert.Equal(t, PaymentRequired.Code, problemDetails.Status)
	assert.Equal(t, PaymentRequired.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotAcceptable(t *testing.T) {
	problemDetails := NewNotAcceptable()
	assert.Equal(t, NotAcceptable.Type, problemDetails.Type)
	assert.Equal(t, NotAcceptable.Code, problemDetails.Status)
	assert.Equal(t, NotAcceptable.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewProxyAuthenticationRequired(t *testing.T) {
	problemDetails := NewProxyAuthenticationRequired()
	assert.Equal(t, ProxyAuthenticationRequired.Type, problemDetails.Type)
	assert.Equal(t, ProxyAuthenticationRequired.Code, problemDetails.Status)
	assert.Equal(t, ProxyAuthenticationRequired.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewRequestTimeout(t *testing.T) {
	problemDetails := NewRequestTimeout()
	assert.Equal(t, RequestTimeout.Type, problemDetails.Type)
	assert.Equal(t, RequestTimeout.Code, problemDetails.Status)
	assert.Equal(t, RequestTimeout.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewGone(t *testing.T) {
	problemDetails := NewGone()
	assert.Equal(t, Gone.Type, problemDetails.Type)
	assert.Equal(t, Gone.Code, problemDetails.Status)
	assert.Equal(t, Gone.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewLengthRequired(t *testing.T) {
	problemDetails := NewLengthRequired()
	assert.Equal(t, LengthRequired.Type, problemDetails.Type)
	assert.Equal(t, LengthRequired.Code, problemDetails.Status)
	assert.Equal(t, LengthRequired.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewPreconditionFailed(t *testing.T) {
	problemDetails := NewPreconditionFailed()
	assert.Equal(t, PreconditionFailed.Type, problemDetails.Type)
	assert.Equal(t, PreconditionFailed.Code, problemDetails.Status)
	assert.Equal(t, PreconditionFailed.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewPayloadTooLarge(t *testing.T) {
	problemDetails := NewPayloadTooLarge()
	assert.Equal(t, PayloadTooLarge.Type, problemDetails.Type)
	assert.Equal(t, PayloadTooLarge.Code, problemDetails.Status)
	assert.Equal(t, PayloadTooLarge.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewURITooLong(t *testing.T) {
	problemDetails := NewURITooLong()
	assert.Equal(t, URITooLong.Type, problemDetails.Type)
	assert.Equal(t, URITooLong.Code, problemDetails.Status)
	assert.Equal(t, URITooLong.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUnsupportedMediaType(t *testing.T) {
	problemDetails := NewUnsupportedMediaType()
	assert.Equal(t, UnsupportedMediaType.Type, problemDetails.Type)
	assert.Equal(t, UnsupportedMediaType.Code, problemDetails.Status)
	assert.Equal(t, UnsupportedMediaType.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewRangeNotSatisfiable(t *testing.T) {
	problemDetails := NewRangeNotSatisfiable()
	assert.Equal(t, RangeNotSatisfiable.Type, problemDetails.Type)
	assert.Equal(t, RangeNotSatisfiable.Code, problemDetails.Status)
	assert.Equal(t, RangeNotSatisfiable.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewExpectationFailed(t *testing.T) {
	problemDetails := NewExpectationFailed()
	assert.Equal(t, ExpectationFailed.Type, problemDetails.Type)
	assert.Equal(t, ExpectationFailed.Code, problemDetails.Status)
	assert.Equal(t, ExpectationFailed.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewMisdirectedRequest(t *testing.T) {
	problemDetails := NewMisdirectedRequest()
	assert.Equal(t, MisdirectedRequest.Type, problemDetails.Type)
	assert.Equal(t, MisdirectedRequest.Code, problemDetails.Status)
	assert.Equal(t, MisdirectedRequest.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUpgradeRequired(t *testing.T) {
	problemDetails := NewUpgradeRequired()
	assert.Equal(t, UpgradeRequired.Type, problemDetails.Type)
	assert.Equal(t, UpgradeRequired.Code, problemDetails.Status)
	assert.Equal(t, UpgradeRequired.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewPreconditionRequired(t *testing.T) {
	problemDetails := NewPreconditionRequired()
	assert.Equal(t, PreconditionRequired.Type, problemDetails.Type)
	assert.Equal(t, PreconditionRequired.Code, problemDetails.Status)
	assert.Equal(t, PreconditionRequired.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewTooManyRequests(t *testing.T) {
	problemDetails := NewTooManyRequests()
	assert.Equal(t, TooManyRequests.Type, problemDetails.Type)
	assert.Equal(t, TooManyRequests.Code, problemDetails.Status)
	assert.Equal(t, TooManyRequests.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewRequestHeaderFieldsTooLarge(t *testing.T) {
	problemDetails := NewRequestHeaderFieldsTooLarge()
	assert.Equal(t, RequestHeaderFieldsTooLarge.Type, problemDetails.Type)
	assert.Equal(t, RequestHeaderFieldsTooLarge.Code, problemDetails.Status)
	assert.Equal(t, RequestHeaderFieldsTooLarge.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewUnavailableForLegalReasons(t *testing.T) {
	problemDetails := NewUnavailableForLegalReasons()
	assert.Equal(t, UnavailableForLegalReasons.Type, problemDetails.Type)
	assert.Equal(t, UnavailableForLegalReasons.Code, problemDetails.Status)
	assert.Equal(t, UnavailableForLegalReasons.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewHTTPVersionNotSupported(t *testing.T) {
	problemDetails := NewHTTPVersionNotSupported()
	assert.Equal(t, HTTPVersionNotSupported.Type, problemDetails.Type)
	assert.Equal(t, HTTPVersionNotSupported.Code, problemDetails.Status)
	assert.Equal(t, HTTPVersionNotSupported.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewVariantAlsoNegotiates(t *testing.T) {
	problemDetails := NewVariantAlsoNegotiates()
	assert.Equal(t, VariantAlsoNegotiates.Type, problemDetails.Type)
	assert.Equal(t, VariantAlsoNegotiates.Code, problemDetails.Status)
	assert.Equal(t, VariantAlsoNegotiates.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNotExtended(t *testing.T) {
	problemDetails := NewNotExtended()
	assert.Equal(t, NotExtended.Type, problemDetails.Type)
	assert.Equal(t, NotExtended.Code, problemDetails.Status)
	assert.Equal(t, NotExtended.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewNetworkAuthenticationRequired(t *testing.T) {
	problemDetails := NewNetworkAuthenticationRequired()
	assert.Equal(t, NetworkAuthenticationRequired.Type, problemDetails.Type)
	assert.Equal(t, NetworkAuthenticationRequired.Code, problemDetails.Status)
	assert.Equal(t, NetworkAuthenticationRequired.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
	doCommonAssertions(t, problemDetails)
}

func TestNewTooEarly(t *testing.T) {
	problemDetails := NewTooEarly()
	assert.Equal(t, TooEarly.Type, problemDetails.Type)
	assert.Equal(t, TooEarly.Code, problemDetails.Status)
	assert.Equal(t, TooEarly.Title, problemDetails.Title)
	assert.Equal(t, "", problemDetails.Detail)
	assert.Equal(t, 0, len(problemDetails.Extensions))
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
