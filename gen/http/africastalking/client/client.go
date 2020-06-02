// Code generated by goa v3.1.3, DO NOT EDIT.
//
// africastalking client HTTP transport
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the africastalking service endpoint HTTP clients.
type Client struct {
	// SmsDeliveryReport Doer is the HTTP client used to make requests to the
	// sms_delivery_report endpoint.
	SmsDeliveryReportDoer goahttp.Doer

	// SmsIncomingMessage Doer is the HTTP client used to make requests to the
	// sms_incoming_message endpoint.
	SmsIncomingMessageDoer goahttp.Doer

	// SmsBulkOptout Doer is the HTTP client used to make requests to the
	// sms_bulk_optout endpoint.
	SmsBulkOptoutDoer goahttp.Doer

	// SmsSubscription Doer is the HTTP client used to make requests to the
	// sms_subscription endpoint.
	SmsSubscriptionDoer goahttp.Doer

	// VoiceNotification Doer is the HTTP client used to make requests to the
	// voice_notification endpoint.
	VoiceNotificationDoer goahttp.Doer

	// TransferEvent Doer is the HTTP client used to make requests to the
	// transfer_event endpoint.
	TransferEventDoer goahttp.Doer

	// UssdNotifier Doer is the HTTP client used to make requests to the
	// ussd_notifier endpoint.
	UssdNotifierDoer goahttp.Doer

	// ValidationNotifier Doer is the HTTP client used to make requests to the
	// validation_notifier endpoint.
	ValidationNotifierDoer goahttp.Doer

	// StatusNotifier Doer is the HTTP client used to make requests to the
	// status_notifier endpoint.
	StatusNotifierDoer goahttp.Doer

	// PaymentNotifier Doer is the HTTP client used to make requests to the
	// payment_notifier endpoint.
	PaymentNotifierDoer goahttp.Doer

	// C2bValidationNotifier Doer is the HTTP client used to make requests to the
	// c2b_validation_notifier endpoint.
	C2bValidationNotifierDoer goahttp.Doer

	// B2cValidationNotifier Doer is the HTTP client used to make requests to the
	// b2c_validation_notifier endpoint.
	B2cValidationNotifierDoer goahttp.Doer

	// IotNotifier Doer is the HTTP client used to make requests to the
	// iot_notifier endpoint.
	IotNotifierDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the africastalking service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		SmsDeliveryReportDoer:     doer,
		SmsIncomingMessageDoer:    doer,
		SmsBulkOptoutDoer:         doer,
		SmsSubscriptionDoer:       doer,
		VoiceNotificationDoer:     doer,
		TransferEventDoer:         doer,
		UssdNotifierDoer:          doer,
		ValidationNotifierDoer:    doer,
		StatusNotifierDoer:        doer,
		PaymentNotifierDoer:       doer,
		C2bValidationNotifierDoer: doer,
		B2cValidationNotifierDoer: doer,
		IotNotifierDoer:           doer,
		RestoreResponseBody:       restoreBody,
		scheme:                    scheme,
		host:                      host,
		decoder:                   dec,
		encoder:                   enc,
	}
}

// SmsDeliveryReport returns an endpoint that makes HTTP requests to the
// africastalking service sms_delivery_report server.
func (c *Client) SmsDeliveryReport() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeSmsDeliveryReportRequest(c.encoder)
		decodeResponse = DecodeSmsDeliveryReportResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSmsDeliveryReportRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SmsDeliveryReportDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "sms_delivery_report", err)
		}
		return decodeResponse(resp)
	}
}

// SmsIncomingMessage returns an endpoint that makes HTTP requests to the
// africastalking service sms_incoming_message server.
func (c *Client) SmsIncomingMessage() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeSmsIncomingMessageRequest(c.encoder)
		decodeResponse = DecodeSmsIncomingMessageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSmsIncomingMessageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SmsIncomingMessageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "sms_incoming_message", err)
		}
		return decodeResponse(resp)
	}
}

// SmsBulkOptout returns an endpoint that makes HTTP requests to the
// africastalking service sms_bulk_optout server.
func (c *Client) SmsBulkOptout() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeSmsBulkOptoutRequest(c.encoder)
		decodeResponse = DecodeSmsBulkOptoutResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSmsBulkOptoutRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SmsBulkOptoutDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "sms_bulk_optout", err)
		}
		return decodeResponse(resp)
	}
}

