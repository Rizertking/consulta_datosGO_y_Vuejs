package api

import (
	"GO_APP/pkg/config" // Ajusta los nombres de los paquetes según tu configuración real
	"GO_APP/pkg/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FetchStocksHandler es el manejador para el endpoint /stocks
func FetchStocksHandler(c *gin.Context) {
	response, err := fetchStocksData("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// fetchStocksData realiza la petición a la API externa y devuelve los datos procesados
func fetchStocksData(nextPageKey string) (model.Response, error) {
	url := config.BaseURL // Usa la constante BaseURL desde el paquete config
	if nextPageKey != "" {
		url += "?next_page=" + nextPageKey
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request: ", err)
		return model.Response{}, err
	}

	req.Header.Set("Authorization", config.AuthToken) // Usa la constante AuthToken
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making HTTP request: ", err)
		return model.Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return model.Response{}, err
	}

	var response model.Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Error unmarshalling response: ", err)
		return model.Response{}, err
	}

	return response, nil
}
