package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func OrderShowByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	query := "SELECT i.Name, t.DateEn FROM trade t join on instrument i where t.InstrumentID = i.Id GROUPBY I.NAME ORDER BY t.DateEn;"
	fmt.Println(query)
	// if err := models.Connection.Exec(context.Background(), "SELECT * from orders where order_id = $1", order_id).Scan(&orders).Error; err != nil {
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
