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

// func (storage *CarsStorage) CreateCar(car models.Car) error {
// 	ctx := context.Background()
// 	tx, err := storage.databasePool.Begin(ctx) //здесь будем пользоваться транзакцией, чтобы проверка пользователей и вставка нового автомобиля выглядели одним запросом с ее точки зрения
// 	defer func() {
// 		err = tx.Rollback(context.Background())
// 		if err != nil {
// 			log.Errorln(err)
// 		}
// 	}()

// 	query := "SELECT id FROM peer WHERE nickname = $1"

// 	id := -1

// 	err = pgxscan.Get(ctx, tx, &id, query, car.Owner.Id)
// 	if err != nil {
// 		log.Errorln(err)
// 		err = tx.Rollback(context.Background()) //если получили ошибку откатываем транзакцию целиком
// 		if err != nil {
// 			log.Errorln(err)
// 		}
// 		return err
// 	}

// 	if id == -1 {
// 		return errors.New("user ")
// 	}

// 	insertQuery := "INSERT INTO cars(user_id, colour, brand, license_plate) VALUES ($1,$2,$3,$4)"

// 	_, err = tx.Exec(context.Background(),insertQuery, car.Owner.Id, car.Colour, car.Brand, car.LicensePlate) //вызываем exec НЕ У соединения а У транзакции

// 	if err != nil {
// 		log.Errorln(err)
// 		err = tx.Rollback(context.Background())
// 		if err != nil {
// 			log.Errorln(err)
// 		}
// 		return err
// 	}
// 	err = tx.Commit(context.Background()) // в конце посылаем транзакцию, база сохранит значения, если до этого ничего не было откачено
// 	if err != nil {
// 		log.Errorln(err)
// 	}

// 	return err
// }
