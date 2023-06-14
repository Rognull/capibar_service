package db

import (
	"capi_api/internals/app/models"
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type CandidateStorage struct {
	databasePool *pgxpool.Pool
}

func NewCandidateStorage(pool *pgxpool.Pool) *CandidateStorage {
	storage := new(CandidateStorage)
	storage.databasePool = pool
	return storage
}

func (storage *CandidateStorage) FindCandidateByNickname(peer_nick string) models.CandidateWithName {
	query := "SELECT candidate.id, candidate.peer_id, t1.nickname as nickname, convocation_id, promises FROM candidate JOIN peer on candidate.peer_id = peer.id"

	var result models.CandidateWithName

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, peer_nick) //забираем по nickname
	
	if err != nil {
		log.Errorln(err)
	}

	return result
}


func (storage *CandidateStorage) GetCandidateList(Convocation uint64) []models.CandidateWithName { //TODO test
	query := "SELECT candidate.peer_id, peer.nickname as nickname, convocation_id, promises FROM candidate JOIN peer on candidate.peer_id = peer.id where candidate.convocation_id = $1"

	
	var dbResult []models.CandidateWithName

	err := pgxscan.Select(context.Background(), storage.databasePool, &dbResult, query, Convocation)

	if err != nil {
		log.Errorln(err)
	}

	return dbResult
}
