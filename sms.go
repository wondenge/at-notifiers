package notifier

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	sms "github.com/wondenge/at-notifiers/gen/sms"
)

// sms service example implementation.
// The example methods log the requests and return zero values.
type smssrvc struct {
	logger log.Logger
}

// NewSms returns the sms service implementation.
func NewSms(logger log.Logger) sms.Service {
	return &smssrvc{logger}
}

// Sent whenever an MSP confirms or rejects delivery of a message.
func (s *smssrvc) DeliveryReport(ctx context.Context, p *sms.DeliveryReport1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("sms.deliveryReport"))
	return
}

// Sent whenever a message is sent to any of your registered shortcodes.
func (s *smssrvc) IncomingMessage(ctx context.Context, p *sms.IncomingMessage1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("sms.incomingMessage"))
	return
}

// Sent whenever a user opts out of receiving messages from your alphanumeric
// sender ID
func (s *smssrvc) BulkSMSOptOut(ctx context.Context, p *sms.BulkSMSOptOut1) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("sms.bulkSMSOptOut"))
	return
}

// Sent whenever someone subscribes or unsubscribes from any of your premium
// SMS products.
func (s *smssrvc) SubNotifier(ctx context.Context, p *sms.SubscriptionNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("sms.subNotifier"))
	return
}
