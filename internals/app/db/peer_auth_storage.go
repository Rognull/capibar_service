package db

import (
	"capi_api/internals/app/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type PeerAuthStorage struct {
	databasePool *pgxpool.Pool
}


func NewPeerAuthStorage(pool *pgxpool.Pool) *PeerAuthStorage {
	storage := new(PeerAuthStorage)
	storage.databasePool = pool
	return storage
}

func (storage *PeerAuthStorage) UpdatePassword(auth models.PeerAuth) error {
	query := "UPDATE peer_auth SET password = $1 where peer_id = $2"

	_, err := storage.databasePool.Exec(context.Background(), query, auth.Password, auth.PeerId) 
	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
