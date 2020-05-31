// Code generated by goa v3.1.3, DO NOT EDIT.
//
// airtime service
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package airtime

import (
	"context"
)

// Service is the airtime service interface.
type Service interface {
	// Airtime Validation Notifications
	Validation(context.Context, *AirtimeValidationPayload) (res *AirtimeValidationResponse, err error)
	// Airtime Status Notifications
	Status(context.Context, *AirtimeStatus) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "airtime"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"validation", "status"}

// AirtimeValidationPayload is the payload type of the airtime service
// validation method.
type AirtimeValidationPayload struct {
	// The transaction id within Africa’s Talking.
	TransactionID *string
	// The phone number of the mobile subscriber receiving the airtime.
	PhoneNumber *string
	// The 3-digist ISO format currency for the value of this transaction
	CurrencyCode *string
	// Amount - in the provided currency - that the client will receive.
	Amount *float64
}

// AirtimeValidationResponse is the result type of the airtime service
// validation method.
type AirtimeValidationResponse struct {
	Status *string
}

// AirtimeStatus is the payload type of the airtime service status method.
type AirtimeStatus struct {
	// The request ID sent back as a response to the airtime send request.
	RequestID *string
	// The Transaction status.
	Status *string
}
