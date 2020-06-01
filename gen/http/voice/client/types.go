// Code generated by goa v3.1.3, DO NOT EDIT.
//
// voice HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

import (
	voice "github.com/wondenge/at-notifiers/gen/voice"
)

// VoiceNotifierRequestBody is the type of the "voice" service "voiceNotifier"
// endpoint HTTP request body.
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

// TransferEventsRequestBody is the type of the "voice" service
// "transferEvents" endpoint HTTP request body.
type TransferEventsRequestBody struct {
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

// NewVoiceNotifierRequestBody builds the HTTP request body from the payload of
// the "voiceNotifier" endpoint of the "voice" service.
func NewVoiceNotifierRequestBody(p *voice.VoiceNotification) *VoiceNotifierRequestBody {
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

// NewTransferEventsRequestBody builds the HTTP request body from the payload
// of the "transferEvents" endpoint of the "voice" service.
func NewTransferEventsRequestBody(p *voice.CallTransferEvent) *TransferEventsRequestBody {
	body := &TransferEventsRequestBody{
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