// Code generated by goa v3.1.3, DO NOT EDIT.
//
// payments HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

import (
	payments "github.com/wondenge/at-notifiers/gen/payments"
	goa "goa.design/goa/v3/pkg"
)

// PaymentNotifierRequestBody is the type of the "payments" service
// "paymentNotifier" endpoint HTTP request body.
type PaymentNotifierRequestBody struct {
	// Unique transactionId generated for every payment.
	TransactionID string `form:"transactionId" json:"transactionId" xml:"transactionId"`
	// Category of the payment.
	Category string `form:"category" json:"category" xml:"category"`
	// Payment provider that facilitated this transaction.
	Provider string `form:"provider" json:"provider" xml:"provider"`
	// Unique ID generated by the payment provider for this transaction.
	ProviderRefID *string `form:"providerRefId,omitempty" json:"providerRefId,omitempty" xml:"providerRefId,omitempty"`
	// Name or number of channel that used to facilitate this payment.
	ProviderChannel string `form:"providerChannel" json:"providerChannel" xml:"providerChannel"`
	// Account name used by a mobile subscriber to initiate this transaction.
	ClientAccount *string `form:"clientAccount,omitempty" json:"clientAccount,omitempty" xml:"clientAccount,omitempty"`
	// Africa’s Talking Payment Product used to facilitate this transaction.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// Type of party providing the funds for this transaction (the Debit Party).
	SourceType string `form:"sourceType" json:"sourceType" xml:"sourceType"`
	// Unique identifier of the party that is providing the funds for this
	// transaction.
	Source string `form:"source" json:"source" xml:"source"`
	// Unique identifier of the party receiving funds in this transaction (the
	// Credit Party).
	DestinationType string `form:"destinationType" json:"destinationType" xml:"destinationType"`
	// Unique identifier of the party  receiving the funds for this transaction.
	Destination string `form:"destination" json:"destination" xml:"destination"`
	// Value being exchanged in this transaction.
	Value string `form:"value" json:"value" xml:"value"`
	// Transaction fee charged by Africa’s Talking for this transaction.
	TransactionFee *string `form:"transactionFee,omitempty" json:"transactionFee,omitempty" xml:"transactionFee,omitempty"`
	// Fee charged by a payment provider to facilitate this transaction.
	ProviderFee *string `form:"providerFee,omitempty" json:"providerFee,omitempty" xml:"providerFee,omitempty"`
	// The final status of this transaction
	Status string `form:"status" json:"status" xml:"status"`
	// A detailed description of this transaction.
	Description string `form:"description" json:"description" xml:"description"`
	// Any metadata that was sent by your application when it initiated this
	// transaction.
	RequestMetadata string `form:"requestMetadata" json:"requestMetadata" xml:"requestMetadata"`
	// Any additional data that we receive from a payment provider for a particular
	// transaction.
	ProviderMetadata string `form:"providerMetadata" json:"providerMetadata" xml:"providerMetadata"`
	// The date and time when a successful transaction was completed.
	TransactionDate *string `form:"transactionDate,omitempty" json:"transactionDate,omitempty" xml:"transactionDate,omitempty"`
}

// C2bNotifierRequestBody is the type of the "payments" service "c2bNotifier"
// endpoint HTTP request body.
type C2bNotifierRequestBody struct {
	// Payment provider that is facilitating this transaction
	Provider string `form:"provider" json:"provider" xml:"provider"`
	// Account name used by a mobile subscriber to initiate this transaction.
	ClientAccount *string `form:"clientAccount,omitempty" json:"clientAccount,omitempty" xml:"clientAccount,omitempty"`
	// Identifies the Africa’s Talking Payment Product used to facilitate this
	// transaction.
	ProductName string `form:"productName" json:"productName" xml:"productName"`
	// Phone number of the mobile subscriber who is initiating the C2B transaction.
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
	// Value being exchanged in this transaction
	Value string `form:"value" json:"value" xml:"value"`
	// Additional data received from a payment provider for a particular transaction
	ProviderMetadata map[string]string `form:"providerMetadata" json:"providerMetadata" xml:"providerMetadata"`
}

