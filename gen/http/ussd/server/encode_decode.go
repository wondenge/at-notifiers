// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ussd HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package server

import (
	"context"
	"io"
	"net/http"

	ussdviews "github.com/wondenge/at-notifiers/gen/ussd/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUssdNotifierResponse returns an encoder for responses returned by the
// ussd ussdNotifier endpoint.
func EncodeUssdNotifierResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*ussdviews.USSDResponse)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/plain")
		enc := encoder(ctx, w)
		body := NewUssdNotifierResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeUssdNotifierRequest returns a decoder for requests sent to the ussd
// ussdNotifier endpoint.
func DecodeUssdNotifierRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UssdNotifierRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewUssdNotifierUSSDPayload(&body)

		return payload, nil
	}
}