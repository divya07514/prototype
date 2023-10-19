package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"prototype/airline_check_in/service"
	"strconv"
	"sync"
)

// Running 120 threads, simulating behaviour where all 120 users try to check in one go
// We try to pick first available seat here, but by taking a lock on the database row using `FOR UPDATE` in the select query
// `FOR UPDATE` re-evaluates when it sees that the row on which it was waiting has changed it status.
// But this code is not Fair as other threads have to wait until the current thread releases the lock.
func main() {
	DB, err := dbConn()
	service.Init()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	var wg sync.WaitGroup
	users := len(service.SeatLayout)
	wg.Add(users)
	log.Info().Msg("all 120 users are trying to book seats simultaneously")

	for i := 1; i <= 120; i++ {
		go func(userId int) {
			user, err := service.GetUser(DB, userId)
			if err != nil {
				log.Error().Msg(err.Error())
				return
			}
			seat, err := book(DB, user)
			if err != nil {
				log.Error().Msg(err.Error())
				log.Error().Msgf("could not book seat for user %s", user.Name)
			} else {
				log.Info().Msgf("booked seat %s for user %s", seat.Name.String, user.Name)
			}
			wg.Done()
		}(i)
	}
	service.Wait(2)
	log.Info().Msg("printing final seat layout")
	fmt.Println()
	service.PrintSeatLayout()
}

func book(db *sql.DB, user *service.User) (*service.Seat, error) {
	txn, err := db.Begin()

	row := txn.QueryRow("select id, name, trip_id, user_id from seats where user_id is null order by id Limit 1 FOR UPDATE")
	if row.Err() != nil {
		return nil, row.Err()
	}

	var seat service.Seat
	err = row.Scan(&seat.Id, &seat.Name, &seat.TripId, &seat.UserId)
	if err != nil {
		return nil, err
	}
	_, err = txn.Exec("update seats set user_id = ? where id =? ", user.Id, seat.Id)
	if err != nil {
		return nil, err
	}
	err = txn.Commit()
	if err != nil {
		return nil, err
	}
	service.Wait(1)
	index, _ := strconv.Atoi(seat.Name.String)
	service.MarkSeat(index)
	return &seat, nil
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
