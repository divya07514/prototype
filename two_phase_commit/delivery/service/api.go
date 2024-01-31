package service

import (
	"database/sql"
	"errors"
	"github.com/rs/zerolog/log"
	"prototype/two_phase_commit/db"
)

func ReserveAgent() (*Agent, error) {
	log.Info().Msg("reserving agent")
	DB, err := db.GetConn()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	txn, _ := DB.Begin()
	row := txn.QueryRow(db.SelectAgent)
	if row.Err() != nil {
		log.Error().Msg(row.Err().Error())
		txn.Rollback()
		return nil, row.Err()
	}
	var agent Agent
	err = row.Scan(&agent.Id, &agent.IsReserved, &agent.OrderId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, errors.New("no delivery agent found")
	}
	_, err = txn.Exec(db.BlockAgent, agent.Id)
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
	return &agent, nil
}

func BookAgent(orderId int) (*Agent, error) {
	log.Info().Msg("booking agent")
	DB, err := db.GetConn()
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	txn, _ := DB.Begin()
	row := txn.QueryRow(db.BookAgent)
	if row.Err() != nil {
		log.Error().Msg(row.Err().Error())
		txn.Rollback()
		return nil, row.Err()
	}
	var agent Agent
	err = row.Scan(&agent.Id, &agent.IsReserved, &agent.OrderId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Error().Msg(err.Error())
		txn.Rollback()
		return nil, errors.New("no delivery agent found")
	}
	_, err = txn.Exec(db.UpdateAgent, orderId, agent.Id)
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
	agent.IsReserved = false
	agent.OrderId = orderId
	return &agent, nil
}
