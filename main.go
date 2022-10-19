package main

import (
	"log"
	"openhms/internal"
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
}
