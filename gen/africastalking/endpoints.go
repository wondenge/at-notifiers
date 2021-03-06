// Code generated by goa v3.1.3, DO NOT EDIT.
//
// africastalking endpoints
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package africastalking

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "africastalking" service endpoints.
type Endpoints struct {
	DeliveryReportNotifier  endpoint.Endpoint
	IncomingMessageNotifier endpoint.Endpoint
	BulkOptOutNotifier      endpoint.Endpoint
	SubNotifier             endpoint.Endpoint
	VoiceNotifier           endpoint.Endpoint
	TransferEventNotifier   endpoint.Endpoint
	UssdNotifier            endpoint.Endpoint
	ValidationNotifier      endpoint.Endpoint
	StatusNotifier          endpoint.Endpoint
	PaymentNotifier         endpoint.Endpoint
	C2bValidationNotifier   endpoint.Endpoint
	B2cValidationNotifier   endpoint.Endpoint
	IotNotifier             endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "africastalking" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		DeliveryReportNotifier:  NewDeliveryReportNotifierEndpoint(s),
		IncomingMessageNotifier: NewIncomingMessageNotifierEndpoint(s),
		BulkOptOutNotifier:      NewBulkOptOutNotifierEndpoint(s),
		SubNotifier:             NewSubNotifierEndpoint(s),
		VoiceNotifier:           NewVoiceNotifierEndpoint(s),
		TransferEventNotifier:   NewTransferEventNotifierEndpoint(s),
		UssdNotifier:            NewUssdNotifierEndpoint(s),
		ValidationNotifier:      NewValidationNotifierEndpoint(s),
		StatusNotifier:          NewStatusNotifierEndpoint(s),
		PaymentNotifier:         NewPaymentNotifierEndpoint(s),
		C2bValidationNotifier:   NewC2bValidationNotifierEndpoint(s),
		B2cValidationNotifier:   NewB2cValidationNotifierEndpoint(s),
		IotNotifier:             NewIotNotifierEndpoint(s),
	}
}

// Use applies the given middleware to all the "africastalking" service
// endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.DeliveryReportNotifier = m(e.DeliveryReportNotifier)
	e.IncomingMessageNotifier = m(e.IncomingMessageNotifier)
	e.BulkOptOutNotifier = m(e.BulkOptOutNotifier)
	e.SubNotifier = m(e.SubNotifier)
	e.VoiceNotifier = m(e.VoiceNotifier)
	e.TransferEventNotifier = m(e.TransferEventNotifier)
	e.UssdNotifier = m(e.UssdNotifier)
	e.ValidationNotifier = m(e.ValidationNotifier)
	e.StatusNotifier = m(e.StatusNotifier)
	e.PaymentNotifier = m(e.PaymentNotifier)
	e.C2bValidationNotifier = m(e.C2bValidationNotifier)
	e.B2cValidationNotifier = m(e.B2cValidationNotifier)
	e.IotNotifier = m(e.IotNotifier)
}

// NewDeliveryReportNotifierEndpoint returns an endpoint function that calls
// the method "delivery_report_notifier" of service "africastalking".
func NewDeliveryReportNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeliveryReportPayload)
		return s.DeliveryReportNotifier(ctx, p)
	}
}

// NewIncomingMessageNotifierEndpoint returns an endpoint function that calls
// the method "incoming_message_notifier" of service "africastalking".
func NewIncomingMessageNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IncomingMessagePayload)
		return s.IncomingMessageNotifier(ctx, p)
	}
}

// NewBulkOptOutNotifierEndpoint returns an endpoint function that calls the
// method "bulk_optOut_notifier" of service "africastalking".
func NewBulkOptOutNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*BulkSMSOptOutPayload)
		return s.BulkOptOutNotifier(ctx, p)
	}
}

// NewSubNotifierEndpoint returns an endpoint function that calls the method
// "sub_notifier" of service "africastalking".
func NewSubNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SubNotificationPayload)
		return s.SubNotifier(ctx, p)
	}
}

// NewVoiceNotifierEndpoint returns an endpoint function that calls the method
// "voice_notifier" of service "africastalking".
func NewVoiceNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*VoiceNotificationPayload)
		return s.VoiceNotifier(ctx, p)
	}
}

// NewTransferEventNotifierEndpoint returns an endpoint function that calls the
// method "transfer_event_notifier" of service "africastalking".
func NewTransferEventNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*TransferEventPayload)
		return s.TransferEventNotifier(ctx, p)
	}
}

// NewUssdNotifierEndpoint returns an endpoint function that calls the method
// "ussd_notifier" of service "africastalking".
func NewUssdNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*USSDPayload)
		res, err := s.UssdNotifier(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUSSDResponse(res, "default")
		return vres, nil
	}
}

// NewValidationNotifierEndpoint returns an endpoint function that calls the
// method "validation_notifier" of service "africastalking".
func NewValidationNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AirtimeValidationPayload)
		return s.ValidationNotifier(ctx, p)
	}
}

// NewStatusNotifierEndpoint returns an endpoint function that calls the method
// "status_notifier" of service "africastalking".
func NewStatusNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AirtimeStatusPayload)
		return s.StatusNotifier(ctx, p)
	}
}

// NewPaymentNotifierEndpoint returns an endpoint function that calls the
// method "payment_notifier" of service "africastalking".
func NewPaymentNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PaymentNotificationPayload)
		return s.PaymentNotifier(ctx, p)
	}
}

// NewC2bValidationNotifierEndpoint returns an endpoint function that calls the
// method "c2b_validation_notifier" of service "africastalking".
func NewC2bValidationNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*C2BValidationNotificationPayload)
		return s.C2bValidationNotifier(ctx, p)
	}
}

// NewB2cValidationNotifierEndpoint returns an endpoint function that calls the
// method "b2c_validation_notifier" of service "africastalking".
func NewB2cValidationNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*B2CValidationNotificationPayload)
		return s.B2cValidationNotifier(ctx, p)
	}
}

// NewIotNotifierEndpoint returns an endpoint function that calls the method
// "iot_notifier" of service "africastalking".
func NewIotNotifierEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*IoTNotificationPayload)
		return s.IotNotifier(ctx, p)
	}
}
