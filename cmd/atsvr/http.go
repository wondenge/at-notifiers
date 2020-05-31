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
	airtime "github.com/wondenge/at-notifiers/gen/airtime"
	health "github.com/wondenge/at-notifiers/gen/health"
	airtimekitsvr "github.com/wondenge/at-notifiers/gen/http/airtime/kitserver"
	airtimesvr "github.com/wondenge/at-notifiers/gen/http/airtime/server"
	healthkitsvr "github.com/wondenge/at-notifiers/gen/http/health/kitserver"
	healthsvr "github.com/wondenge/at-notifiers/gen/http/health/server"
	iotkitsvr "github.com/wondenge/at-notifiers/gen/http/iot/kitserver"
	iotsvr "github.com/wondenge/at-notifiers/gen/http/iot/server"
	paymentskitsvr "github.com/wondenge/at-notifiers/gen/http/payments/kitserver"
	paymentssvr "github.com/wondenge/at-notifiers/gen/http/payments/server"
	smskitsvr "github.com/wondenge/at-notifiers/gen/http/sms/kitserver"
	smssvr "github.com/wondenge/at-notifiers/gen/http/sms/server"
	swaggerkitsvr "github.com/wondenge/at-notifiers/gen/http/swagger/kitserver"
	swaggersvr "github.com/wondenge/at-notifiers/gen/http/swagger/server"
	ussdkitsvr "github.com/wondenge/at-notifiers/gen/http/ussd/kitserver"
	ussdsvr "github.com/wondenge/at-notifiers/gen/http/ussd/server"
	voicekitsvr "github.com/wondenge/at-notifiers/gen/http/voice/kitserver"
	voicesvr "github.com/wondenge/at-notifiers/gen/http/voice/server"
	iot "github.com/wondenge/at-notifiers/gen/iot"
	payments "github.com/wondenge/at-notifiers/gen/payments"
	sms "github.com/wondenge/at-notifiers/gen/sms"
	ussd "github.com/wondenge/at-notifiers/gen/ussd"
	voice "github.com/wondenge/at-notifiers/gen/voice"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, smsEndpoints *sms.Endpoints, voiceEndpoints *voice.Endpoints, ussdEndpoints *ussd.Endpoints, airtimeEndpoints *airtime.Endpoints, paymentsEndpoints *payments.Endpoints, iotEndpoints *iot.Endpoints, healthEndpoints *health.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

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
		smsDeliveryReportHandler       *kithttp.Server
		smsIncomingMessageHandler      *kithttp.Server
		smsBulkSMSOptOutHandler        *kithttp.Server
		smsSubNotifierHandler          *kithttp.Server
		smsServer                      *smssvr.Server
		voiceVoiceNotifierHandler      *kithttp.Server
		voiceTransferEventsHandler     *kithttp.Server
		voiceServer                    *voicesvr.Server
		ussdUssdNotifierHandler        *kithttp.Server
		ussdServer                     *ussdsvr.Server
		airtimeValidationHandler       *kithttp.Server
		airtimeStatusHandler           *kithttp.Server
		airtimeServer                  *airtimesvr.Server
		paymentsPaymentNotifierHandler *kithttp.Server
		paymentsC2bNotifierHandler     *kithttp.Server
		paymentsB2cNotifierHandler     *kithttp.Server
		paymentsServer                 *paymentssvr.Server
		iotIotNotifierHandler          *kithttp.Server
		iotServer                      *iotsvr.Server
		healthShowHandler              *kithttp.Server
		healthServer                   *healthsvr.Server
		swaggerServer                  *swaggersvr.Server
	)
	{
		eh := errorHandler(logger)
		smsDeliveryReportHandler = kithttp.NewServer(
			endpoint.Endpoint(smsEndpoints.DeliveryReport),
			smskitsvr.DecodeDeliveryReportRequest(mux, dec),
			smskitsvr.EncodeDeliveryReportResponse(enc),
		)
		smsIncomingMessageHandler = kithttp.NewServer(
			endpoint.Endpoint(smsEndpoints.IncomingMessage),
			smskitsvr.DecodeIncomingMessageRequest(mux, dec),
			smskitsvr.EncodeIncomingMessageResponse(enc),
		)
		smsBulkSMSOptOutHandler = kithttp.NewServer(
			endpoint.Endpoint(smsEndpoints.BulkSMSOptOut),
			smskitsvr.DecodeBulkSMSOptOutRequest(mux, dec),
			smskitsvr.EncodeBulkSMSOptOutResponse(enc),
		)
		smsSubNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(smsEndpoints.SubNotifier),
			smskitsvr.DecodeSubNotifierRequest(mux, dec),
			smskitsvr.EncodeSubNotifierResponse(enc),
		)
		smsServer = smssvr.New(smsEndpoints, mux, dec, enc, eh, nil)
		voiceVoiceNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(voiceEndpoints.VoiceNotifier),
			voicekitsvr.DecodeVoiceNotifierRequest(mux, dec),
			voicekitsvr.EncodeVoiceNotifierResponse(enc),
		)
		voiceTransferEventsHandler = kithttp.NewServer(
			endpoint.Endpoint(voiceEndpoints.TransferEvents),
			voicekitsvr.DecodeTransferEventsRequest(mux, dec),
			voicekitsvr.EncodeTransferEventsResponse(enc),
		)
		voiceServer = voicesvr.New(voiceEndpoints, mux, dec, enc, eh, nil)
		ussdUssdNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(ussdEndpoints.UssdNotifier),
			ussdkitsvr.DecodeUssdNotifierRequest(mux, dec),
			ussdkitsvr.EncodeUssdNotifierResponse(enc),
		)
		ussdServer = ussdsvr.New(ussdEndpoints, mux, dec, enc, eh, nil)
		airtimeValidationHandler = kithttp.NewServer(
			endpoint.Endpoint(airtimeEndpoints.Validation),
			airtimekitsvr.DecodeValidationRequest(mux, dec),
			airtimekitsvr.EncodeValidationResponse(enc),
		)
		airtimeStatusHandler = kithttp.NewServer(
			endpoint.Endpoint(airtimeEndpoints.Status),
			airtimekitsvr.DecodeStatusRequest(mux, dec),
			airtimekitsvr.EncodeStatusResponse(enc),
		)
		airtimeServer = airtimesvr.New(airtimeEndpoints, mux, dec, enc, eh, nil)
		paymentsPaymentNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(paymentsEndpoints.PaymentNotifier),
			paymentskitsvr.DecodePaymentNotifierRequest(mux, dec),
			paymentskitsvr.EncodePaymentNotifierResponse(enc),
		)
		paymentsC2bNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(paymentsEndpoints.C2bNotifier),
			paymentskitsvr.DecodeC2bNotifierRequest(mux, dec),
			paymentskitsvr.EncodeC2bNotifierResponse(enc),
		)
		paymentsB2cNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(paymentsEndpoints.B2cNotifier),
			paymentskitsvr.DecodeB2cNotifierRequest(mux, dec),
			paymentskitsvr.EncodeB2cNotifierResponse(enc),
		)
		paymentsServer = paymentssvr.New(paymentsEndpoints, mux, dec, enc, eh, nil)
		iotIotNotifierHandler = kithttp.NewServer(
			endpoint.Endpoint(iotEndpoints.IotNotifier),
			iotkitsvr.DecodeIotNotifierRequest(mux, dec),
			iotkitsvr.EncodeIotNotifierResponse(enc),
		)
		iotServer = iotsvr.New(iotEndpoints, mux, dec, enc, eh, nil)
		healthShowHandler = kithttp.NewServer(
			endpoint.Endpoint(healthEndpoints.Show),
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			healthkitsvr.EncodeShowResponse(enc),
		)
		healthServer = healthsvr.New(healthEndpoints, mux, dec, enc, eh, nil)
		swaggerServer = swaggersvr.New(nil, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	smskitsvr.MountDeliveryReportHandler(mux, smsDeliveryReportHandler)
	smskitsvr.MountIncomingMessageHandler(mux, smsIncomingMessageHandler)
	smskitsvr.MountBulkSMSOptOutHandler(mux, smsBulkSMSOptOutHandler)
	smskitsvr.MountSubNotifierHandler(mux, smsSubNotifierHandler)
	voicekitsvr.MountVoiceNotifierHandler(mux, voiceVoiceNotifierHandler)
	voicekitsvr.MountTransferEventsHandler(mux, voiceTransferEventsHandler)
	ussdkitsvr.MountUssdNotifierHandler(mux, ussdUssdNotifierHandler)
	airtimekitsvr.MountValidationHandler(mux, airtimeValidationHandler)
	airtimekitsvr.MountStatusHandler(mux, airtimeStatusHandler)
	paymentskitsvr.MountPaymentNotifierHandler(mux, paymentsPaymentNotifierHandler)
	paymentskitsvr.MountC2bNotifierHandler(mux, paymentsC2bNotifierHandler)
	paymentskitsvr.MountB2cNotifierHandler(mux, paymentsB2cNotifierHandler)
	iotkitsvr.MountIotNotifierHandler(mux, iotIotNotifierHandler)
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
	for _, m := range smsServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range voiceServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range ussdServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range airtimeServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range paymentsServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range iotServer.Mounts {
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
