package rctx

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func InitAppTerm(ctx context.Context, onTerm func()) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-sig
		if onTerm != nil {
			onTerm()
		}
		cancel()
	}()
	return ctx
}
