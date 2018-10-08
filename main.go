package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/minhthuy30197/template/config"
	"github.com/minhthuy30197/template/controller"
	"github.com/minhthuy30197/template/model"
	"github.com/minhthuy30197/template/router"
)

func main() {
	r := gin.Default()

	// Đọc cấu hình từ file json
	config := config.SetupConfig()

	// Khởi tạo controller
	c := controller.NewController()
	c.Config = config

	// Kết nối CSDL
	dbConfig := config.Database
	db := model.ConnectDb(dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Address)
	defer db.Close()
	c.DB = db

	err := model.MigrationDb(db, config.ServiceName)
	if err != nil {
		panic(err)
	}

	router.SetupRouter(config, r, c)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
