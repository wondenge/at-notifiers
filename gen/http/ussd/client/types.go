// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ussd HTTP client types
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

import (
	ussd "github.com/wondenge/at-notifiers/gen/ussd"
	ussdviews "github.com/wondenge/at-notifiers/gen/ussd/views"
)

// UssdNotifierRequestBody is the type of the "ussd" service "ussdNotifier"
// endpoint HTTP request body.
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

// UssdNotifierResponseBody is the type of the "ussd" service "ussdNotifier"
// endpoint HTTP response body.
type UssdNotifierResponseBody struct {
	// Plain text response back to AT gateway
	Response *string `form:"response,omitempty" json:"response,omitempty" xml:"response,omitempty"`
}

// NewUssdNotifierRequestBody builds the HTTP request body from the payload of
// the "ussdNotifier" endpoint of the "ussd" service.
func NewUssdNotifierRequestBody(p *ussd.USSDPayload) *UssdNotifierRequestBody {
	body := &UssdNotifierRequestBody{
		SessionID:   p.SessionID,
		PhoneNumber: p.PhoneNumber,
		NetworkCode: p.NetworkCode,
		ServiceCode: p.ServiceCode,
		Text:        p.Text,
	}
	return body
}

// NewUssdNotifierUSSDResponseCreated builds a "ussd" service "ussdNotifier"
// endpoint result from a HTTP "Created" response.
func NewUssdNotifierUSSDResponseCreated(body *UssdNotifierResponseBody) *ussdviews.USSDResponseView {
	v := &ussdviews.USSDResponseView{
		Response: body.Response,
	}

	return v
}
