package main

import (
	"github.com/thomaspaulin/snc-server-go/snc"
	"github.com/gin-gonic/gin"
	"net/http"
)



//------------------------------------------------------------------------------------------------//
// Matches
//------------------------------------------------------------------------------------------------//
func GetDivisions(c *gin.Context) {
	// todo handle deleted case
	db := DB.Find(&snc.Division{})
	if db.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, db.Value)
}

func CreateDivision(c *gin.Context) {
	d := snc.Division{}
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		if err := DB.Create(d); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
	}
}
