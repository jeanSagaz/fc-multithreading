package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jeanSagaz/multithreading/internal/application/dto"
)

func GetApiCepService(zipCode string) (*dto.ApiCepResponse, error) {
	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", zipCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiCepResponse dto.ApiCepResponse
	err = json.Unmarshal(res, &apiCepResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println(url)
	return &apiCepResponse, nil
}
