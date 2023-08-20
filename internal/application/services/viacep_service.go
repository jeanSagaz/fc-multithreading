package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jeanSagaz/multithreading/internal/application/dto"
)

func GetViaCepService(zipCode string) (*dto.ViaCepResponse, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", zipCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var viaCepResponse dto.ViaCepResponse
	err = json.Unmarshal(res, &viaCepResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println(url)
	return &viaCepResponse, nil
}
