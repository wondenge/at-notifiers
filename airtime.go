package notifier

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	airtime "github.com/wondenge/at-notifiers/gen/airtime"
)

// airtime service example implementation.
// The example methods log the requests and return zero values.
type airtimesrvc struct {
	logger log.Logger
}

// NewAirtime returns the airtime service implementation.
func NewAirtime(logger log.Logger) airtime.Service {
	return &airtimesrvc{logger}
}

// Airtime Validation Notifications
func (s *airtimesrvc) Validation(ctx context.Context, p *airtime.AirtimeValidationPayload) (res *airtime.AirtimeValidationResponse, err error) {
	res = &airtime.AirtimeValidationResponse{}
	s.logger.Log("info", fmt.Sprintf("airtime.validation"))
	return
}

// Airtime Status Notifications
func (s *airtimesrvc) Status(ctx context.Context, p *airtime.AirtimeStatus) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("airtime.status"))
	return
}
