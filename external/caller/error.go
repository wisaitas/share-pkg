package caller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/wisaitas/shared-pkg/response"
)

type ServiceError struct {
	StatusCode int
	Response   *response.ApiResponse[any]
}

func (e *ServiceError) Error() string {
	if e.Response != nil {
		return fmt.Sprintf("service error %d: %s", e.StatusCode, e.Response.Message)
	}
	return fmt.Sprintf("service error %d", e.StatusCode)
}

// IsStatusCategory checks if the error is a ServiceError with a status code in the specified range
func IsStatusCategory(err error, minCode, maxCode int) bool {
	var serviceErr *ServiceError
	if errors.As(err, &serviceErr) {
		return serviceErr.StatusCode >= minCode && serviceErr.StatusCode < maxCode
	}
	return false
}

// IsStatusCode checks if the error is a ServiceError with the specified status code
func IsStatusCode(err error, code int) bool {
	var serviceErr *ServiceError
	if errors.As(err, &serviceErr) {
		return serviceErr.StatusCode == code
	}
	return false
}

// Status Category Checkers
func Is1xxInformational(err error) bool {
	return IsStatusCategory(err, 100, 200)
}

func Is2xxSuccess(err error) bool {
	return IsStatusCategory(err, 200, 300)
}

func Is3xxRedirection(err error) bool {
	return IsStatusCategory(err, 300, 400)
}

func Is4xxClientError(err error) bool {
	return IsStatusCategory(err, 400, 500)
}

func Is5xxServerError(err error) bool {
	return IsStatusCategory(err, 500, 600)
}

// 1xx Informational
func IsContinue(err error) bool {
	return IsStatusCode(err, http.StatusContinue)
}

func IsSwitchingProtocols(err error) bool {
	return IsStatusCode(err, http.StatusSwitchingProtocols)
}

func IsProcessing(err error) bool {
	return IsStatusCode(err, http.StatusProcessing)
}

// 2xx Success
func IsOK(err error) bool {
	return IsStatusCode(err, http.StatusOK)
}

func IsCreated(err error) bool {
	return IsStatusCode(err, http.StatusCreated)
}

func IsAccepted(err error) bool {
	return IsStatusCode(err, http.StatusAccepted)
}

func IsNoContent(err error) bool {
	return IsStatusCode(err, http.StatusNoContent)
}

// 3xx Redirection
func IsMultipleChoices(err error) bool {
	return IsStatusCode(err, http.StatusMultipleChoices)
}

func IsMovedPermanently(err error) bool {
	return IsStatusCode(err, http.StatusMovedPermanently)
}

func IsFound(err error) bool {
	return IsStatusCode(err, http.StatusFound)
}

func IsSeeOther(err error) bool {
	return IsStatusCode(err, http.StatusSeeOther)
}

func IsNotModified(err error) bool {
	return IsStatusCode(err, http.StatusNotModified)
}

func IsTemporaryRedirect(err error) bool {
	return IsStatusCode(err, http.StatusTemporaryRedirect)
}

func IsPermanentRedirect(err error) bool {
	return IsStatusCode(err, http.StatusPermanentRedirect)
}

// 4xx Client Error
func IsBadRequest(err error) bool {
	return IsStatusCode(err, http.StatusBadRequest)
}

func IsUnauthorized(err error) bool {
	return IsStatusCode(err, http.StatusUnauthorized)
}

func IsPaymentRequired(err error) bool {
	return IsStatusCode(err, http.StatusPaymentRequired)
}

func IsForbidden(err error) bool {
	return IsStatusCode(err, http.StatusForbidden)
}

func IsNotFound(err error) bool {
	return IsStatusCode(err, http.StatusNotFound)
}

func IsMethodNotAllowed(err error) bool {
	return IsStatusCode(err, http.StatusMethodNotAllowed)
}

func IsNotAcceptable(err error) bool {
	return IsStatusCode(err, http.StatusNotAcceptable)
}

func IsProxyAuthRequired(err error) bool {
	return IsStatusCode(err, http.StatusProxyAuthRequired)
}

func IsRequestTimeout(err error) bool {
	return IsStatusCode(err, http.StatusRequestTimeout)
}

func IsConflict(err error) bool {
	return IsStatusCode(err, http.StatusConflict)
}

func IsGone(err error) bool {
	return IsStatusCode(err, http.StatusGone)
}

func IsLengthRequired(err error) bool {
	return IsStatusCode(err, http.StatusLengthRequired)
}

