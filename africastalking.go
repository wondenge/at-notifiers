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
func (s *africastalkingsrvc) DeliveryReportNotifier(ctx context.Context, p *africastalking.DeliveryReportPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.delivery_report_notifier"))
	return
}

// Adds new SMS Incoming Message to our callback URL and return its ID.
func (s *africastalkingsrvc) IncomingMessageNotifier(ctx context.Context, p *africastalking.IncomingMessagePayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.incoming_message_notifier"))
	return
}

// Adds new SMS Bulk OptOut to our callback URL and return its ID.
func (s *africastalkingsrvc) BulkOptOutNotifier(ctx context.Context, p *africastalking.BulkSMSOptOutPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.bulk_optOut_notifier"))
	return
}

// Adds new SMS subscription to our callback URL and return its ID.
func (s *africastalkingsrvc) SubNotifier(ctx context.Context, p *africastalking.SubNotificationPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.sub_notifier"))
	return
}

// Adds new Voice Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) VoiceNotifier(ctx context.Context, p *africastalking.VoiceNotificationPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.voice_notifier"))
	return
}

// Adds new Event Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) TransferEventNotifier(ctx context.Context, p *africastalking.TransferEventPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.transfer_event_notifier"))
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
func (s *africastalkingsrvc) StatusNotifier(ctx context.Context, p *africastalking.AirtimeStatusPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.status_notifier"))
	return
}

// Adds new Payment Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) PaymentNotifier(ctx context.Context, p *africastalking.PaymentNotificationPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.payment_notifier"))
	return
}

// Adds new C2B Validation Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) C2bValidationNotifier(ctx context.Context, p *africastalking.C2BValidationNotificationPayload) (res string, err error) {
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
func (s *africastalkingsrvc) IotNotifier(ctx context.Context, p *africastalking.IoTNotificationPayload) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("africastalking.iot_notifier"))
	return
}
