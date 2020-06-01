// Code generated by goa v3.1.3, DO NOT EDIT.
//
// payments go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountPaymentNotifierHandler configures the mux to serve the "payments"
// service "paymentNotifier" endpoint.
func MountPaymentNotifierHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/callbacks/africastalking/payments/events", f)
}

// MountC2bNotifierHandler configures the mux to serve the "payments" service
// "c2bNotifier" endpoint.
func MountC2bNotifierHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/callbacks/africastalking/payments/c2b/validation", f)
}

// MountB2cNotifierHandler configures the mux to serve the "payments" service
// "b2cNotifier" endpoint.
func MountB2cNotifierHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/callbacks/africastalking/payments/b2c/validation", f)
}