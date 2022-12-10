package main

import (
	"fmt"
	"github.com/openhms/openhms-api/internal"
	"github.com/openhms/openhms-api/pkg/router"
	"log"
	"os"
)

func main() {
	jsonFile, jsonError := os.Open("config.json")
	if jsonError != nil {
		log.Fatalf("Error opening file %v", jsonError)
	}
	config, err := internal.ParseConfig(jsonFile)
	if err != nil {
		log.Fatalf("Error parsing json file %v", err)
	}
	db, connectionError := internal.OpenDatabase(&config.Settings)
	if connectionError != nil {
		log.Fatalf("Error occurred while connecting to database %s", connectionError)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	port := config.Settings.Port
	r, v1 := router.InitRouter()
	router.Routes(v1)
	serverErr := r.Run(fmt.Sprintf(":%d", port))
	if serverErr != nil {
		return
	}
}
