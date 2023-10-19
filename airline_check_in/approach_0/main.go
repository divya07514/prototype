package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"os"
	"prototype/airline_check_in/service"
	"strconv"
)

// Asking user to pick a seat of his choice
// Different users can choose the same seat. There is no conflict detection mechanism present here.
func main() {
	service.Init()
	DB, err := dbConn()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	userId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	user, err := service.GetUser(DB, userId)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Info().Msgf("hello %s, which seat do you want", user.Name)
	seatNum, err := service.ReadInt()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	seats, err := service.GetSeat(DB, seatNum)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	err = book(DB, user, seats)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	service.Wait(2)
	log.Info().Msg("printing final seat layout")
	fmt.Println()
	service.PrintSeatLayout()
}

func book(db *sql.DB, user *service.User, seats *service.Seat) error {
	txn, err := db.Begin()
	if err != nil {
		return err
	}
	log.Info().Msgf("booking seat %s for user %s", seats.Name.String, user.Name)

	_, err = txn.Exec("Select id from seats where id = ?", seats.Id)
	if err != nil {
		return err
	}

	log.Info().Msg("transaction got the seat")
	service.Wait(1)

	_, err = txn.Exec("update seats set user_id = ? where id =? ", user.Id, seats.Id)
	if err != nil {
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	service.Wait(1)
	index, _ := strconv.Atoi(seats.Name.String)
	service.MarkSeat(index)
	log.Info().Msgf("%s booked the seat %s", user.Name, seats.Name)
	return nil
}

func dbConn() (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/prototype")
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return db, nil
}
