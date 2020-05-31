// Code generated by goa v3.1.3, DO NOT EDIT.
//
// iot client HTTP transport
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

// Client lists the iot service endpoint HTTP clients.
type Client struct {
	// IotNotifier Doer is the HTTP client used to make requests to the iotNotifier
	// endpoint.
	IotNotifierDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the iot service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		IotNotifierDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// IotNotifier returns an endpoint that makes HTTP requests to the iot service
// iotNotifier server.
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
			return nil, goahttp.ErrRequestError("iot", "iotNotifier", err)
		}
		return decodeResponse(resp)
	}
}
