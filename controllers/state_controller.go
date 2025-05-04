package controllers

import (
	"encoding/json"
	"github.com/danilosmaciel/api-go-gin/database"
	"github.com/danilosmaciel/api-go-gin/dtos"
	"github.com/danilosmaciel/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

func StateImport(c *gin.Context) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://servicodados.ibge.gov.br/api/v1/localidades/estados", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var states []models.State
	err = json.NewDecoder(resp.Body).Decode(&states)

	if err != nil {
		panic(err)
	}

	database.DB.CreateInBatches(states, len(states))

	print("importacao de estados finalizada!")

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func getStates(c *gin.Context, states *[]models.State) {
	errState := database.DB.Find(&states)

	if errState.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error finding states"})
	}

	if states == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "table state is empty",
		})
	}
}

func StateHandler(c *gin.Context) {
	var states []models.State
	errState := database.DB.Preload("Cities").Find(&states).Error
	if errState != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "impossible data load",
		})
		return
	}

	var statesdDto []dtos.StateDTO

	for _, state := range states {
		stateDto := dtos.StateDTO{
			Code:  state.Code,
			Sigla: state.Sigla,
		}

		for _, city := range state.Cities {
			citName := strings.ReplaceAll(city.Name, "'", " ")
			citName = strings.Trim(citName, " ")

			stateDto.Cities = append(stateDto.Cities, dtos.CityDTO{IbgeCode: city.IbgeCode, Name: citName})
		}

		statesdDto = append(statesdDto, stateDto)
	}

	c.JSON(http.StatusOK, statesdDto)
}
