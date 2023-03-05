package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zahra-ash0uri/mabna-task/models"
)

func TradeShowByRecent(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := models.Connection.Query(context.Background(), "SELECT i.Name, t.DateEn FROM trade t join on instrument i where t.InstrumentID = i.Id GROUPBY I.NAME ORDER BY t.DateEn;")
	if err != nil {
		log.Fatal("error while executing query")
	}

	for rows.Next() {
		values, err := rows.Values()
		fmt.Println(values...)
		if err != nil {
			log.Fatal("error while iterating dataset")
		}
	}

	// if err := models.Connection.Query(context.Background(), "SELECT * from orders where order_id = $1", order_id).Scan(&orders).Error; err != nil {
	// 	fmt.Printf("error show order by id: %3v \n", err)
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }

	// if (orders != nil) {
	//   c.JSON(http.StatusOK, orders)
	// } else {
	//   c.JSON(http.StatusOK, json.RawMessage(`[]`))
	// }
}
