package notifier

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	payments "github.com/wondenge/at-notifiers/gen/payments"
)

// payments service example implementation.
// The example methods log the requests and return zero values.
type paymentssrvc struct {
	logger log.Logger
}

// NewPayments returns the payments service implementation.
func NewPayments(logger log.Logger) payments.Service {
	return &paymentssrvc{logger}
}

// Payment Notifications
func (s *paymentssrvc) PaymentNotifier(ctx context.Context, p *payments.PaymentNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("payments.paymentNotifier"))
	return
}

// C2B Validation Notifications
func (s *paymentssrvc) C2bNotifier(ctx context.Context, p *payments.C2BValidationNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("payments.c2bNotifier"))
	return
}

// B2C Validation Notifications
func (s *paymentssrvc) B2cNotifier(ctx context.Context, p *payments.B2CValidationNotificationPayload) (res *payments.B2CValidationNotificationResponse, err error) {
	res = &payments.B2CValidationNotificationResponse{}
	s.logger.Log("info", fmt.Sprintf("payments.b2cNotifier"))
	return
}
