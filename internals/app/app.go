package app

import (
	"capi_api/api"
	"capi_api/api/middleware"
	db3 "capi_api/internals/app/db"
	"capi_api/internals/app/handlers"
	"capi_api/internals/app/processors"
	"capi_api/internals/cfg"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

type appServer struct {
	config cfg.Cfg
	srv    *http.Server
	db     *pgxpool.Pool
	logger *zerolog.Logger
}

func NewServer(config cfg.Cfg, logger *zerolog.Logger) *appServer { 
	return &appServer{
		config: config,
		logger: logger,
	}
}

func (server *appServer) Serve(ctx context.Context) error {
	server.logger.Info().Msg("Starting server")
	server.logger.Info().Msg(server.config.GetDBString())

	var err error

	db, err := pgxpool.Connect(ctx, server.config.GetDBString()) 
	if err != nil {
		server.logger.Err(err)

		err = fmt.Errorf("failed to connect to the database %w. ", err)

		return err
	}

	server.db = db

	voteStrorage := db3.NewVoteStorage(server.db)
	peerStrorage := db3.NewPeerStorage(server.db)
	peerAuthStrorage := db3.NewPeerAuthStorage(server.db)
	candidateStrorage := db3.NewCandidateStorage(server.db)

	candidateProcessor := processors.NewCandidateProcessor(candidateStrorage)
	voteProcessor := processors.NewVoteProcessor(voteStrorage)
	peerProcessor := processors.NewPeerProcessor(peerStrorage)
	peerAuthProcessor := processors.NewPeerAuthProcessor(peerAuthStrorage)

	voteHandler := handlers.NewVoteHandler(voteProcessor)
	peerHandler := handlers.NewPeerHandler(peerProcessor)
	peerAuthHandler := handlers.NewPeerAuthHandler(peerAuthProcessor)
	candidateHandler := handlers.NewCandidateHandler(candidateProcessor)

	routes := api.CreateRoutes(voteHandler, peerHandler, peerAuthHandler, candidateHandler) 
	routes.Use(middleware.RequestLog)                                                      

	server.srv = &http.Server{
		Addr:    "0.0.0.0:" + server.config.Port,
		Handler: routes,
	}

	server.logger.Info().Msg("Server started.")

	err = server.srv.ListenAndServe() 

	if err != nil && err != http.ErrServerClosed {
		server.logger.Err(err).Msg("Failure while serving")

		return err
	}

	return nil
}

func (server *appServer) Shutdown() error {
	server.logger.Info().Msg("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	server.db.Close()

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		server.logger.Err(err)

		err = fmt.Errorf("server shutdown failed %w. ", err)

		return err
	}

	server.logger.Info().Msg("Bye!")

	return nil
}
