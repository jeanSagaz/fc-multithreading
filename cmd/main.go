package main

import (
	"fmt"
	"time"

	"github.com/jeanSagaz/multithreading/internal/application/dto"
	"github.com/jeanSagaz/multithreading/internal/application/services"
)

func main() {
	apiCep := make(chan dto.Request)
	viaCep := make(chan dto.Request)

	// ApiCep
	go func() {
		// time.Sleep(time.Second * 4)
		msg := dto.Request{ZipCode: "30530-440"}
		apiCep <- msg
	}()

	// ViaCep
	go func() {
		// time.Sleep(time.Second * 2)
		msg := dto.Request{ZipCode: "30530440"}
		viaCep <- msg
	}()

	select {
	case msg := <-apiCep:
		apiCepResponse, err := services.GetApiCepService(msg.ZipCode)
		if err != nil {
			panic(err)
		}

		fmt.Println("ApiCep")
		fmt.Println(apiCepResponse)

	case msg := <-viaCep:
		viaCepResponse, err := services.GetViaCepService(msg.ZipCode)
		if err != nil {
			panic(err)
		}

		fmt.Println("ViaCep")
		fmt.Println(viaCepResponse)

	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
