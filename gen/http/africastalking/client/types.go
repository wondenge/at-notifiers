// Code generated by goa v3.1.3, DO NOT EDIT.
//
// africastalking HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

import (
	africastalking "github.com/wondenge/at-notifiers/gen/africastalking"
	africastalkingviews "github.com/wondenge/at-notifiers/gen/africastalking/views"
	goa "goa.design/goa/v3/pkg"
)

// DeliveryReportNotifierRequestBody is the type of the "africastalking"
// service "delivery_report_notifier" endpoint HTTP request body.
type DeliveryReportNotifierRequestBody struct {
	// A unique identifier for each message.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The status of the message.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Mobile phone number message was sent out to.
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// A unique identifier for the Telco that handled the message.
	NetworkCode *string `form:"networkCode,omitempty" json:"networkCode,omitempty" xml:"networkCode,omitempty"`
	// Only provided if status is Rejected or Failed.
	FailureReason *string `form:"failureReason,omitempty" json:"failureReason,omitempty" xml:"failureReason,omitempty"`
	// Number of times the request to send a message to the device was retried
	// before it succeeded or definitely failed.
	RetryCount *string `form:"retryCount,omitempty" json:"retryCount,omitempty" xml:"retryCount,omitempty"`
}

// IncomingMessageNotifierRequestBody is the type of the "africastalking"
// service "incoming_message_notifier" endpoint HTTP request body.
type IncomingMessageNotifierRequestBody struct {
	// The date and time when the message was received.
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// The number that sent the message.
	From *string `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	// The internal ID that we use to store this message.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Parameter required when responding to an on-demand user request with a
	// premium message.
	LinkID *string `form:"linkId,omitempty" json:"linkId,omitempty" xml:"linkId,omitempty"`
	// The message content.
	Text *string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// The number to which the message was sent.
	To *string `form:"to,omitempty" json:"to,omitempty" xml:"to,omitempty"`
	// A unique identifier for the telco that handled the message.
	NetworkCode *string `form:"networkCode,omitempty" json:"networkCode,omitempty" xml:"networkCode,omitempty"`
}

// BulkOptOutNotifierRequestBody is the type of the "africastalking" service
// "bulk_optOut_notifier" endpoint HTTP request body.
type BulkOptOutNotifierRequestBody struct {
	// Shortcode/Alphanumeric Sender ID the user opted out from.
	SenderID *string `form:"senderId,omitempty" json:"senderId,omitempty" xml:"senderId,omitempty"`
	// Mobile phone number of the subscriber who opted out.
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

// SubNotifierRequestBody is the type of the "africastalking" service
// "sub_notifier" endpoint HTTP request body.
type SubNotifierRequestBody struct {
	// Mobile phone number to subscribe or unsubscribe.
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// The short code that has this product.
	ShortCode *string `form:"shortCode,omitempty" json:"shortCode,omitempty" xml:"shortCode,omitempty"`
	// The product keyword that the user has subscribed or unsubscribed from.
	Keyword *string `form:"keyword,omitempty" json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The type of the update.
	UpdateType *string `form:"updateType,omitempty" json:"updateType,omitempty" xml:"updateType,omitempty"`
}

