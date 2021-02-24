package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"./models"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body

	var login models.Login

	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(value, &login)

	c.JSON(200, gin.H{
		"Nome":  string(login.Name),
		"Senha": string(login.Password),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Print("Test")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)
	r.GET("/path/:name/:age", PathParameters)
	r.Run(":8081")
}
