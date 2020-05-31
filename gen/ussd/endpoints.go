// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ussd endpoints
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package ussd

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "ussd" service endpoints.
type Endpoints struct {
	UssdNotifier endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "ussd" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		UssdNotifier: NewUssdNotifierEndpoint(s),
	}
}

// Use applies the given middleware to all the "ussd" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.UssdNotifier = m(e.UssdNotifier)
}

// NewUssdNotifierEndpoint returns an endpoint function that calls the method
// "ussdNotifier" of service "ussd".
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
