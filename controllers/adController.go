package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaodrts/Drts.Ads.GO.API/entities"
)

func Create(c *gin.Context) {

	fmt.Printf("chamou o criate, SERVER")

	var newAd entities.Ad

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if err := c.BindJSON(&newAd); err != nil {
		return
	}

	_, err := entities.Insert(newAd)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newAd)
}

func Get(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	listAds, err := entities.Get("0")

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, listAds)
}

func GetById(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	id := c.Param("id")

	listAds, err := entities.GetById(id)

	if err != nil {
		return
	}

	if listAds.ID == "0" || listAds.ID == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ad not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, listAds)
}

func Update(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	id := c.Param("id")

	var newAd entities.Ad

	if err := c.BindJSON(&newAd); err != nil {
		return
	}

	_, err := entities.Update(id, newAd)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ad not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, "record updated successfully.")
}

func Delete(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	id := c.Param("id")

	_, err := entities.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ad not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, "record deleted successfully.")
}
