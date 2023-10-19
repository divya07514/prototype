package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"prototype/airline_check_in/service"
	"strconv"
	"time"
)

// Randomly assigning a seat to a single user
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
	seats, _ := service.GetSeatAtRandom(DB)
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
	log.Info().Msgf("selecting a random seat for the user %s out of total 120 seats", user.Name)
	rand.Seed(time.Now().UnixMilli())
	service.Wait(1)
	txn, err := db.Begin()
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
	log.Info().Msgf("%s booked the seat %s", user.Name, seats.Name.String)
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