// VoiceNotifierRequestBody is the type of the "africastalking" service
// "voice_notifier" endpoint HTTP request body.
type VoiceNotifierRequestBody struct {
	// Lets us know whether the call is in session state
	IsActive string `form:"isActive,omitempty" json:"isActive,omitempty" xml:"isActive,omitempty"`
	// A unique identifier generated during each call session
	SessionID *string `form:"sessionId,omitempty" json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Whether this is an inbound or outbound call
	Direction *string `form:"direction,omitempty" json:"direction,omitempty" xml:"direction,omitempty"`
	// Africa’s Talking phone number, in international format
	DestinationNumber *string `form:"destinationNumber,omitempty" json:"destinationNumber,omitempty" xml:"destinationNumber,omitempty"`
	// The phone number of the phone user in the call, in international format.
	CallerNumber *string `form:"callerNumber,omitempty" json:"callerNumber,omitempty" xml:"callerNumber,omitempty"`
	// The code of the country the callerNumber is calling from.
	CallerCountryCode *string `form:"callerCountryCode,omitempty" json:"callerCountryCode,omitempty" xml:"callerCountryCode,omitempty"`
	// The time the call began.
	CallStartTime *string `form:"callStartTime,omitempty" json:"callStartTime,omitempty" xml:"callStartTime,omitempty"`
	// The digits that a user enters in response to a getDigits request
	DtmfDigits *string `form:"dtmfDigits,omitempty" json:"dtmfDigits,omitempty" xml:"dtmfDigits,omitempty"`
	// The URL of the recording made for this call
	RecordingURL *string `form:"recordingUrl,omitempty" json:"recordingUrl,omitempty" xml:"recordingUrl,omitempty"`
	// The duration of the call in seconds.
	DurationInSeconds *string `form:"durationInSeconds,omitempty" json:"durationInSeconds,omitempty" xml:"durationInSeconds,omitempty"`
	// The currency used to bill this call (e.g KES, USD, GBP).
	CurrencyCode *string `form:"currencyCode,omitempty" json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	// The total cost of the call.
	Amount *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	// The final status of the call.
	CallSessionState *string `form:"callSessionState,omitempty" json:"callSessionState,omitempty" xml:"callSessionState,omitempty"`
	// The number which a call was forwarded to if the Dial action was used.
	DialDestinationNumber *string `form:"dialDestinationNumber,omitempty" json:"dialDestinationNumber,omitempty" xml:"dialDestinationNumber,omitempty"`
	// The duration of the dialed call if the Dial action was used.
	DialDurationInSeconds *string `form:"dialDurationInSeconds,omitempty" json:"dialDurationInSeconds,omitempty" xml:"dialDurationInSeconds,omitempty"`
	// The time the dial action began if the Dial action was used.
	DialStartTime *string `form:"dialStartTime,omitempty" json:"dialStartTime,omitempty" xml:"dialStartTime,omitempty"`
	// The reason a call could have ended
	HangupCause *string `form:"hangupCause,omitempty" json:"hangupCause,omitempty" xml:"hangupCause,omitempty"`
}

// TransferEventNotifierRequestBody is the type of the "africastalking" service
// "transfer_event_notifier" endpoint HTTP request body.
type TransferEventNotifierRequestBody struct {
	CallSessionState *string `form:"callSessionState,omitempty" json:"callSessionState,omitempty" xml:"callSessionState,omitempty"`
	IsActive         string  `form:"isActive,omitempty" json:"isActive,omitempty" xml:"isActive,omitempty"`
	Status           *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// +2347XXXXXXXXX:20, (20 is the duration in seconds)
	CallTransferParam *string `form:"callTransferParam,omitempty" json:"callTransferParam,omitempty" xml:"callTransferParam,omitempty"`
	// Number call was transferred to
	CallTransferredToNumber *string `form:"callTransferredToNumber,omitempty" json:"callTransferredToNumber,omitempty" xml:"callTransferredToNumber,omitempty"`
	CallTransferState       *string `form:"callTransferState,omitempty" json:"callTransferState,omitempty" xml:"callTransferState,omitempty"`
	CallTransferHangupCause *string `form:"callTransferHangupCause,omitempty" json:"callTransferHangupCause,omitempty" xml:"callTransferHangupCause,omitempty"`
}

// UssdNotifierRequestBody is the type of the "africastalking" service
// "ussd_notifier" endpoint HTTP request body.
type UssdNotifierRequestBody struct {
	// A unique value generated when the session starts.
	SessionID *string `form:"sessionId,omitempty" json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Mobile number of the subscriber interacting with USSD application.
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Telco of the mobile number interacting with USSD application.
	NetworkCode *string `form:"networkCode,omitempty" json:"networkCode,omitempty" xml:"networkCode,omitempty"`
	// USSD code assigned to application.
	ServiceCode *string `form:"serviceCode,omitempty" json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// Shows the user input.
	Text *string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
}

