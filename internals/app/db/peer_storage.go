package db

import (
	"capi_api/internals/app/models"
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type PeerStorage struct {
	databasePool *pgxpool.Pool
}

func NewPeerStorage(pool *pgxpool.Pool) *PeerStorage {
	storage := new(PeerStorage)
	storage.databasePool = pool
	return storage
}

func (storage *PeerStorage) FindPeerByNickname(peer_nick string) models.Peer {
	query := "SELECT * FROM peer  WHERE peer.nickname = $1"

	var result models.Peer

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, peer_nick) //забираем по nickname

	if err != nil {
		log.Errorln(err)
	}

	return result
}
