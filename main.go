package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	UpdatedAt string `json:"updated_at"`
	Type      string `json:"type"`
	BuyUsd    string `json:"buy_usd"`
	SellUsd   string `json:"sell_usd"`
	BuyEur    string `json:"buy_eur"`
	SellEur   string `json:"sell_eur"`
	BuyRub    string `json:"buy_rub"`
	SellRub   string `json:"sell_rub"`
	BuyKzt    string `json:"buy_kzt"`
	SellKzt   string `json:"sell_kzt"`
}

func MakeRequest(url *string, token *string) (request *http.Request) {
	request, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Authorization", "Bearer "+*token)
	return
}

func SendRequest(request *http.Request) (body []byte, status_code int8) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body, int8(resp.StatusCode)
}

func main() {
	var url string = "https://data.fx.kg/api/v1/average"
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	var token string = os.Getenv("TOKEN")

	request := MakeRequest(&url, &token)
	body, _ := SendRequest(request)

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Buying\nUSD: %s\nEUR: %s\nRUB: %s\nKZT: %s\n\n", response.BuyUsd, response.BuyEur, response.BuyRub, response.BuyKzt)
	fmt.Printf("Selling\nUSD: %s\nEUR: %s\nRUB: %s\nKZT: %s\n", response.SellUsd, response.SellEur, response.SellRub, response.SellKzt)
}
