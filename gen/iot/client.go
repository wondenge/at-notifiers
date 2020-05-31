// Code generated by goa v3.1.3, DO NOT EDIT.
//
// iot client
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package iot

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "iot" service client.
type Client struct {
	IotNotifierEndpoint endpoint.Endpoint
}

// NewClient initializes a "iot" service client given the endpoints.
func NewClient(iotNotifier endpoint.Endpoint) *Client {
	return &Client{
		IotNotifierEndpoint: iotNotifier,
	}
}

// IotNotifier calls the "iotNotifier" endpoint of the "iot" service.
func (c *Client) IotNotifier(ctx context.Context, p *IoTNotification) (res string, err error) {
	var ires interface{}
	ires, err = c.IotNotifierEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
