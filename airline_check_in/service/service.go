package service

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var SeatLayout []int

func Init() {
	SeatLayout = make([]int, 120)
}

func PrintSeatLayout() {
	i := 1
	for idx := range SeatLayout {
		if SeatLayout[idx] == 0 {
			fmt.Print("." + " ")
		} else {
			fmt.Print("x" + " ")
		}
		if (i+idx)%10 == 0 {
			fmt.Println()
		}
	}
}

func MarkSeat(index int) {
	SeatLayout[index-1] = 1
}

func Reset() {
	Init()
}

func Seats() int {
	return len(SeatLayout)
}

func GetUser(db *sql.DB, userId int) (*User, error) {
	var id int
	var name string
	query := "SELECT * FROM users WHERE id = ?"
	err := db.QueryRow(query, userId).Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Msg(err.Error())
		} else if err != nil {
			log.Error().Msg(err.Error())
		}
		return nil, err
	}
	return &User{
		Id:   id,
		Name: name,
	}, nil
}

func GetSeat(db *sql.DB, seatNum int) (*Seat, error) {
	var id int
	var name sql.NullString
	var tripId sql.NullString
	var userId sql.NullString
	query := "SELECT * FROM seats WHERE id = ?"
	err := db.QueryRow(query, seatNum).Scan(&id, &name, &tripId, &userId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Msg(err.Error())
		} else if err != nil {
			log.Error().Msg(err.Error())
		}
		return nil, err
	}
	return &Seat{
		Id:     id,
		Name:   name,
		TripId: tripId,
		UserId: tripId,
	}, nil
}

func GetSeatAtRandom(db *sql.DB) (*Seat, error) {
	var id int
	var name sql.NullString
	var tripId sql.NullString
	var userId sql.NullString
	rand.Seed(time.Now().UnixMilli())
	randSeat := rand.Intn(Seats())
	query := "SELECT * FROM seats WHERE id = ?"
	err := db.QueryRow(query, randSeat).Scan(&id, &name, &tripId, &userId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Msg(err.Error())
		} else if err != nil {
			log.Error().Msg(err.Error())
		}
		return nil, err
	}
	return &Seat{
		Id:     id,
		Name:   name,
		TripId: tripId,
		UserId: userId,
	}, nil
}

func ReadInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func Wait(i int) {
	duration := time.Duration(i) * time.Second
	time.Sleep(duration)
}