// B2cNotifierRequestBody is the type of the "payments" service "b2cNotifier"
// endpoint HTTP request body.
type B2cNotifierRequestBody struct {
	// The transaction id within Africa’s Talking.
	TransactionID *string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// The phone number of the mobile subscriber receiving the B2C payment.
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// The 3-digist ISO format currency for the value of this transaction
	CurrencyCode *string `form:"currencyCode,omitempty" json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	// Amount - in the provided currency - that the client will receive.
	Amount *float64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// The IPv4 address that initiated the B2C transaction.
	SourceIPAddress *string `form:"sourceIpAddress,omitempty" json:"sourceIpAddress,omitempty" xml:"sourceIpAddress,omitempty"`
	// A map of metadata associated with this request.
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// B2cNotifierResponseBody is the type of the "payments" service "b2cNotifier"
// endpoint HTTP response body.
type B2cNotifierResponseBody struct {
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewPaymentNotifierRequestBody builds the HTTP request body from the payload
// of the "paymentNotifier" endpoint of the "payments" service.
func NewPaymentNotifierRequestBody(p *payments.PaymentNotification) *PaymentNotifierRequestBody {
	body := &PaymentNotifierRequestBody{
		TransactionID:    p.TransactionID,
		Category:         p.Category,
		Provider:         p.Provider,
		ProviderRefID:    p.ProviderRefID,
		ProviderChannel:  p.ProviderChannel,
		ClientAccount:    p.ClientAccount,
		ProductName:      p.ProductName,
		SourceType:       p.SourceType,
		Source:           p.Source,
		DestinationType:  p.DestinationType,
		Destination:      p.Destination,
		Value:            p.Value,
		TransactionFee:   p.TransactionFee,
		ProviderFee:      p.ProviderFee,
		Status:           p.Status,
		Description:      p.Description,
		RequestMetadata:  p.RequestMetadata,
		ProviderMetadata: p.ProviderMetadata,
		TransactionDate:  p.TransactionDate,
	}
	return body
}

// NewC2bNotifierRequestBody builds the HTTP request body from the payload of
// the "c2bNotifier" endpoint of the "payments" service.
func NewC2bNotifierRequestBody(p *payments.C2BValidationNotification) *C2bNotifierRequestBody {
	body := &C2bNotifierRequestBody{
		Provider:      p.Provider,
		ClientAccount: p.ClientAccount,
		ProductName:   p.ProductName,
		PhoneNumber:   p.PhoneNumber,
		Value:         p.Value,
	}
	if p.ProviderMetadata != nil {
		body.ProviderMetadata = make(map[string]string, len(p.ProviderMetadata))
		for key, val := range p.ProviderMetadata {
			tk := key
			tv := val
			body.ProviderMetadata[tk] = tv
		}
	}
	return body
}

// NewB2cNotifierRequestBody builds the HTTP request body from the payload of
// the "b2cNotifier" endpoint of the "payments" service.
func NewB2cNotifierRequestBody(p *payments.B2CValidationNotificationPayload) *B2cNotifierRequestBody {
	body := &B2cNotifierRequestBody{
		TransactionID:   p.TransactionID,
		PhoneNumber:     p.PhoneNumber,
		CurrencyCode:    p.CurrencyCode,
		Amount:          p.Amount,
		SourceIPAddress: p.SourceIPAddress,
	}
	if p.Metadata != nil {
		body.Metadata = make(map[string]string, len(p.Metadata))
		for key, val := range p.Metadata {
			tk := key
			tv := val
			body.Metadata[tk] = tv
		}
	}
	return body
}

// NewB2cNotifierB2CValidationNotificationResponseCreated builds a "payments"
// service "b2cNotifier" endpoint result from a HTTP "Created" response.
func NewB2cNotifierB2CValidationNotificationResponseCreated(body *B2cNotifierResponseBody) *payments.B2CValidationNotificationResponse {
	v := &payments.B2CValidationNotificationResponse{
		Status: body.Status,
	}

	return v
}

// ValidateB2cNotifierResponseBody runs the validations defined on
// B2cNotifierResponseBody
func ValidateB2cNotifierResponseBody(body *B2cNotifierResponseBody) (err error) {
	if body.Status != nil {
		if !(*body.Status == "Validated" || *body.Status == "Failed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"Validated", "Failed"}))
		}
	}
	return
}