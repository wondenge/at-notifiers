package notifier

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	voice "github.com/wondenge/at-notifiers/gen/voice"
)

// voice service example implementation.
// The example methods log the requests and return zero values.
type voicesrvc struct {
	logger log.Logger
}

// NewVoice returns the voice service implementation.
func NewVoice(logger log.Logger) voice.Service {
	return &voicesrvc{logger}
}

// Voice Notification delivered to our callback URL
func (s *voicesrvc) VoiceNotifier(ctx context.Context, p *voice.VoiceNotification) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("voice.voiceNotifier"))
	return
}

// Event Notifications sent from AT after call transfer initiated.
func (s *voicesrvc) TransferEvents(ctx context.Context, p *voice.CallTransferEvent) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("voice.transferEvents"))
	return
}
