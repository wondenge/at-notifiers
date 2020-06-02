package listener

import (
	"context"
	"fmt"
	"github.com/boltdb/bolt"

	"github.com/go-kit/kit/log"
	africastalking "github.com/wondenge/at-notifiers/gen/africastalking"
)

// africastalking service example implementation.
// The example methods log the requests and return zero values.
type africastalkingsrvc struct {
	db     *Bolt
	logger log.Logger
}

// NewAfricastalking returns the africastalking service implementation.
func NewAfricastalking(db *bolt.DB, logger log.Logger) africastalking.Service {

	// Setup Database
	boltdb, err := NewBoltDB(db)
	if err != nil {
		return nil
	}

	// Build and return service implementation
	return &africastalkingsrvc{boltdb, logger}
}

// Adds new SMS Delivery Report to our callback URL and return its ID.
func (s *africastalkingsrvc) DeliveryReportNotifier(ctx context.Context, p *africastalking.DeliveryReportPayload) (res string, err error) {

	res, err = s.db.NewID("DeliveryReport")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.DeliveryReportPayload{
		ID:            p.ID,
		Status:        p.Status,
		PhoneNumber:   p.PhoneNumber,
		NetworkCode:   p.NetworkCode,
		FailureReason: p.NetworkCode,
		RetryCount:    p.RetryCount,
	}

	if err = s.db.Save("DeliveryReport", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.delivery_report_notifier"))

	return res, nil
}

// Adds new SMS Incoming Message to our callback URL and return its ID.
func (s *africastalkingsrvc) IncomingMessageNotifier(ctx context.Context, p *africastalking.IncomingMessagePayload) (res string, err error) {

	res, err = s.db.NewID("IncomingMessage")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.IncomingMessagePayload{
		Date:        p.Date,
		From:        p.From,
		ID:          p.ID,
		LinkID:      p.LinkID,
		Text:        p.Text,
		To:          p.To,
		NetworkCode: p.NetworkCode,
	}

	if err = s.db.Save("IncomingMessage", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.incoming_message_notifier"))

	return res, nil
}

// Adds new SMS Bulk OptOut to our callback URL and return its ID.
func (s *africastalkingsrvc) BulkOptOutNotifier(ctx context.Context, p *africastalking.BulkSMSOptOutPayload) (res string, err error) {

	res, err = s.db.NewID("BulkSMSOptOut")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.BulkSMSOptOutPayload{
		SenderID:    p.SenderID,
		PhoneNumber: p.PhoneNumber,
	}

	if err = s.db.Save("BulkSMSOptOut", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.bulk_optOut_notifier"))

	return res, nil
}

// Adds new SMS subscription to our callback URL and return its ID.
func (s *africastalkingsrvc) SubNotifier(ctx context.Context, p *africastalking.SubNotificationPayload) (res string, err error) {

	res, err = s.db.NewID("SubNotification")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.SubNotificationPayload{
		PhoneNumber: p.PhoneNumber,
		ShortCode:   p.ShortCode,
		Keyword:     p.Keyword,
		UpdateType:  p.UpdateType,
	}

	if err = s.db.Save("SubNotification", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.sub_notifier"))

	return res, nil
}

// Adds new Voice Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) VoiceNotifier(ctx context.Context, p *africastalking.VoiceNotificationPayload) (res string, err error) {

	res, err = s.db.NewID("VoiceNotification")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.VoiceNotificationPayload{
		IsActive:              p.IsActive,
		SessionID:             p.SessionID,
		Direction:             p.Direction,
		DestinationNumber:     p.DestinationNumber,
		CallerNumber:          p.CallerNumber,
		CallerCountryCode:     p.CallerCountryCode,
		CallStartTime:         p.CallStartTime,
		DtmfDigits:            p.DtmfDigits,
		RecordingURL:          p.RecordingURL,
		DurationInSeconds:     p.DurationInSeconds,
		CurrencyCode:          p.CurrencyCode,
		Amount:                p.Amount,
		CallSessionState:      p.CallSessionState,
		DialDestinationNumber: p.DialDestinationNumber,
		DialDurationInSeconds: p.DialDurationInSeconds,
		DialStartTime:         p.DialStartTime,
		HangupCause:           p.HangupCause,
	}

	if err = s.db.Save("VoiceNotification", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.voice_notifier"))

	return res, nil
}

// Adds new Event Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) TransferEventNotifier(ctx context.Context, p *africastalking.TransferEventPayload) (res string, err error) {

	res, err = s.db.NewID("TransferEvent")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.TransferEventPayload{
		CallSessionState:        p.CallSessionState,
		IsActive:                p.IsActive,
		Status:                  p.Status,
		CallTransferParam:       p.CallTransferParam,
		CallTransferredToNumber: p.CallTransferredToNumber,
		CallTransferState:       p.CallTransferState,
		CallTransferHangupCause: p.CallTransferHangupCause,
	}

	if err = s.db.Save("TransferEvent", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.transfer_event_notifier"))

	return res, nil
}

// Adds new USSD Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) UssdNotifier(ctx context.Context, p *africastalking.USSDPayload) (res *africastalking.USSDResponse, err error) {
	res = &africastalking.USSDResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.ussd_notifier"))
	return res, nil
}

// Adds new Airtime Validation Notification to our callback URL and return its
// ID.
func (s *africastalkingsrvc) ValidationNotifier(ctx context.Context, p *africastalking.AirtimeValidationPayload) (res *africastalking.AirtimeValidationResponse, err error) {
	res = &africastalking.AirtimeValidationResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.validation_notifier"))
	return res, nil
}

// Adds new Airtime Status Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) StatusNotifier(ctx context.Context, p *africastalking.AirtimeStatusPayload) (res string, err error) {

	res, err = s.db.NewID("AirtimeStatus")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.AirtimeStatusPayload{
		RequestID: p.RequestID,
		Status:    p.Status,
	}

	if err = s.db.Save("AirtimeStatus", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.status_notifier"))

	return res, nil
}

// Adds new Payment Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) PaymentNotifier(ctx context.Context, p *africastalking.PaymentNotificationPayload) (res string, err error) {

	res, err = s.db.NewID("PaymentNotification")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.PaymentNotificationPayload{
		TransactionID:    p.TransactionID,
		Category:         p.Category,
		Provider:         p.Provider,
		ProviderRefID:    p.ProviderRefID,
		ProviderChannel:  p.ProviderChannel,
		ClientAccount:    p.ClientAccount,
		ProductName:      p.ProductName,
		SourceType:       p.SourceType,
		Source:           p.Source,
		DestinationType:  p.DestinationType,
		Destination:      p.Destination,
		Value:            p.Value,
		TransactionFee:   p.TransactionFee,
		ProviderFee:      p.ProviderFee,
		Status:           p.Status,
		Description:      p.Description,
		RequestMetadata:  p.RequestMetadata,
		ProviderMetadata: p.ProviderMetadata,
		TransactionDate:  p.TransactionDate,
	}

	if err = s.db.Save("PaymentNotification", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.payment_notifier"))

	return res, nil
}

// Adds new C2B Validation Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) C2bValidationNotifier(ctx context.Context, p *africastalking.C2BValidationNotificationPayload) (res string, err error) {

	res, err = s.db.NewID("C2BValidationNotification")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.C2BValidationNotificationPayload{
		Provider:         p.Provider,
		ClientAccount:    p.ClientAccount,
		ProductName:      p.ProductName,
		PhoneNumber:      p.PhoneNumber,
		Value:            p.Value,
		ProviderMetadata: p.ProviderMetadata,
	}

	if err = s.db.Save("C2BValidationNotification", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.c2b_validation_notifier"))

	return res, nil
}

// Adds new B2C Validation Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) B2cValidationNotifier(ctx context.Context, p *africastalking.B2CValidationNotificationPayload) (res *africastalking.B2CValidationNotificationResponse, err error) {
	res = &africastalking.B2CValidationNotificationResponse{}
	s.logger.Log("info", fmt.Sprintf("africastalking.b2c_validation_notifier"))
	return res, nil
}

// Adds new IoT Notification to our callback URL and return its ID.
func (s *africastalkingsrvc) IotNotifier(ctx context.Context, p *africastalking.IoTNotificationPayload) (res string, err error) {

	res, err = s.db.NewID("IoTNotification")
	if err != nil {
		return "", err // internal error
	}

	sn := africastalking.IoTNotificationPayload{
		Payload: p.Payload,
		Topic:   p.Topic,
	}

	if err = s.db.Save("IoTNotification", res, &sn); err != nil {
		return "", err // internal error
	}

	s.logger.Log("info", fmt.Sprintf("africastalking.iot_notifier"))

	return res, nil
}
