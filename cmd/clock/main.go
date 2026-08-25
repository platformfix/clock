package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/platformfix/clock/internal/clock"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	slog.Info("starting clock")
	clock.Run(ctx, os.Stdout, time.Second)
	slog.Info("shutting down")
}
