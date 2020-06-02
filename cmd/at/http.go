package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	africastalking "github.com/wondenge/at-notifiers/gen/africastalking"
	health "github.com/wondenge/at-notifiers/gen/health"
	africastalkingkitsvr "github.com/wondenge/at-notifiers/gen/http/africastalking/kitserver"
	africastalkingsvr "github.com/wondenge/at-notifiers/gen/http/africastalking/server"
	healthkitsvr "github.com/wondenge/at-notifiers/gen/http/health/kitserver"
	healthsvr "github.com/wondenge/at-notifiers/gen/http/health/server"
	swaggerkitsvr "github.com/wondenge/at-notifiers/gen/http/swagger/kitserver"
	swaggersvr "github.com/wondenge/at-notifiers/gen/http/swagger/server"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, africastalkingEndpoints *africastalking.Endpoints, healthEndpoints *health.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		africastalkingSmsDeliveryReportHandler     *kithttp.Server
		africastalkingSmsIncomingMessageHandler    *kithttp.Server
		africastalkingSmsBulkOptoutHandler         *kithttp.Server
		africastalkingSmsSubscriptionHandler       *kithttp.Server
		africastalkingVoiceNotificationHandler     *kithttp.Server
		africastalkingTransferEventHandler         *kithttp.Server
		africastalkingUssdNotifierHandler          *kithttp.Server
		africastalkingValidationNotifierHandler    *kithttp.Server
		africastalkingStatusNotifierHandler        *kithttp.Server
		africastalkingPaymentNotifierHandler       *kithttp.Server
		africastalkingC2bValidationNotifierHandler *kithttp.Server
		africastalkingB2cValidationNotifierHandler *kithttp.Server
		africastalkingIotNotifierHandler           *kithttp.Server
		africastalkingServer                       *africastalkingsvr.Server
		healthShowHandler                          *kithttp.Server
		healthServer                               *healthsvr.Server
		swaggerServer                              *swaggersvr.Server
	)
	{
		eh := errorHandler(logger)
		africastalkingSmsDeliveryReportHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.SmsDeliveryReport),
			africastalkingkitsvr.DecodeSmsDeliveryReportRequest(mux, dec),
			africastalkingkitsvr.EncodeSmsDeliveryReportResponse(enc),
		)
		africastalkingSmsIncomingMessageHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.SmsIncomingMessage),
			africastalkingkitsvr.DecodeSmsIncomingMessageRequest(mux, dec),
			africastalkingkitsvr.EncodeSmsIncomingMessageResponse(enc),
		)
		africastalkingSmsBulkOptoutHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.SmsBulkOptout),
			africastalkingkitsvr.DecodeSmsBulkOptoutRequest(mux, dec),
			africastalkingkitsvr.EncodeSmsBulkOptoutResponse(enc),
		)
		africastalkingSmsSubscriptionHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.SmsSubscription),
			africastalkingkitsvr.DecodeSmsSubscriptionRequest(mux, dec),
			africastalkingkitsvr.EncodeSmsSubscriptionResponse(enc),
		)
		africastalkingVoiceNotificationHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.VoiceNotification),
			africastalkingkitsvr.DecodeVoiceNotificationRequest(mux, dec),
			africastalkingkitsvr.EncodeVoiceNotificationResponse(enc),
		)
		africastalkingTransferEventHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.TransferEvent),
			africastalkingkitsvr.DecodeTransferEventRequest(mux, dec),
			africastalkingkitsvr.EncodeTransferEventResponse(enc),
		)
		africastalkingUssdNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.UssdNotifier),
			africastalkingkitsvr.DecodeUssdNotifierRequest(mux, dec),
			africastalkingkitsvr.EncodeUssdNotifierResponse(enc),
		)
		africastalkingValidationNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.ValidationNotifier),
			africastalkingkitsvr.DecodeValidationNotifierRequest(mux, dec),
			africastalkingkitsvr.EncodeValidationNotifierResponse(enc),
		)
		africastalkingStatusNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.StatusNotifier),
			africastalkingkitsvr.DecodeStatusNotifierRequest(mux, dec),
			africastalkingkitsvr.EncodeStatusNotifierResponse(enc),
		)
		africastalkingPaymentNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.PaymentNotifier),
			africastalkingkitsvr.DecodePaymentNotifierRequest(mux, dec),
			africastalkingkitsvr.EncodePaymentNotifierResponse(enc),
		)
		africastalkingC2bValidationNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.C2bValidationNotifier),
			africastalkingkitsvr.DecodeC2bValidationNotifierRequest(mux, dec),
			africastalkingkitsvr.EncodeC2bValidationNotifierResponse(enc),
		)
		africastalkingB2cValidationNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.B2cValidationNotifier),
			africastalkingkitsvr.DecodeB2cValidationNotifierRequest(mux, dec),
			africastalkingkitsvr.EncodeB2cValidationNotifierResponse(enc),
		)
		africastalkingIotNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(africastalkingEndpoints.IotNotifier),
			africastalkingkitsvr.DecodeIotNotifierRequest(mux, dec),
			africastalkingkitsvr.EncodeIotNotifierResponse(enc),
		)
		africastalkingServer = africastalkingsvr.New(africastalkingEndpoints, mux, dec, enc, eh, nil)
		healthShowHandler = kithttp.NewServer(
			endpoint.Endpoint(healthEndpoints.Show),
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			healthkitsvr.EncodeShowResponse(enc),
		)
		healthServer = healthsvr.New(healthEndpoints, mux, dec, enc, eh, nil)
		swaggerServer = swaggersvr.New(nil, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	africastalkingkitsvr.MountSmsDeliveryReportHandler(mux, africastalkingSmsDeliveryReportHandler)
	africastalkingkitsvr.MountSmsIncomingMessageHandler(mux, africastalkingSmsIncomingMessageHandler)
	africastalkingkitsvr.MountSmsBulkOptoutHandler(mux, africastalkingSmsBulkOptoutHandler)
	africastalkingkitsvr.MountSmsSubscriptionHandler(mux, africastalkingSmsSubscriptionHandler)
	africastalkingkitsvr.MountVoiceNotificationHandler(mux, africastalkingVoiceNotificationHandler)
	africastalkingkitsvr.MountTransferEventHandler(mux, africastalkingTransferEventHandler)
	africastalkingkitsvr.MountUssdNotifierHandler(mux, africastalkingUssdNotifierHandler)
	africastalkingkitsvr.MountValidationNotifierHandler(mux, africastalkingValidationNotifierHandler)
	africastalkingkitsvr.MountStatusNotifierHandler(mux, africastalkingStatusNotifierHandler)
	africastalkingkitsvr.MountPaymentNotifierHandler(mux, africastalkingPaymentNotifierHandler)
	africastalkingkitsvr.MountC2bValidationNotifierHandler(mux, africastalkingC2bValidationNotifierHandler)
	africastalkingkitsvr.MountB2cValidationNotifierHandler(mux, africastalkingB2cValidationNotifierHandler)
	africastalkingkitsvr.MountIotNotifierHandler(mux, africastalkingIotNotifierHandler)
	healthkitsvr.MountShowHandler(mux, healthShowHandler)
	swaggerkitsvr.MountGenHTTPOpenapiJSON(mux)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(logger)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range africastalkingServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range healthServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range swaggerServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host))
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host))

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
