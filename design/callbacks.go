package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

/*
	 SMS Callback Service.
 	 This service hosts the following notifications.
     1. Delivery Reports.
     2. Incoming Messages.
     3. Bulk SMS Opt Out.
     4. Subscription Notification.
*/
var _ = Service("africastalking", func() {
	HTTP(func() {
		Path("/callbacks/africastalking")
	})

	// Sent whenever an MSP confirms or rejects delivery of a message.
	Method("sms_delivery_report", func() {

		Description("Adds new SMS Delivery Report to our callback URL and return its ID.")
		Payload(DeliveryReport)
		Result(String)
		HTTP(func() {
			POST("/sms/deliveryreport")
			Response(StatusCreated)
		})
	})

	// Sent whenever a message is sent to any of your registered
	// shortcodes.
	Method("sms_incoming_message", func() {

		Description("Adds new SMS Incoming Message to our callback URL and return its ID.")
		Payload(IncomingMessage)
		Result(String)
		HTTP(func() {
			POST("/sms/incomingmessage")
			Response(StatusCreated)
		})
	})

	// Sent whenever a user opts out of receiving messages from your
	// alphanumeric sender ID
	Method("sms_bulk_optout", func() {
		Description("Adds new SMS Bulk OptOut to our callback URL and return its ID.")
		Payload(BulkSMSOptOut)
		Result(String)
		HTTP(func() {
			POST("/sms/bulksmsoptout")
			Response(StatusCreated)
		})
	})

	// Sent whenever someone subscribes or unsubscribes from any of
	// your premium SMS products.
	Method("sms_subscription", func() {
		Description("Adds new SMS subscription to our callback URL and return its ID.")
		Payload(SubscriptionNotification)
		Result(String)
		HTTP(func() {
			POST("/sms/subscription")
			Response(StatusCreated)
		})
	})

	// Voice notifications sent from Africa'sTalking gateway.
	Method("voice_notification", func() {
		Description("Adds new Voice Notification to our callback URL and return its ID.")
		Payload(VoiceNotification)
		Result(String)
		HTTP(func() {
			POST("/voice/notifications")
			Response(StatusCreated)
		})
	})

	// Sent from AT after call transfer initiated.
	Method("transfer_event", func() {
		Description("Adds new Event Notification to our callback URL and return its ID.")
		Payload(CallTransferEvent)
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
		Payload(AirtimeStatus)
		Result(String)
		HTTP(func() {
			POST("/airtime/status")
			Response(StatusCreated)
		})
	})

	// Payment Notifications
	Method("payment_notifier", func() {
		Description("Adds new Payment Notification to our callback URL and return its ID.")
		Payload(PaymentNotification)
		Result(String)
		HTTP(func() {
			POST("/payments/events")
			Response(StatusCreated)
		})
	})

	// C2B Validation Notifications
	Method("c2b_validation_notifier", func() {
		Description("Adds new C2B Validation Notification to our callback URL and return its ID.")
		Payload(C2BValidationNotification)
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
		Payload(IoTNotification)
		Result(String)
		HTTP(func() {
			POST("/iot/events")
			Response(StatusCreated)
		})
	})
})
