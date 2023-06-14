package db

import (
	"capi_api/internals/app/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type VoteStorage struct {
	databasePool *pgxpool.Pool
}


func NewVoteStorage(pool *pgxpool.Pool) *VoteStorage {
	storage := new(VoteStorage)
	storage.databasePool = pool
	return storage
}

func (storage *VoteStorage) InsertVote(vote models.Vote) error {
	query := "INSERT INTO Vote (peer_id, candidate_id, election_id) VALUES ($1, $2, $3)"

	_, err := storage.databasePool.Exec(context.Background(),query, vote.Id_voter , vote.Id_candidate , vote.Id_election) //транзакция не нужна, у нас только один запрос

	if err != nil {
		log.Errorln(err)
		return err
	}
	
	return nil
}

