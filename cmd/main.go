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

func main() { //точка входа нашего сервера
	logger := new(zerolog.Logger)

	config := cfg.LoadAndStoreConfig() //грузим конфигурацию

	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст для работы контекстнозависимых частей системы

	c := make(chan os.Signal, 1) //создаем канал для сигналов системы

	signal.Notify(c, os.Interrupt) //в случае поступления сигнала завершения - уведомляем наш канал

	server := app.NewServer(config, logger) // создаем сервер

	go func() { //горутина для ловли сообщений системы
		oscall := <-c //если таки что то пришло

		logger.Info().Msg(fmt.Sprintf("system call:%+v", oscall))

		if err := server.Shutdown(); err != nil { //выключаем сервер
			logger.Err(err)
		}

		cancel() //отменяем контекст
	}()

	if err := server.Serve(ctx); err != nil { //запускаем сервер
		logger.Err(err)
	}
}
