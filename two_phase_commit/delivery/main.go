package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"net/http"
	"prototype/two_phase_commit/delivery/service"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1Path := r.Group("/delivery/agent")

	{
		v1Path.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, "up and running")
		})

		v1Path.POST("/reserve", func(c *gin.Context) {
			agent, err := service.ReserveAgent()
			if err != nil {
				c.JSON(429, err)
				return
			}
			c.JSON(200, agent.Id)
		})
		v1Path.POST("/book", func(c *gin.Context) {
			var req service.BookAgentRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.AbortWithStatus(400)
			}
			agent, err := service.BookAgent(req.OrderId)
			if err != nil {
				c.JSON(429, err)
				return
			}
			c.JSON(200, service.BookAgentResponse{
				AgentId: agent.Id,
				OrderId: req.OrderId,
			})
		})
	}
	log.Info().Msg("running delivery service on port 8082")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
