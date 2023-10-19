package service

import "database/sql"

type User struct {
	Id   int
	Name string
}

type Seat struct {
	Id     int
	Name   sql.NullString
	TripId sql.NullString
	UserId sql.NullString
}

type Trip struct {
	Id   int
	Name string
}
