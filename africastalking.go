package listener

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	africastalking "github.com/wondenge/at-notifiers/gen/africastalking"
)

// africastalking service example implementation.
// The example methods log the requests and return zero values.
type africastalkingsrvc struct {
	logger log.Logger
}

// NewAfricastalking returns the africastalking service implementation.
func NewAfricastalking(logger log.Logger) africastalking.Service {
	return &africastalkingsrvc{logger}
}

// Adds new SMS Delivery Report to our callback URL and return its ID.
func (s *africastalkingsrvc) SmsDeliveryReport(ctx context.Context, p *africastalking.DeliveryReport) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.sms_delivery_report"))
	return
}

// Adds new SMS Incoming Message to our callback URL and return its ID.
func (s *africastalkingsrvc) SmsIncomingMessage(ctx context.Context, p *africastalking.IncomingMessage) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.sms_incoming_message"))
	return
}

// Adds new SMS Bulk OptOut to our callback URL and return its ID.
func (s *africastalkingsrvc) SmsBulkOptout(ctx context.Context, p *africastalking.BulkSMSOptOut) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.sms_bulk_optout"))
	return
}

// Adds new SMS subscription to our callback URL and return its ID.
func (s *africastalkingsrvc) SmsSubscription(ctx context.Context, p *africastalking.SubscriptionNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.sms_subscription"))
	return
}

// Adds new Voice Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) VoiceNotification(ctx context.Context, p *africastalking.VoiceNotification1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.voice_notification"))
	return
}

// Adds new Event Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) TransferEvent(ctx context.Context, p *africastalking.CallTransferEvent) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.transfer_event"))
	return
}

// Adds new USSD Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) UssdNotifier(ctx context.Context, p *africastalking.USSDPayload) (res *africastalking.USSDResponse, err error) {
	res = &africastalking.USSDResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.ussd_notifier"))
	return
}

// Adds new Airtime Validation Notification to our callback URL and return its
// ID.
func (s *africastalkingsrvc) ValidationNotifier(ctx context.Context, p *africastalking.AirtimeValidationPayload) (res *africastalking.AirtimeValidationResponse, err error) {
	res = &africastalking.AirtimeValidationResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.validation_notifier"))
	return
}

// Adds new Airtime Status Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) StatusNotifier(ctx context.Context, p *africastalking.AirtimeStatus) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.status_notifier"))
	return
}

// Adds new Payment Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) PaymentNotifier(ctx context.Context, p *africastalking.PaymentNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.payment_notifier"))
	return
}

// Adds new C2B Validation Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) C2bValidationNotifier(ctx context.Context, p *africastalking.C2BValidationNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.c2b_validation_notifier"))
	return
}

// Adds new B2C Validation Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) B2cValidationNotifier(ctx context.Context, p *africastalking.B2CValidationNotificationPayload) (res *africastalking.B2CValidationNotificationResponse, err error) {
	res = &africastalking.B2CValidationNotificationResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.b2c_validation_notifier"))
	return
}

// Adds new IoT Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) IotNotifier(ctx context.Context, p *africastalking.IoTNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.iot_notifier"))
	return
}