// ValidationNotifierRequestBody is the type of the "africastalking" service
// "validation_notifier" endpoint HTTP request body.
type ValidationNotifierRequestBody struct {
	// The transaction id within Africa’s Talking.
	TransactionID *string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// The phone number of the mobile subscriber receiving the airtime.
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// The 3-digist ISO format currency for the value of this transaction
	CurrencyCode *string `form:"currencyCode,omitempty" json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	// Amount - in the provided currency - that the client will receive.
	Amount *float64 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

// StatusNotifierRequestBody is the type of the "africastalking" service
// "status_notifier" endpoint HTTP request body.
type StatusNotifierRequestBody struct {
	// The request ID sent back as a response to the airtime send request.
	RequestID *string `form:"requestId,omitempty" json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The Transaction status.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// PaymentNotifierRequestBody is the type of the "africastalking" service
// "payment_notifier" endpoint HTTP request body.
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

// C2bValidationNotifierRequestBody is the type of the "africastalking" service
// "c2b_validation_notifier" endpoint HTTP request body.
type C2bValidationNotifierRequestBody struct {
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

// B2cValidationNotifierRequestBody is the type of the "africastalking" service
// "b2c_validation_notifier" endpoint HTTP request body.
type B2cValidationNotifierRequestBody struct {
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

// IotNotifierRequestBody is the type of the "africastalking" service
// "iot_notifier" endpoint HTTP request body.
type IotNotifierRequestBody struct {
	// The MQTT packet sent by the publishing device.
	Payload *string `form:"payload,omitempty" json:"payload,omitempty" xml:"payload,omitempty"`
	// Message channel to which the message was sent by the publishing MQTT client
	Topic *string `form:"topic,omitempty" json:"topic,omitempty" xml:"topic,omitempty"`
}

// UssdNotifierResponseBody is the type of the "africastalking" service
// "ussd_notifier" endpoint HTTP response body.
type UssdNotifierResponseBody struct {
	// Plain text response back to AT gateway
	Response *string `form:"response,omitempty" json:"response,omitempty" xml:"response,omitempty"`
}

// ValidationNotifierResponseBody is the type of the "africastalking" service
// "validation_notifier" endpoint HTTP response body.
type ValidationNotifierResponseBody struct {
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// B2cValidationNotifierResponseBody is the type of the "africastalking"
// service "b2c_validation_notifier" endpoint HTTP response body.
type B2cValidationNotifierResponseBody struct {
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewDeliveryReportNotifierRequestBody builds the HTTP request body from the
// payload of the "delivery_report_notifier" endpoint of the "africastalking"
// service.
func NewDeliveryReportNotifierRequestBody(p *africastalking.DeliveryReportPayload) *DeliveryReportNotifierRequestBody {
	body := &DeliveryReportNotifierRequestBody{
		ID:            p.ID,
		Status:        p.Status,
		PhoneNumber:   p.PhoneNumber,
		NetworkCode:   p.NetworkCode,
		FailureReason: p.FailureReason,
		RetryCount:    p.RetryCount,
	}
	return body
}

// NewIncomingMessageNotifierRequestBody builds the HTTP request body from the
// payload of the "incoming_message_notifier" endpoint of the "africastalking"
// service.
func NewIncomingMessageNotifierRequestBody(p *africastalking.IncomingMessagePayload) *IncomingMessageNotifierRequestBody {
	body := &IncomingMessageNotifierRequestBody{
		Date:        p.Date,
		From:        p.From,
		ID:          p.ID,
		LinkID:      p.LinkID,
		Text:        p.Text,
		To:          p.To,
		NetworkCode: p.NetworkCode,
	}
	return body
}

// NewBulkOptOutNotifierRequestBody builds the HTTP request body from the
// payload of the "bulk_optOut_notifier" endpoint of the "africastalking"
// service.
func NewBulkOptOutNotifierRequestBody(p *africastalking.BulkSMSOptOutPayload) *BulkOptOutNotifierRequestBody {
	body := &BulkOptOutNotifierRequestBody{
		SenderID:    p.SenderID,
		PhoneNumber: p.PhoneNumber,
	}
	return body
}

// NewSubNotifierRequestBody builds the HTTP request body from the payload of
// the "sub_notifier" endpoint of the "africastalking" service.
func NewSubNotifierRequestBody(p *africastalking.SubNotificationPayload) *SubNotifierRequestBody {
	body := &SubNotifierRequestBody{
		PhoneNumber: p.PhoneNumber,
		ShortCode:   p.ShortCode,
		Keyword:     p.Keyword,
		UpdateType:  p.UpdateType,
	}
	return body
}

// NewVoiceNotifierRequestBody builds the HTTP request body from the payload of
// the "voice_notifier" endpoint of the "africastalking" service.
func NewVoiceNotifierRequestBody(p *africastalking.VoiceNotificationPayload) *VoiceNotifierRequestBody {
	body := &VoiceNotifierRequestBody{
		IsActive:              p.IsActive,
		SessionID:             p.SessionID,
		Direction:             p.Direction,
		DestinationNumber:     p.DestinationNumber,
		CallerNumber:          p.CallerNumber,
		CallerCountryCode:     p.CallerCountryCode,
		CallStartTime:         p.CallStartTime,
		DtmfDigits:            p.DtmfDigits,
		RecordingURL:          p.RecordingURL,
		DurationInSeconds:     p.DurationInSeconds,
		CurrencyCode:          p.CurrencyCode,
		Amount:                p.Amount,
		CallSessionState:      p.CallSessionState,
		DialDestinationNumber: p.DialDestinationNumber,
		DialDurationInSeconds: p.DialDurationInSeconds,
		DialStartTime:         p.DialStartTime,
		HangupCause:           p.HangupCause,
	}
	return body
}

// NewTransferEventNotifierRequestBody builds the HTTP request body from the
// payload of the "transfer_event_notifier" endpoint of the "africastalking"
// service.
func NewTransferEventNotifierRequestBody(p *africastalking.TransferEventPayload) *TransferEventNotifierRequestBody {
	body := &TransferEventNotifierRequestBody{
		CallSessionState:        p.CallSessionState,
		IsActive:                p.IsActive,
		Status:                  p.Status,
		CallTransferParam:       p.CallTransferParam,
		CallTransferredToNumber: p.CallTransferredToNumber,
		CallTransferState:       p.CallTransferState,
		CallTransferHangupCause: p.CallTransferHangupCause,
	}
	return body
}

// NewUssdNotifierRequestBody builds the HTTP request body from the payload of
// the "ussd_notifier" endpoint of the "africastalking" service.
func NewUssdNotifierRequestBody(p *africastalking.USSDPayload) *UssdNotifierRequestBody {
	body := &UssdNotifierRequestBody{
		SessionID:   p.SessionID,
		PhoneNumber: p.PhoneNumber,
		NetworkCode: p.NetworkCode,
		ServiceCode: p.ServiceCode,
		Text:        p.Text,
	}
	return body
}

// NewValidationNotifierRequestBody builds the HTTP request body from the
// payload of the "validation_notifier" endpoint of the "africastalking"
// service.
func NewValidationNotifierRequestBody(p *africastalking.AirtimeValidationPayload) *ValidationNotifierRequestBody {
	body := &ValidationNotifierRequestBody{
		TransactionID: p.TransactionID,
		PhoneNumber:   p.PhoneNumber,
		CurrencyCode:  p.CurrencyCode,
		Amount:        p.Amount,
	}
	return body
}

// NewStatusNotifierRequestBody builds the HTTP request body from the payload
// of the "status_notifier" endpoint of the "africastalking" service.
func NewStatusNotifierRequestBody(p *africastalking.AirtimeStatusPayload) *StatusNotifierRequestBody {
	body := &StatusNotifierRequestBody{
		RequestID: p.RequestID,
		Status:    p.Status,
	}
	return body
}

// NewPaymentNotifierRequestBody builds the HTTP request body from the payload
// of the "payment_notifier" endpoint of the "africastalking" service.
func NewPaymentNotifierRequestBody(p *africastalking.PaymentNotificationPayload) *PaymentNotifierRequestBody {
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

// NewC2bValidationNotifierRequestBody builds the HTTP request body from the
// payload of the "c2b_validation_notifier" endpoint of the "africastalking"
// service.
func NewC2bValidationNotifierRequestBody(p *africastalking.C2BValidationNotificationPayload) *C2bValidationNotifierRequestBody {
	body := &C2bValidationNotifierRequestBody{
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

// NewB2cValidationNotifierRequestBody builds the HTTP request body from the
// payload of the "b2c_validation_notifier" endpoint of the "africastalking"
// service.
func NewB2cValidationNotifierRequestBody(p *africastalking.B2CValidationNotificationPayload) *B2cValidationNotifierRequestBody {
	body := &B2cValidationNotifierRequestBody{
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

// NewIotNotifierRequestBody builds the HTTP request body from the payload of
// the "iot_notifier" endpoint of the "africastalking" service.
func NewIotNotifierRequestBody(p *africastalking.IoTNotificationPayload) *IotNotifierRequestBody {
	body := &IotNotifierRequestBody{
		Payload: p.Payload,
		Topic:   p.Topic,
	}
	return body
}

// NewUssdNotifierUSSDResponseCreated builds a "africastalking" service
// "ussd_notifier" endpoint result from a HTTP "Created" response.
func NewUssdNotifierUSSDResponseCreated(body *UssdNotifierResponseBody) *africastalkingviews.USSDResponseView {
	v := &africastalkingviews.USSDResponseView{
		Response: body.Response,
	}

	return v
}

// NewValidationNotifierAirtimeValidationResponseCreated builds a
// "africastalking" service "validation_notifier" endpoint result from a HTTP
// "Created" response.
func NewValidationNotifierAirtimeValidationResponseCreated(body *ValidationNotifierResponseBody) *africastalking.AirtimeValidationResponse {
	v := &africastalking.AirtimeValidationResponse{
		Status: body.Status,
	}

	return v
}

// NewB2cValidationNotifierB2CValidationNotificationResponseCreated builds a
// "africastalking" service "b2c_validation_notifier" endpoint result from a
// HTTP "Created" response.
func NewB2cValidationNotifierB2CValidationNotificationResponseCreated(body *B2cValidationNotifierResponseBody) *africastalking.B2CValidationNotificationResponse {
	v := &africastalking.B2CValidationNotificationResponse{
		Status: body.Status,
	}

	return v
}

// ValidateValidationNotifierResponseBody runs the validations defined on
// validation_notifier_response_body
func ValidateValidationNotifierResponseBody(body *ValidationNotifierResponseBody) (err error) {
	if body.Status != nil {
		if !(*body.Status == "Validated" || *body.Status == "Failed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"Validated", "Failed"}))
		}
	}
	return
}

// ValidateB2cValidationNotifierResponseBody runs the validations defined on
// b2c_validation_notifier_response_body
func ValidateB2cValidationNotifierResponseBody(body *B2cValidationNotifierResponseBody) (err error) {
	if body.Status != nil {
		if !(*body.Status == "Validated" || *body.Status == "Failed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"Validated", "Failed"}))
		}
	}
	return
}
