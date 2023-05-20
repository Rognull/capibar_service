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

type VoteStorage struct {
	databasePool *pgxpool.Pool
}

// type userCar struct { //наша joined структура
// 	Id_voter   int64 `db:"id_voter"` // TODO измениить мапы полей
// 	Id_Vote_one string `db:"Vote_one"`
// 	Id_Vote_two string `db:"Vote_two"`
// }

func NewVoteStorage(pool *pgxpool.Pool) *VoteStorage {
	storage := new(VoteStorage)
	storage.databasePool = pool
	return storage
}

func (storage *VoteStorage) InsertVote(vote models.Vote) error {
	query := "INSERT INTO Vote(Id_voter, Id_candidate, Id_election VALUES ($1, $2, $3)"

	_, err := storage.databasePool.Exec(context.Background(),query, vote.Id_voter , vote.Id_candidate , vote.Id_election) //транзакция не нужна, у нас только один запрос

	if err != nil {
		log.Errorln(err)
		return err
	}
	
	return nil
}




// func convertJoinedQueryToCar(input userCar) models.Car { //сворачиваем плоскую структуру в нашу рабочую модель
// 	return models.Vote{
// 		Id_voter:Id_voter,
// 		Id_Vote_one:input.Id_Vote_one,
// 		Id_Vote_two:input.Id_Vote_two,
// 	}
// }