package main

import (
	"context"
	"fmt"

	"github.com/zahra-ash0uri/mabna-task/models"
)

func main() {
	conn, err := models.ConnectToPostgres()
	if err != nil {
		fmt.Println("Can not connect to database...")
		return
	}
	fmt.Println(conn)
	defer conn.Close(context.Background())

}
