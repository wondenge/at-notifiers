// Code generated by goa v3.1.3, DO NOT EDIT.
//
// payments go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/at-notifiers/gen/http/payments/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodePaymentNotifierResponse returns a go-kit EncodeResponseFunc suitable
// for encoding payments paymentNotifier responses.
func EncodePaymentNotifierResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePaymentNotifierResponse(encoder)
}

// DecodePaymentNotifierRequest returns a go-kit DecodeRequestFunc suitable for
// decoding payments paymentNotifier requests.
func DecodePaymentNotifierRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePaymentNotifierRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeC2bNotifierResponse returns a go-kit EncodeResponseFunc suitable for
// encoding payments c2bNotifier responses.
func EncodeC2bNotifierResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeC2bNotifierResponse(encoder)
}

// DecodeC2bNotifierRequest returns a go-kit DecodeRequestFunc suitable for
// decoding payments c2bNotifier requests.
func DecodeC2bNotifierRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeC2bNotifierRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeB2cNotifierResponse returns a go-kit EncodeResponseFunc suitable for
// encoding payments b2cNotifier responses.
func EncodeB2cNotifierResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeB2cNotifierResponse(encoder)
}

// DecodeB2cNotifierRequest returns a go-kit DecodeRequestFunc suitable for
// decoding payments b2cNotifier requests.
func DecodeB2cNotifierRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeB2cNotifierRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}