package deboxapi

import (
	"encoding/json"
)

// BadResponseError : special sls error, not valid json format
type BadResponseError struct {
	RespBody   string
	RespHeader map[string][]string
	HTTPCode   int
}

func (e BadResponseError) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e BadResponseError) Error() string {
	return e.String()
}

// NewBadResponseError ...
func NewBadResponseError(body string, header map[string][]string, httpCode int) *BadResponseError {
	return &BadResponseError{
		RespBody:   body,
		RespHeader: header,
		HTTPCode:   httpCode,
	}
}

// mockErrorRetry : for mock the error retry logic
type mockErrorRetry struct {
	Err      Error
	RetryCnt int // RetryCnt-- after each retry. When RetryCnt > 0, return Err, else return nil, if set it BigUint, it equivalents to always failing.
}

func (e mockErrorRetry) Error() string {
	return e.Err.String()
}

// Error is an error containing extra information returned by the DeBox API.
type Error struct {
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
	// ResponseParameters
	HTTPCode int32 `json:"httpCode"`
	// Code      string `json:"errorCode"`
	// Message   string `json:"errorMessage"`
	RequestID string `json:"requestID"`
}

func (e Error) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e Error) Error() string {
	return e.Message
}
