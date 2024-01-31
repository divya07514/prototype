package service

type Agent struct {
	Id         int
	IsReserved bool
	OrderId    int
}

type BookAgentRequest struct {
	OrderId int
}

type BookAgentResponse struct {
	AgentId int
	OrderId int
}
