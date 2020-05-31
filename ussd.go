package notifier

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	ussd "github.com/wondenge/at-notifiers/gen/ussd"
)

// ussd service example implementation.
// The example methods log the requests and return zero values.
type ussdsrvc struct {
	logger log.Logger
}

// NewUssd returns the ussd service implementation.
func NewUssd(logger log.Logger) ussd.Service {
	return &ussdsrvc{logger}
}

// Callback URL that sends request data our App using HTTP POST.
func (s *ussdsrvc) UssdNotifier(ctx context.Context, p *ussd.USSDPayload) (res *ussd.USSDResponse, err error) {
	res = &ussd.USSDResponse{}
	s.logger.Log("info", fmt.Sprintf("ussd.ussdNotifier"))
	return
}
