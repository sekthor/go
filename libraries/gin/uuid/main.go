package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Entity struct {
	UUID uuid.UUID `json:"uuid"`
}

func GetRandomEntity(c *gin.Context) {
	c.JSON(http.StatusOK, &Entity{
		UUID: uuid.New(),
	})
}

func GetEntity(c *gin.Context) {

	id := c.Param("id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		msg := fmt.Sprintf("'%s' is not a valid uuid", id)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}

	c.JSON(http.StatusOK, &Entity{
		UUID: uuid,
	})
}

func main() {
	router := gin.New()
	router.GET("random", GetRandomEntity)
	router.GET("entity/:id", GetEntity)
	err := router.Run()

	if err != nil {
		log.Fatal("could not start server")
	}
}
