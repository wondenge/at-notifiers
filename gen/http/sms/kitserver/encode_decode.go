// Code generated by goa v3.1.3, DO NOT EDIT.
//
// sms go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/at-notifiers/gen/http/sms/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeDeliveryReportResponse returns a go-kit EncodeResponseFunc suitable
// for encoding sms deliveryReport responses.
func EncodeDeliveryReportResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeDeliveryReportResponse(encoder)
}

// DecodeDeliveryReportRequest returns a go-kit DecodeRequestFunc suitable for
// decoding sms deliveryReport requests.
func DecodeDeliveryReportRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeDeliveryReportRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeIncomingMessageResponse returns a go-kit EncodeResponseFunc suitable
// for encoding sms incomingMessage responses.
func EncodeIncomingMessageResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeIncomingMessageResponse(encoder)
}

// DecodeIncomingMessageRequest returns a go-kit DecodeRequestFunc suitable for
// decoding sms incomingMessage requests.
func DecodeIncomingMessageRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeIncomingMessageRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeBulkSMSOptOutResponse returns a go-kit EncodeResponseFunc suitable for
// encoding sms bulkSMSOptOut responses.
func EncodeBulkSMSOptOutResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeBulkSMSOptOutResponse(encoder)
}

// DecodeBulkSMSOptOutRequest returns a go-kit DecodeRequestFunc suitable for
// decoding sms bulkSMSOptOut requests.
func DecodeBulkSMSOptOutRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeBulkSMSOptOutRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeSubNotifierResponse returns a go-kit EncodeResponseFunc suitable for
// encoding sms subNotifier responses.
func EncodeSubNotifierResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSubNotifierResponse(encoder)
}

// DecodeSubNotifierRequest returns a go-kit DecodeRequestFunc suitable for
// decoding sms subNotifier requests.
func DecodeSubNotifierRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeSubNotifierRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}