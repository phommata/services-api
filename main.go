package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phommata/services-api/config"
	"github.com/phommata/services-api/controllers"
)

func main(){
	r := gin.Default()

	appConf := config.AppConfig()
	server := fmt.Sprintf(":%d", appConf.Server.Port)

	// Setup service repository
	serviceRepo := controllers.New(appConf)

	// Routes
	r.GET("/services", serviceRepo.FindServices)

	// Run the server
	r.Run(server)
}