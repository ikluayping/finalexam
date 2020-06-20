package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikluayping/finalexam/database"
	"github.com/ikluayping/finalexam/middleware"
)

func createCustomerHandler(c *gin.Context) {
	customer := new(Customer)
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	repository := database.Repository()
	in := make(map[string]string)
	in["name"] = customer.Name
	in["email"] = customer.Email
	in["status"] = customer.Status
	row, _ := repository.Create(in)
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func getCustomers(c *gin.Context) {
	repository := database.Repository()
	rows, err := repository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var customers []Customer
	for rows.Next() {
		cust := Customer{}
		err := rows.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		customers = append(customers, cust)
	}
	c.JSON(http.StatusOK, customers)
}

func getCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	repository := database.Repository()
	row, err := repository.GetOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	customer := new(Customer)
	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func updateCustomerHandler(c *gin.Context) {
	id := c.Param("id")
	repository := database.Repository()
	row, err := repository.GetOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	customer := new(Customer)
	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	in := make(map[string]string)
	in["name"] = customer.Name
	in["email"] = customer.Email
	in["status"] = customer.Status

	if err := repository.Update(customer.ID, in); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func deleteCustByID(c *gin.Context) {
	id := c.Param("id")
	repository := database.Repository()
	if err := repository.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}

//SetupRouter :
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Authentication)
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomerByIDHandler)
	router.POST("/customers", createCustomerHandler)
	router.PUT("/customers/:id", updateCustomerHandler)
	router.DELETE("/customers/:id", deleteCustByID)

	return router
}
