// Code generated by goa v3.1.3, DO NOT EDIT.
//
// airtime HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

import (
	airtime "github.com/wondenge/at-notifiers/gen/airtime"
	goa "goa.design/goa/v3/pkg"
)

// ValidationRequestBody is the type of the "airtime" service "validation"
// endpoint HTTP request body.
type ValidationRequestBody struct {
	// The transaction id within Africa’s Talking.
	TransactionID *string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// The phone number of the mobile subscriber receiving the airtime.
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// The 3-digist ISO format currency for the value of this transaction
	CurrencyCode *string `form:"currencyCode,omitempty" json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	// Amount - in the provided currency - that the client will receive.
	Amount *float64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

// StatusRequestBody is the type of the "airtime" service "status" endpoint
// HTTP request body.
type StatusRequestBody struct {
	// The request ID sent back as a response to the airtime send request.
	RequestID *string `form:"requestId,omitempty" json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The Transaction status.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// ValidationResponseBody is the type of the "airtime" service "validation"
// endpoint HTTP response body.
type ValidationResponseBody struct {
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewValidationRequestBody builds the HTTP request body from the payload of
// the "validation" endpoint of the "airtime" service.
func NewValidationRequestBody(p *airtime.AirtimeValidationPayload) *ValidationRequestBody {
	body := &ValidationRequestBody{
		TransactionID: p.TransactionID,
		PhoneNumber:   p.PhoneNumber,
		CurrencyCode:  p.CurrencyCode,
		Amount:        p.Amount,
	}
	return body
}

// NewStatusRequestBody builds the HTTP request body from the payload of the
// "status" endpoint of the "airtime" service.
func NewStatusRequestBody(p *airtime.AirtimeStatus) *StatusRequestBody {
	body := &StatusRequestBody{
		RequestID: p.RequestID,
		Status:    p.Status,
	}
	return body
}

// NewValidationAirtimeValidationResponseCreated builds a "airtime" service
// "validation" endpoint result from a HTTP "Created" response.
func NewValidationAirtimeValidationResponseCreated(body *ValidationResponseBody) *airtime.AirtimeValidationResponse {
	v := &airtime.AirtimeValidationResponse{
		Status: body.Status,
	}

	return v
}

// ValidateValidationResponseBody runs the validations defined on
// ValidationResponseBody
func ValidateValidationResponseBody(body *ValidationResponseBody) (err error) {
	if body.Status != nil {
		if !(*body.Status == "Validated" || *body.Status == "Failed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"Validated", "Failed"}))
		}
	}
	return
}
