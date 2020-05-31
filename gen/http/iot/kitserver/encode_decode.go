// Code generated by goa v3.1.3, DO NOT EDIT.
//
// iot go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/at-notifiers/gen/http/iot/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeIotNotifierResponse returns a go-kit EncodeResponseFunc suitable for
// encoding iot iotNotifier responses.
func EncodeIotNotifierResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeIotNotifierResponse(encoder)
}

// DecodeIotNotifierRequest returns a go-kit DecodeRequestFunc suitable for
// decoding iot iotNotifier requests.
func DecodeIotNotifierRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeIotNotifierRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}
