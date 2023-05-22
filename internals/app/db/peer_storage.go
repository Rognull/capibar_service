package db

import (
	"capi_api/internals/app/models"
	"context"

	// "errors"
	// "fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	//"models"
)

type PeerStorage struct {
	databasePool *pgxpool.Pool
}

// type userCar struct { //наша joined структура
// 	Id_voter   int64 `db:"id_voter"` // TODO измениить мапы полей
// 	Id_Vote_one string `db:"Vote_one"`
// 	Id_Vote_two string `db:"Vote_two"`
// }

func NewPeerStorage(pool *pgxpool.Pool) *PeerStorage {
	storage := new(PeerStorage)
	storage.databasePool = pool
	return storage
}

// func (storage *PeerStorage) GetPeerList(nicknameFilter string, codeFilter string) []models.Peer { //TODO test
// 	query := "SELECT peer.nickname AS nickname, peer.code as code  FROM peer WHERE 1=1"

// 	placeholderNum := 1
// 	args := make([]interface{},0)
// 	if userIdFilter != 0 { //задаем фильтры через плейсхолдеры
// 		query += fmt.Sprintf(" AND users.id = $%d", placeholderNum)
// 		args = append(args, userIdFilter) //сразу же добавляем аргумент для фильтра
// 		placeholderNum++ //увеличиваем номер плейсхолдеру
// 	}
// 	if brandFilter != "" {
// 		query += fmt.Sprintf(" AND brand ILIKE $%d", placeholderNum)
// 		args = append(args, fmt.Sprintf("%%%s%%",brandFilter))
// 		placeholderNum++
// 	}
// 	if colourFilter != "" {
// 		query += fmt.Sprintf(" AND colour ILIKE $%d", placeholderNum)
// 		args = append(args, fmt.Sprintf("%%%s%%",colourFilter))
// 		placeholderNum++
// 	}
// 	if licenseFilter != "" {
// 		query += fmt.Sprintf(" AND license_plate ILIKE $%d", placeholderNum)
// 		args = append(args, fmt.Sprintf("%%%s%%",licenseFilter))
// 		placeholderNum++
// 	}

// 	var dbResult []userCar

// 	err := pgxscan.Select(context.Background(), storage.databasePool, &dbResult, query, args...)

// 	if err != nil {
// 		log.Errorln(err)
// 	}

// 	result := make([]models.Car, len(dbResult))

// 	for idx, dbEntity := range dbResult {
// 		result[idx] = convertJoinedQueryToCar(dbEntity) //заполняем результат
// 	}

// 	return result
// }

func (storage *PeerStorage) FindPeerByNickname(peer_nick string) models.Peer {
	query := "SELECT * FROM peer  WHERE peer.nickname = $1"

	var result models.Peer

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, peer_nick) //забираем по nickname
	
	if err != nil {
		log.Errorln(err)
	}

	return result
}
