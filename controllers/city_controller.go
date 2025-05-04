package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/danilosmaciel/api-go-gin/database"
	"github.com/danilosmaciel/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CityImport(c *gin.Context) {
	var allCities []models.City
	var states []models.State

	getStates(c, &states)

	for _, s := range states {
		citiesTEmp := getCitysByState(s.Sigla)
		allCities = append(allCities, citiesTEmp...)
	}

	database.DB.CreateInBatches(allCities, len(allCities))

	print("importacao de cidades finalizada!")

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func getCitysByState(sigla string) []models.City {
	client := &http.Client{}
	var cities []models.City
	url := fmt.Sprintf("https://servicodados.ibge.gov.br/api/v1/localidades/estados/%v/distritos", sigla)
	req, err := http.NewRequest("GET", url, nil)
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

	var data []map[string]interface{}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return cities
	}

	if err := json.Unmarshal(
		body,
		&data,
	); err != nil {
		fmt.Println("Erro ao fazer unmarshalling:", err)
		return cities
	}

	for _, d := range data {
		// Navegue pela estrutura aninhada para obter o StateID
		municipio, ok := d["municipio"].(map[string]interface{})
		if !ok {
			fmt.Println("Erro: campo 'municipio' não encontrado ou inválido")
			return cities
		}

		regiaoImediata, ok := municipio["regiao-imediata"].(map[string]interface{})
		if !ok {
			fmt.Println("Erro: campo 'regiao-imediata' não encontrado ou inválido")
			return cities
		}

		regiaoIntermediaria, ok := regiaoImediata["regiao-intermediaria"].(map[string]interface{})
		if !ok {
			fmt.Println("Erro: campo 'regiao-intermediaria' não encontrado ou inválido")
			return cities
		}

		uf, ok := regiaoIntermediaria["UF"].(map[string]interface{})
		if !ok {
			fmt.Println("Erro: campo 'UF' não encontrado ou inválido")
			return cities
		}

		// Crie a cidade e preencha os campos básicos
		cities = append(cities, models.City{
			IbgeCode: int32(d["id"].(float64)),
			Name:     d["nome"].(string),
			StateID:  int32(uf["id"].(float64)),
		})
		// Atribua o ID da UF ao StateID
	}

	return cities
}
