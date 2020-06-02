package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("africastalking", func() {
	HTTP(func() {
		Path("/callbacks/africastalking")
	})

	// Sent whenever an MSP confirms or rejects delivery of a message.
	Method("delivery_report_notifier", func() {

		Description("Adds new SMS Delivery Report to our callback URL and return its ID.")
		Payload(DeliveryReportPayload)
		Result(String)
		HTTP(func() {
			POST("/sms/deliveryreport")
			Response(StatusCreated)
		})
	})

	// Sent whenever a message is sent to any of your registered
	// shortcodes.
	Method("incoming_message_notifier", func() {

		Description("Adds new SMS Incoming Message to our callback URL and return its ID.")
		Payload(IncomingMessagePayload)
		Result(String)
		HTTP(func() {
			POST("/sms/incomingmessage")
			Response(StatusCreated)
		})
	})

	// Sent whenever a user opts out of receiving messages from your
	// alphanumeric sender ID
	Method("bulk_optOut_notifier", func() {
		Description("Adds new SMS Bulk OptOut to our callback URL and return its ID.")
		Payload(BulkSMSOptOutPayload)
		Result(String)
		HTTP(func() {
			POST("/sms/bulksmsoptout")
			Response(StatusCreated)
		})
	})

	// Sent whenever someone subscribes or unsubscribes from any of
	// your premium SMS products.
	Method("sub_notifier", func() {
		Description("Adds new SMS subscription to our callback URL and return its ID.")
		Payload(SubNotificationPayload)
		Result(String)
		HTTP(func() {
			POST("/sms/subscription")
			Response(StatusCreated)
		})
	})

	// Voice notifications sent from Africa'sTalking gateway.
	Method("voice_notifier", func() {
		Description("Adds new Voice Notification to our callback URL and return its ID.")
		Payload(VoiceNotificationPayload)
		Result(String)
		HTTP(func() {
			POST("/voice/notifications")
			Response(StatusCreated)
		})
	})

	// Sent from AT after call transfer initiated.
	Method("transfer_event_notifier", func() {
		Description("Adds new Event Notification to our callback URL and return its ID.")
		Payload(TransferEventPayload)
		Result(String)
		HTTP(func() {
			POST("/voice/transferevents")
			Response(StatusCreated)
		})
	})

	// USSD Notifications sent from Africa'sTalking gateway.
	Method("ussd_notifier", func() {
		Description("Adds new USSD Notification to our callback URL and return its ID.")
		Payload(USSDPayload)
		Result(USSDResponse)
		HTTP(func() {
			POST("/ussd/sessions")
			Response(StatusCreated)
		})
	})

	Method("validation_notifier", func() {
		Description("Adds new Airtime Validation Notification to our callback URL and return its ID.")
		Payload(AirtimeValidationPayload)
		Result(AirtimeValidationResponse)
		HTTP(func() {
			POST("/airtime/validation")
			Response(StatusCreated)
		})
	})

	Method("status_notifier", func() {
		Description("Adds new Airtime Status Notification to our callback URL and return its ID.")
		Payload(AirtimeStatusPayload)
		Result(String)
		HTTP(func() {
			POST("/airtime/status")
			Response(StatusCreated)
		})
	})

	// Payment Notifications
	Method("payment_notifier", func() {
		Description("Adds new Payment Notification to our callback URL and return its ID.")
		Payload(PaymentNotificationPayload)
		Result(String)
		HTTP(func() {
			POST("/payments/events")
			Response(StatusCreated)
		})
	})

	// C2B Validation Notifications
	Method("c2b_validation_notifier", func() {
		Description("Adds new C2B Validation Notification to our callback URL and return its ID.")
		Payload(C2BValidationNotificationPayload)
		Result(String)
		HTTP(func() {
			POST("/payments/c2b/validation")
			Response(StatusCreated)
		})
	})

	// B2C Validation Notifications.
	Method("b2c_validation_notifier", func() {
		Description("Adds new B2C Validation Notification to our callback URL and return its ID.")
		Payload(B2CValidationNotificationPayload)
		Result(B2CValidationNotificationResponse)
		HTTP(func() {
			POST("/payments/b2c/validation")
			Response(StatusCreated)
		})
	})

	// IoT Notifications
	Method("iot_notifier", func() {
		Description("Adds new IoT Notification to our callback URL and return its ID.")
		Payload(IoTNotificationPayload)
		Result(String)
		HTTP(func() {
			POST("/iot/events")
			Response(StatusCreated)
		})
	})
})
