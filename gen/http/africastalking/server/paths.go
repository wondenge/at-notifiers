// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the africastalking service.
//
// Command:
// $ goa gen github.com/wondenge/at-notifiers/design

package server

// DeliveryReportNotifierAfricastalkingPath returns the URL path to the africastalking service delivery_report_notifier HTTP endpoint.
func DeliveryReportNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/sms/deliveryreport"
}

// IncomingMessageNotifierAfricastalkingPath returns the URL path to the africastalking service incoming_message_notifier HTTP endpoint.
func IncomingMessageNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/sms/incomingmessage"
}

// BulkOptOutNotifierAfricastalkingPath returns the URL path to the africastalking service bulk_optOut_notifier HTTP endpoint.
func BulkOptOutNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/sms/bulksmsoptout"
}

// SubNotifierAfricastalkingPath returns the URL path to the africastalking service sub_notifier HTTP endpoint.
func SubNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/sms/subscription"
}

// VoiceNotifierAfricastalkingPath returns the URL path to the africastalking service voice_notifier HTTP endpoint.
func VoiceNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/voice/notifications"
}

// TransferEventNotifierAfricastalkingPath returns the URL path to the africastalking service transfer_event_notifier HTTP endpoint.
func TransferEventNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/voice/transferevents"
}

// UssdNotifierAfricastalkingPath returns the URL path to the africastalking service ussd_notifier HTTP endpoint.
func UssdNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/ussd/sessions"
}

// ValidationNotifierAfricastalkingPath returns the URL path to the africastalking service validation_notifier HTTP endpoint.
func ValidationNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/airtime/validation"
}

// StatusNotifierAfricastalkingPath returns the URL path to the africastalking service status_notifier HTTP endpoint.
func StatusNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/airtime/status"
}

// PaymentNotifierAfricastalkingPath returns the URL path to the africastalking service payment_notifier HTTP endpoint.
func PaymentNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/payments/events"
}

// C2bValidationNotifierAfricastalkingPath returns the URL path to the africastalking service c2b_validation_notifier HTTP endpoint.
func C2bValidationNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/payments/c2b/validation"
}

// B2cValidationNotifierAfricastalkingPath returns the URL path to the africastalking service b2c_validation_notifier HTTP endpoint.
func B2cValidationNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/payments/b2c/validation"
}

// IotNotifierAfricastalkingPath returns the URL path to the africastalking service iot_notifier HTTP endpoint.
func IotNotifierAfricastalkingPath() string {
	return "/callbacks/africastalking/iot/events"
}
