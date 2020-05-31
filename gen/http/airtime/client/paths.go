// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the airtime service.
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package client

// ValidationAirtimePath returns the URL path to the airtime service validation HTTP endpoint.
func ValidationAirtimePath() string {
	return "/callbacks/africastalking/airtime/validation"
}

// StatusAirtimePath returns the URL path to the airtime service status HTTP endpoint.
func StatusAirtimePath() string {
	return "/callbacks/africastalking/airtime/status"
}
