package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	toml "github.com/pelletier/go-toml"
	"test.com/repositories"
	"test.com/structs"
)

func setupDB() (*gorm.DB, error) {
	configFile, err := toml.LoadFile("config.toml")
	if err != nil {
		return nil, err
	}
	dbConfig := configFile.Get("database").(*toml.Tree)
	username := dbConfig.Get("Username").(string)
	password := dbConfig.Get("Password").(string)
	dbName := dbConfig.Get("DbName").(string)
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	gormDB, err := setupDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer gormDB.Close()

	gormDB.AutoMigrate(&structs.User{}, &structs.Product{})

	r := gin.Default()

	dbHandler := repositories.DbHandler{Db: gormDB}

	setupRoutes(r, dbHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}
}