// SmsSubscription returns an endpoint that makes HTTP requests to the
// africastalking service sms_subscription server.
func (c *Client) SmsSubscription() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeSmsSubscriptionRequest(c.encoder)
		decodeResponse = DecodeSmsSubscriptionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSmsSubscriptionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SmsSubscriptionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "sms_subscription", err)
		}
		return decodeResponse(resp)
	}
}

// VoiceNotification returns an endpoint that makes HTTP requests to the
// africastalking service voice_notification server.
func (c *Client) VoiceNotification() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeVoiceNotificationRequest(c.encoder)
		decodeResponse = DecodeVoiceNotificationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildVoiceNotificationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VoiceNotificationDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "voice_notification", err)
		}
		return decodeResponse(resp)
	}
}

// TransferEvent returns an endpoint that makes HTTP requests to the
// africastalking service transfer_event server.
func (c *Client) TransferEvent() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeTransferEventRequest(c.encoder)
		decodeResponse = DecodeTransferEventResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildTransferEventRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TransferEventDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "transfer_event", err)
		}
		return decodeResponse(resp)
	}
}

// UssdNotifier returns an endpoint that makes HTTP requests to the
// africastalking service ussd_notifier server.
func (c *Client) UssdNotifier() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeUssdNotifierRequest(c.encoder)
		decodeResponse = DecodeUssdNotifierResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUssdNotifierRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UssdNotifierDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "ussd_notifier", err)
		}
		return decodeResponse(resp)
	}
}

// ValidationNotifier returns an endpoint that makes HTTP requests to the
// africastalking service validation_notifier server.
func (c *Client) ValidationNotifier() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeValidationNotifierRequest(c.encoder)
		decodeResponse = DecodeValidationNotifierResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildValidationNotifierRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ValidationNotifierDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "validation_notifier", err)
		}
		return decodeResponse(resp)
	}
}

// StatusNotifier returns an endpoint that makes HTTP requests to the
// africastalking service status_notifier server.
func (c *Client) StatusNotifier() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeStatusNotifierRequest(c.encoder)
		decodeResponse = DecodeStatusNotifierResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStatusNotifierRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StatusNotifierDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "status_notifier", err)
		}
		return decodeResponse(resp)
	}
}

// PaymentNotifier returns an endpoint that makes HTTP requests to the
// africastalking service payment_notifier server.
func (c *Client) PaymentNotifier() endpoint.Endpoint {
	var (
		encodeRequest  = EncodePaymentNotifierRequest(c.encoder)
		decodeResponse = DecodePaymentNotifierResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPaymentNotifierRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PaymentNotifierDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "payment_notifier", err)
		}
		return decodeResponse(resp)
	}
}

// C2bValidationNotifier returns an endpoint that makes HTTP requests to the
// africastalking service c2b_validation_notifier server.
func (c *Client) C2bValidationNotifier() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeC2bValidationNotifierRequest(c.encoder)
		decodeResponse = DecodeC2bValidationNotifierResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildC2bValidationNotifierRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.C2bValidationNotifierDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "c2b_validation_notifier", err)
		}
		return decodeResponse(resp)
	}
}

// B2cValidationNotifier returns an endpoint that makes HTTP requests to the
// africastalking service b2c_validation_notifier server.
func (c *Client) B2cValidationNotifier() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeB2cValidationNotifierRequest(c.encoder)
		decodeResponse = DecodeB2cValidationNotifierResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildB2cValidationNotifierRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.B2cValidationNotifierDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "b2c_validation_notifier", err)
		}
		return decodeResponse(resp)
	}
}

// IotNotifier returns an endpoint that makes HTTP requests to the
// africastalking service iot_notifier server.
func (c *Client) IotNotifier() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeIotNotifierRequest(c.encoder)
		decodeResponse = DecodeIotNotifierResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildIotNotifierRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.IotNotifierDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("africastalking", "iot_notifier", err)
		}
		return decodeResponse(resp)
	}
}