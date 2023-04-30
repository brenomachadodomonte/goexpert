package main

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"strconv"
	"time"
)

type QuotationData struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Quotation struct {
	Data QuotationData `json:"USDBRL"`
}

type QuotationResponse struct {
	Value float64 `json:"value"`
}

type JsonError struct {
	Message string `json:"message"`
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	err = createTable(db)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", func(writer http.ResponseWriter, request *http.Request) {
		QuotationHandler(db, writer, request)
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func QuotationHandler(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	quotation, err := GetQuotation()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(&JsonError{Message: err.Error()})
	}

	resp, err := NewQuotationResponseFromQuotation(quotation)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(&JsonError{Message: err.Error()})
	}

	err = saveQuotation(db, resp.Value)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(&JsonError{Message: err.Error()})
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}

func saveQuotation(db *sql.DB, value float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	currentTime := time.Now().Format("2006-01-02 15:04")
	stmt, err := db.Prepare("insert into quotations(value, date) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, value, currentTime)
	if err != nil {
		return err
	}
	return nil
}

func createTable(db *sql.DB) error {
	table := "CREATE TABLE IF NOT EXISTS quotations(id integer primary key, value float, date string)"
	stmt, err := db.Prepare(table)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	return nil
}

func GetQuotation() (*Quotation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		return nil, err
	}

	return &quotation, nil
}

func NewQuotationResponseFromQuotation(quotation *Quotation) (*QuotationResponse, error) {
	valueStr := quotation.Data.Bid
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return nil, err
	}
	return &QuotationResponse{Value: value}, nil
}
