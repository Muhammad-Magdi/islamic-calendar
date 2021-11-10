package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/muhammad-magdi/islamic-calendar/prayers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	r := gin.Default()

	prayersRouter := prayers.NewPrayersRouter()
	r.GET("/v0/prayer-times", prayersRouter.GetPrayerTimes)

	r.Run(fmt.Sprintf(":%s", port))
}
