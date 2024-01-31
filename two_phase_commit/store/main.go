package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"prototype/two_phase_commit/store/service"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1Path := r.Group("/store/food")
	{
		v1Path.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, "up and running")
		})
		v1Path.POST("/reserve", func(c *gin.Context) {
			var req service.ReserveFoodRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.AbortWithStatus(400)
			}
			_, err := service.ReserveFood(req.FoodId)
			if err != nil {
				c.AbortWithStatus(500)
			}
			c.JSON(200, service.ReserveFoodResponse{Reserved: true})
		})

		v1Path.POST("/book", func(c *gin.Context) {
			var req service.BookFoodRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.AbortWithStatus(400)
			}
			food, err := service.BookFood(req.OrderId, req.FoodId)
			if err != nil {
				log.Error().Msg(err.Error())
				c.AbortWithError(500, err)
			}
			c.JSON(200, service.BookFoodResponse{OrderId: food.OrderId})
		})

	}

	log.Info().Msg("running food service on port 8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
