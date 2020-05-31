package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/go-kit/kit/log"
	notifier "github.com/wondenge/at-notifiers"
	airtime "github.com/wondenge/at-notifiers/gen/airtime"
	health "github.com/wondenge/at-notifiers/gen/health"
	iot "github.com/wondenge/at-notifiers/gen/iot"
	payments "github.com/wondenge/at-notifiers/gen/payments"
	sms "github.com/wondenge/at-notifiers/gen/sms"
	ussd "github.com/wondenge/at-notifiers/gen/ussd"
	voice "github.com/wondenge/at-notifiers/gen/voice"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "local", "Server host (valid values: local, docker)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup gokit logger.
	var (
		logger log.Logger
	)
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Initialize the services.
	var (
		smsSvc      sms.Service
		voiceSvc    voice.Service
		ussdSvc     ussd.Service
		airtimeSvc  airtime.Service
		paymentsSvc payments.Service
		iotSvc      iot.Service
		healthSvc   health.Service
	)
	{
		smsSvc = notifier.NewSms(logger)
		voiceSvc = notifier.NewVoice(logger)
		ussdSvc = notifier.NewUssd(logger)
		airtimeSvc = notifier.NewAirtime(logger)
		paymentsSvc = notifier.NewPayments(logger)
		iotSvc = notifier.NewIot(logger)
		healthSvc = notifier.NewHealth(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		smsEndpoints      *sms.Endpoints
		voiceEndpoints    *voice.Endpoints
		ussdEndpoints     *ussd.Endpoints
		airtimeEndpoints  *airtime.Endpoints
		paymentsEndpoints *payments.Endpoints
		iotEndpoints      *iot.Endpoints
		healthEndpoints   *health.Endpoints
	)
	{
		smsEndpoints = sms.NewEndpoints(smsSvc)
		voiceEndpoints = voice.NewEndpoints(voiceSvc)
		ussdEndpoints = ussd.NewEndpoints(ussdSvc)
		airtimeEndpoints = airtime.NewEndpoints(airtimeSvc)
		paymentsEndpoints = payments.NewEndpoints(paymentsSvc)
		iotEndpoints = iot.NewEndpoints(iotSvc)
		healthEndpoints = health.NewEndpoints(healthSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "local":
		{
			addr := "http://localhost:3000"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, smsEndpoints, voiceEndpoints, ussdEndpoints, airtimeEndpoints, paymentsEndpoints, iotEndpoints, healthEndpoints, &wg, errc, logger, *dbgF)
		}

	case "docker":
		{
			addr := "http://0.0.0.0:8000"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, smsEndpoints, voiceEndpoints, ussdEndpoints, airtimeEndpoints, paymentsEndpoints, iotEndpoints, healthEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: local|docker)\n", *hostF)
	}

	// Wait for signal.
	logger.Log("info", fmt.Sprintf("exiting (%v)", <-errc))

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Log("info", fmt.Sprintf("exited"))
}
