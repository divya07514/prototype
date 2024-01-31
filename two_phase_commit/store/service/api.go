package service

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"prototype/two_phase_commit/db"
)

func ReserveFood(foodId int) (*Food, error) {
	log.Info().Msg("reserving food")
	DB, err := db.GetConn()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	txn, _ := DB.Begin()
	row := txn.QueryRow(db.SelectFood, foodId)
	if row.Err() != nil {
		log.Error().Msg(row.Err().Error())
		txn.Rollback()
		return nil, row.Err()
	}
	var food Food
	err = row.Scan(&food.Id, &food.FoodId, &food.IsReserved, &food.OrderId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, errors.New("no food packet found")
	}
	_, err = txn.Exec(db.BlockFood, food.Id)
	if err != nil {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, err
	}
	err = txn.Commit()
	if err != nil {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, err
	}
	return &food, nil
}

func BookFood(orderId int, foodId int) (*Food, error) {
	log.Info().Msg("booking food")
	DB, err := db.GetConn()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	txn, _ := DB.Begin()
	row := txn.QueryRow(db.BookFood, foodId)
	if row.Err() != nil {
		log.Error().Msg(row.Err().Error())
		txn.Rollback()
		return nil, row.Err()
	}
	var food Food
	err = row.Scan(&food.Id, &food.FoodId, &food.IsReserved, &food.OrderId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, errors.New("no food packet found")
	}
	_, err = txn.Exec(db.UpdateFood, orderId, food.Id)
	if err != nil {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, err
	}
	err = txn.Commit()
	if err != nil {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, err
	}
	food.IsReserved = false
	food.OrderId = orderId
	return &food, nil
}
