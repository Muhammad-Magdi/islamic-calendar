package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammad-magdi/islamic-calendar/fasting"
	"github.com/muhammad-magdi/islamic-calendar/prayers"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")

	r := gin.Default()

	prayersRouter := prayers.NewPrayersRouter()
	r.GET("/v0/prayer-times", prayersRouter.GetPrayerTimes)
	r.GET("/v0/calculation-methods", prayersRouter.GetCalculationMethods)

	fastingRouter := fasting.NewFastingRouter(fasting.NewFastingService())
	r.GET("/v0/current-month-fasting-days", fastingRouter.GetCurrentMonthFastingDays)

	r.Static("/docs", "./swagger-ui")

	r.Run(fmt.Sprintf(":%s", port))
}
