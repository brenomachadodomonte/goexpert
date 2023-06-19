package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type API struct {
	Name        string
	URL         string
	GetResponse GetResponse
}

type GetResponse func([]byte) (ApiResponse, error)

type ApiResponse struct {
	Code     string
	Address  string
	City     string
	State    string
	District string
}

func (api *API) Print(response ApiResponse) string {
	return fmt.Sprintf("API: %s, Code: %s, Address: %s, City: %s, State: %s, District: %s\n",
		api.Name,
		response.Code,
		response.Address,
		response.City,
		response.State,
		response.District)
}

func (api *API) CallAPI(cep string) string {
	url := strings.Replace(api.URL, "{cep}", cep, 1)

	req, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Error while requesting: %s\n", err.Error())
	}
	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Sprintf("Error while getting request body: %s\n", err.Error())
	}

	response, err := api.GetResponse(data)
	if err != nil {
		return fmt.Sprintf("Error while getting response: %s\n", err.Error())
	}
	return api.Print(response)
}

func GetResponseViaCep(data []byte) (ApiResponse, error) {
	type alias struct {
		Code     string `json:"cep"`
		Address  string `json:"logradouro"`
		City     string `json:"localidade"`
		State    string `json:"uf"`
		District string `json:"bairro"`
	}
	var a alias
	err := json.Unmarshal(data, &a)
	if err != nil {
		return ApiResponse{}, err
	}
	return ApiResponse(a), nil
}

func GetResponseApiCep(data []byte) (ApiResponse, error) {
	type alias struct {
		Code     string `json:"code"`
		Address  string `json:"address"`
		City     string `json:"city"`
		State    string `json:"state"`
		District string `json:"district"`
	}
	var a alias
	err := json.Unmarshal(data, &a)
	if err != nil {
		return ApiResponse{}, err
	}
	return ApiResponse(a), nil
}

func main() {
	apiCep := API{Name: "API CEP", URL: "https://cdn.apicep.com/file/apicep/{cep}.json", GetResponse: GetResponseApiCep}
	viaCep := API{Name: "VIA CEP", URL: "http://viacep.com.br/ws/{cep}/json/", GetResponse: GetResponseViaCep}

	channelApiCep := make(chan string)
	channelViaCep := make(chan string)

	if len(os.Args) == 1 {
		fmt.Println("Your must provide a CEP")
		fmt.Println("EX: go run main.go 64001‑040")
		return
	}

	cep := os.Args[1] //get first argument

	go func() {
		channelApiCep <- apiCep.CallAPI(cep)
	}()

	go func() {
		channelViaCep <- viaCep.CallAPI(cep)
	}()

	select {
	case messageApiCep := <-channelApiCep:
		fmt.Println(messageApiCep)
	case messageViaCep := <-channelViaCep:
		println(messageViaCep)
	case <-time.After(time.Second):
		println("Timout")
	}
}
