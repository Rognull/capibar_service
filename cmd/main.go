package main

import (
	"capi_api/internals/app"
	"capi_api/internals/cfg"
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/rs/zerolog"
)

func main() { 
	logger := new(zerolog.Logger)

	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1) 

	signal.Notify(c, os.Interrupt) 

	server := app.NewServer(config, logger)

	go func() { 
		oscall := <-c 

		logger.Info().Msg(fmt.Sprintf("system call:%+v", oscall))

		if err := server.Shutdown(); err != nil { 
			logger.Err(err)
		}

		cancel()
	}()

	if err := server.Serve(ctx); err != nil {
		logger.Err(err)
	}
}
