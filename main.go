package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var customers = []Customer{
	{Name: "Irineu", Id: "1", Age: 10},
	{Name: "Jubileu", Id: "2", Age: 20},
	{Name: "Zebedeu", Id: "3", Age: 30},
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

func createCustomer(c *gin.Context) {
	var newCustomer Customer

	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newCustomer)
}

func getCustomerByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range customers {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}

func main() {
	router := gin.Default()

	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomerByID)
	router.POST("/customers", createCustomer)

	router.Run("localhost:8000")
}
