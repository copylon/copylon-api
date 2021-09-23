package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	var jsonFile *os.File
	var err error
	if jsonFile, err = os.Open("config.json"); err == nil {
		config := ParseConfig(jsonFile)
		if db, connectionError := OpenDatabase(&config.Settings); connectionError == nil && db != nil {
			sqlDB, _ := db.DB()

			defer func(sqlDB *sql.DB) {
				log.Println("Closing database connection after closing application")
				sqlError := sqlDB.Close()
				if sqlError != nil {
					panic("Couldn't close database connection. Exiting")
				}
			}(sqlDB)
		} else if connectionError != nil {
			log.Fatalf("Error occurred %s", connectionError)
		}
		server := gin.Default()
		server.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello"})
			return
		})
		serveError := server.Run(":8080")
		if serveError != nil {
			return 
		}
	} else {
		log.Fatalf("Error opening file %s", err)
	}
}
