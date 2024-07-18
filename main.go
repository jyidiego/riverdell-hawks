package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type club struct {
	ID       int    `json:"id"`
	TeamName string `json:"teamname"`
	Flight   string `json:"flight"`
}

var teams = []club{
	{ID: 0, TeamName: "RiverDell Hawks", Flight: "U12A"},
	{ID: 1, TeamName: "Maroons", Flight: "U12A"},
	{ID: 2, TeamName: "WC", Flight: "U12A"},
}

func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, teams)
}

func main() {
	router := gin.Default()
	router.GET("/teams", getTeams)

	router.Run(":8080")
}
