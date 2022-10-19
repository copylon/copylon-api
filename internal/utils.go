package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

func ParseConfig(file *os.File) (*Configuration, error) {
	if file == nil {
		log.Fatalf("Configuration file was not found or couldn't be found. Make sure the path is correct and the user has sufficient permission to open it.")
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var config Configuration
	err := json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, errors.New("error parsing json file. Make sure it is valid")
	}
	return &config, nil
}

func OpenDatabase(config *ConfigurationSettings) (*gorm.DB, error) {
	var dsn string
	switch {
	case config.Adapter == "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.AdapterOptions.Auth.Username,
			config.AdapterOptions.Auth.Password,
			config.AdapterOptions.Server,
			config.AdapterOptions.Port,
			config.AdapterOptions.Database)
		return gorm.Open(mysql.Open(dsn))
	case config.Adapter == "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Athens",
			config.AdapterOptions.Server,
			config.AdapterOptions.Auth.Username,
			config.AdapterOptions.Auth.Password,
			config.AdapterOptions.Database,
			config.AdapterOptions.Port)
		return gorm.Open(postgres.Open(dsn))
	default:
		log.Println("Currently, only PostgreSQL and MySQL are supported. Try using one of those database types.")
	}
	return nil, errors.New("database not currently supported")
}
