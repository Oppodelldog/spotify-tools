package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/Oppodelldog/spotify-sleep-timer/app/handler"
	"github.com/Oppodelldog/spotify-sleep-timer/app/sleep"
	"github.com/Oppodelldog/spotify-sleep-timer/config"
	"github.com/Oppodelldog/spotify-sleep-timer/logger"
)

func main() {
	ctx := NewSignalContext(os.Interrupt)

	config.Init()
	sleep.StartTimerWorker(ctx)

	s := &http.Server{
		Addr:    config.BindAddress,
		Handler: handler.Router(),
	}

	go closeServer(ctx, s)

	err := s.ListenAndServe()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		logger.Std.Error(err.Error())
	}

	logger.Std.Debug("server stopped")
}

func closeServer(ctx context.Context, s *http.Server) {
	<-ctx.Done()

	err := s.Close()
	if err != nil {
		logger.Std.Error(err.Error())
	}
}

func NewSignalContext(signals ...os.Signal) context.Context {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, signals...)

		<-c
		cancelFunc()
	}()

	return ctx
}
