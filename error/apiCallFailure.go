package error

import "fmt"

// NewAPIClientError creates a new API call error
func NewAPIClientError(apiURL string, httpStatusCode *int, httpResponseBody *string, err error) APIClientError {
	return &apiClientErrorImpl{apiURL, httpStatusCode, httpResponseBody, err.Error()}
}

// APIClientError represents an database query failure error interface
type APIClientError interface {
	Error() string
	GetAPIURL() string
	GetHTTPStatusCode() *int
	GetHTTPResponseBody() *string
}

type apiClientErrorImpl struct {
	apiURL           string
	httpStatusCode   *int
	httpResponseBody *string
	error            string
}

// Error returns the error string
func (e apiClientErrorImpl) Error() string {
	return fmt.Sprintf("%v", e.error)
}

// GetAPIURL gets API URL
func (e apiClientErrorImpl) GetAPIURL() string {
	return e.apiURL
}

// GetHTTPStatusCode gets http status code
func (e apiClientErrorImpl) GetHTTPStatusCode() *int {
	return e.httpStatusCode
}

// GetHTTPResponseBody gets http status code
func (e apiClientErrorImpl) GetHTTPResponseBody() *string {
	return e.httpResponseBody
}