func IsPreconditionFailed(err error) bool {
	return IsStatusCode(err, http.StatusPreconditionFailed)
}

func IsPayloadTooLarge(err error) bool {
	return IsStatusCode(err, http.StatusRequestEntityTooLarge)
}

func IsURITooLong(err error) bool {
	return IsStatusCode(err, http.StatusRequestURITooLong)
}

func IsUnsupportedMediaType(err error) bool {
	return IsStatusCode(err, http.StatusUnsupportedMediaType)
}

func IsRangeNotSatisfiable(err error) bool {
	return IsStatusCode(err, http.StatusRequestedRangeNotSatisfiable)
}

func IsExpectationFailed(err error) bool {
	return IsStatusCode(err, http.StatusExpectationFailed)
}

func IsTeapot(err error) bool {
	return IsStatusCode(err, http.StatusTeapot)
}

func IsMisdirectedRequest(err error) bool {
	return IsStatusCode(err, http.StatusMisdirectedRequest)
}

func IsUnprocessableEntity(err error) bool {
	return IsStatusCode(err, http.StatusUnprocessableEntity)
}

func IsLocked(err error) bool {
	return IsStatusCode(err, http.StatusLocked)
}

func IsFailedDependency(err error) bool {
	return IsStatusCode(err, http.StatusFailedDependency)
}

func IsTooEarly(err error) bool {
	return IsStatusCode(err, http.StatusTooEarly)
}

func IsUpgradeRequired(err error) bool {
	return IsStatusCode(err, http.StatusUpgradeRequired)
}

func IsPreconditionRequired(err error) bool {
	return IsStatusCode(err, http.StatusPreconditionRequired)
}

func IsTooManyRequests(err error) bool {
	return IsStatusCode(err, http.StatusTooManyRequests)
}

func IsRequestHeaderFieldsTooLarge(err error) bool {
	return IsStatusCode(err, http.StatusRequestHeaderFieldsTooLarge)
}

func IsUnavailableForLegalReasons(err error) bool {
	return IsStatusCode(err, http.StatusUnavailableForLegalReasons)
}

// 5xx Server Error
func IsInternalServerError(err error) bool {
	return IsStatusCode(err, http.StatusInternalServerError)
}

func IsNotImplemented(err error) bool {
	return IsStatusCode(err, http.StatusNotImplemented)
}

func IsBadGateway(err error) bool {
	return IsStatusCode(err, http.StatusBadGateway)
}

func IsServiceUnavailable(err error) bool {
	return IsStatusCode(err, http.StatusServiceUnavailable)
}

func IsGatewayTimeout(err error) bool {
	return IsStatusCode(err, http.StatusGatewayTimeout)
}

func IsHTTPVersionNotSupported(err error) bool {
	return IsStatusCode(err, http.StatusHTTPVersionNotSupported)
}

func IsVariantAlsoNegotiates(err error) bool {
	return IsStatusCode(err, http.StatusVariantAlsoNegotiates)
}

func IsInsufficientStorage(err error) bool {
	return IsStatusCode(err, http.StatusInsufficientStorage)
}

func IsLoopDetected(err error) bool {
	return IsStatusCode(err, http.StatusLoopDetected)
}

func IsNotExtended(err error) bool {
	return IsStatusCode(err, http.StatusNotExtended)
}

func IsNetworkAuthenticationRequired(err error) bool {
	return IsStatusCode(err, http.StatusNetworkAuthenticationRequired)
}

// GetStatusCode returns the status code from a ServiceError
func GetStatusCode(err error) (int, bool) {
	var serviceErr *ServiceError
	if errors.As(err, &serviceErr) {
		return serviceErr.StatusCode, true
	}
	return 0, false
}

// GetStatusText returns a formatted status text for a ServiceError
func GetStatusText(err error) string {
	var serviceErr *ServiceError
	if errors.As(err, &serviceErr) {
		return fmt.Sprintf("%d %s", serviceErr.StatusCode, http.StatusText(serviceErr.StatusCode))
	}
	return ""
}

// HasStatusCode checks if the error has the specified status code
func HasStatusCode(err error, statusCode int) bool {
	return IsStatusCode(err, statusCode)
}

// HasAnyStatusCode checks if the error has any of the specified status codes
func HasAnyStatusCode(err error, statusCodes ...int) bool {
	var serviceErr *ServiceError
	if errors.As(err, &serviceErr) {
		for _, code := range statusCodes {
			if serviceErr.StatusCode == code {
				return true
			}
		}
	}
	return false
}
