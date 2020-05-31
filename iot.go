package notifier

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	iot "github.com/wondenge/at-notifiers/gen/iot"
)

// iot service example implementation.
// The example methods log the requests and return zero values.
type iotsrvc struct {
	logger log.Logger
}

// NewIot returns the iot service implementation.
func NewIot(logger log.Logger) iot.Service {
	return &iotsrvc{logger}
}

// IoT Notifications
func (s *iotsrvc) IotNotifier(ctx context.Context, p *iot.IoTNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("iot.iotNotifier"))
	return
}
