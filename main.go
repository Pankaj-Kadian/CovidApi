package main

import (
	"covidapi/app/api"
	"covidapi/app/handlers"
	"covidapi/mongodb"
	"fmt"
	"os"
	"time"

	_ "covidapi/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	e := echo.New()
	go func() {
		for {
			final_data := api.GettingData()
			// fmt.Println(final_data)
			mongodb.InsertingNewData(final_data)
			time.Sleep(30 * time.Minute)
		}
	}()
	e.GET("GetStateData", handlers.GetCases)
	e.GET("/GetAllData", handlers.GetAllCases)
	e.GET("/GetByGeoLocation", handlers.GetDataFromGeoLocation)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	e.Logger.Fatal(e.Start(address))
}
