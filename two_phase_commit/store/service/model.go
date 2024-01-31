package service

type ReserveFoodRequest struct {
	FoodId int
}

type Food struct {
	Id         int
	FoodId     int
	IsReserved bool
	OrderId    int
}

type ReserveFoodResponse struct {
	Reserved bool
}

type BookFoodRequest struct {
	OrderId int
	FoodId  int
}

type BookFoodResponse struct {
	OrderId int
}
