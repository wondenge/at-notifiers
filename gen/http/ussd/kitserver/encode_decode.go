// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ussd go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/at-notifiers/gen/http/ussd/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeUssdNotifierResponse returns a go-kit EncodeResponseFunc suitable for
// encoding ussd ussdNotifier responses.
func EncodeUssdNotifierResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeUssdNotifierResponse(encoder)
}

// DecodeUssdNotifierRequest returns a go-kit DecodeRequestFunc suitable for
// decoding ussd ussdNotifier requests.
func DecodeUssdNotifierRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeUssdNotifierRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}