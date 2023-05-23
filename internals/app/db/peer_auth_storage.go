package db

import (
	"capi_api/internals/app/models"
	"context"

	// "errors"
	// "fmt"
	// "github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	//"models"
)

type PeerAuthStorage struct {
	databasePool *pgxpool.Pool
}

// type userCar struct { //наша joined структура
// 	Id_voter   int64 `db:"id_voter"` // TODO измениить мапы полей
// 	Id_Vote_one string `db:"Vote_one"`
// 	Id_Vote_two string `db:"Vote_two"`
// }

func NewPeerAuthStorage(pool *pgxpool.Pool) *PeerAuthStorage {
	storage := new(PeerAuthStorage)
	storage.databasePool = pool
	return storage
}

func (storage *PeerAuthStorage) UpdatePassword(auth models.PeerAuth) error {
	query := "UPDATE peer_auth SET password = $1 where peer_id = $2"

	_, err := storage.databasePool.Exec(context.Background(), query, auth.Password, auth.PeerId) //транзакция не нужна, у нас только один запрос
	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
