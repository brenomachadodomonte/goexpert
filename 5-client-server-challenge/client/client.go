package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Quotation struct {
	Value float64 `json:"value"`
}

func main() {
	quotation, err := GetQuotation()
	if err != nil {
		panic(err)
	}

	err = SaveFile(quotation)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cotação Atual: $1.00 = R$%.4f", quotation.Value)
}

func GetQuotation() (*Quotation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		return nil, err
	}

	return &quotation, nil
}

func SaveFile(quotation *Quotation) error {
	filename := "cotacao.txt"

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	valueStr := fmt.Sprintf("Dólar: %.4f", quotation.Value)
	_, err = file.Write([]byte(valueStr))
	if err != nil {
		return err
	}
	return nil
}
