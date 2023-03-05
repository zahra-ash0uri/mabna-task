package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zahra-ash0uri/mabna-task/controllers"
	"github.com/zahra-ash0uri/mabna-task/models"
)

func main() {
	fmt.Printf("Started at : %3v \n", time.Now())

	conn, err := models.ConnectToPostgres()
	if err != nil {
		fmt.Println("Can not connect to database...")
		return
	}
	fmt.Println(conn)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/show", controllers.TradeShowByRecent)
	router.Run(":4000")

}
