package main

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

type Configuration struct {
	Settings ConfigurationSettings `json:"settings"`
}

type ConfigurationSettings struct {
	Adapter        string         `json:"adapter"`
	AdapterOptions AdapterOptions `json:"adapterOptions"`
}

type AdapterOptions struct {
	Server   string      `json:"server"`
	Port     int64       `json:"port"`
	Database string      `json:"database"`
	UseTLS bool        `json:"useTLS"`
	Auth   AuthOptions `json:"auth"`
}

type AuthOptions struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ParseConfig(file *os.File) *Configuration {
	if file == nil {
		log.Fatalf("Configuration file was not found or couldn't be found. Make sure the path is correct and the user has sufficient permission to open it.")
	}
	byteValue, _ := ioutil.ReadAll(file)
	var config Configuration
	json.Unmarshal(byteValue, &config)
	defer file.Close()
	return &config
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

