package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Game struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	PlayTime  float32 `json:"playtime"`
}

var library = []Game{
	{ID: "1", Title: "Sekiro : Shadow Die Twice", Publisher: "Activision", PlayTime: 80.5},
	{ID: "2", Title: "Elden Ring", Publisher: "Bandai Namco", PlayTime: 120.2},
	{ID: "3", Title: "Dishonored 2", Publisher: "Bethesda", PlayTime: 40.0},
}

func main() {
	router := gin.Default()
	router.GET("/games", getGameList)
	router.GET("/games/:id", getGameById)
	router.POST("/games", postNewGame)
	router.PUT("/games/:id", updateGame)

	router.Run("localhost:8080")
}

func getGameList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, library)
}

func getGameById(context *gin.Context) {
	var id = context.Param("id")

	for _, v := range library {
		if v.ID == id {
			context.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Game not found"})
}

func postNewGame(context *gin.Context) {
	var newGame Game

	if err := context.BindJSON(&newGame); err != nil {
		return
	}

	library = append(library, newGame)
	context.IndentedJSON(http.StatusOK, library)
}

func updateGame(context *gin.Context) {
	var id = context.Param("id")
	var updatedGame Game

	if err := context.BindJSON(&updatedGame); err != nil {
		return
	}

	for i := 0; i < len(library); i++ {
		v := library[i]
		if v.ID == id {
			if updatedGame.Publisher != "" {
				v.Publisher = updatedGame.Publisher
			}
			if updatedGame.PlayTime != 0.0 {
				v.PlayTime = updatedGame.PlayTime
			}

			library[i] = v
			context.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to update game data"})
}
