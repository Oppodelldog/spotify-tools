package sleep

import (
	"context"
	"time"

	"github.com/Oppodelldog/spotify-sleep-timer/app/storage"
	"github.com/Oppodelldog/spotify-sleep-timer/app/timer"
	"github.com/Oppodelldog/spotify-sleep-timer/logger"
	"github.com/Oppodelldog/spotify-sleep-timer/spotify/authorization"
)

func StartTimerWorker(ctx context.Context) {
	go func() {
		var t = time.NewTicker(time.Second)
		defer t.Stop()

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			case <-t.C:
				work()
			}
		}

		logger.Std.Debug("timer worker down")
	}()

	logger.Std.Debug("timer worker up")
}

func work() {
	storage.MutateAll(func(s *storage.Session) {
		checkForPause(s)
	})
}

func refreshToken(s *storage.Session) {
	if s.Spotify.RefreshErr == nil {
		token, err := authorization.Token(s.Spotify.RefreshToken)
		if err != nil {
			logger.Std.Errorf("error refreshing token failed: %v", s.ID, err)

			s.Spotify.RefreshErr = err

			return
		}

		s.Spotify.AccessToken = token.AccessToken
		s.Spotify.ExpiresIn = token.ExpiresIn
		s.Spotify.RefreshedAt = time.Now()

		return
	}
}

func checkForPause(s *storage.Session) {
	t := s.Timer
	if !t.IsSet || !t.AsDue().IsOverdue() {
		return
	}

	s.Timer = timer.Timer{}

	if s.Spotify.TokenDue().IsOverdue() {
		refreshToken(s)
	}

	if err := pause(s.Spotify.AccessToken); err == nil {
		logger.Std.Debugf("pause player for %v", s.ID)
	} else {
		logger.Std.Errorf("error pausing player for %v: %v", s.ID, err)
	}
}
