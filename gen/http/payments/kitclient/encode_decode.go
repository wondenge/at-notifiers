// Code generated by goa v3.1.3, DO NOT EDIT.
//
// payments go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/at-notifiers/gen/http/payments/client"
	goahttp "goa.design/goa/v3/http"
)

// EncodePaymentNotifierRequest returns a go-kit EncodeRequestFunc suitable for
// encoding payments paymentNotifier requests.
func EncodePaymentNotifierRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodePaymentNotifierRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodePaymentNotifierResponse returns a go-kit DecodeResponseFunc suitable
// for decoding payments paymentNotifier responses.
func DecodePaymentNotifierResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodePaymentNotifierResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeC2bNotifierRequest returns a go-kit EncodeRequestFunc suitable for
// encoding payments c2bNotifier requests.
func EncodeC2bNotifierRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeC2bNotifierRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeC2bNotifierResponse returns a go-kit DecodeResponseFunc suitable for
// decoding payments c2bNotifier responses.
func DecodeC2bNotifierResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeC2bNotifierResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}

// EncodeB2cNotifierRequest returns a go-kit EncodeRequestFunc suitable for
// encoding payments b2cNotifier requests.
func EncodeB2cNotifierRequest(encoder func(*http.Request) goahttp.Encoder) kithttp.EncodeRequestFunc {
	enc := client.EncodeB2cNotifierRequest(encoder)
	return func(_ context.Context, r *http.Request, v interface{}) error {
		return enc(r, v)
	}
}

// DecodeB2cNotifierResponse returns a go-kit DecodeResponseFunc suitable for
// decoding payments b2cNotifier responses.
func DecodeB2cNotifierResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeB2cNotifierResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
