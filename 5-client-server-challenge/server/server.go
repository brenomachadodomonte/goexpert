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

type QuotationModel struct {
	ID    int     `json:"id"`
	Value float64 `json:"value"`
	Date  string  `json:"date"`
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
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		GetQuotationsHandler(db, writer)
	})

	mux.HandleFunc("/cotacao", func(writer http.ResponseWriter, request *http.Request) {
		QuotationHandler(db, writer)
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func QuotationHandler(db *sql.DB, writer http.ResponseWriter) {
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

func GetQuotationsHandler(db *sql.DB, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")

	quotations, err := getAllQuotations(db)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(&JsonError{Message: err.Error()})
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(quotations)
}

func saveQuotation(db *sql.DB, value float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	currentTime := time.Now().Format("2006-01-02 15:04:05.000000")
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

func getAllQuotations(db *sql.DB) ([]QuotationModel, error) {
	rows, err := db.Query("select id, value, date from quotations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotations []QuotationModel
	for rows.Next() {
		var quotation QuotationModel
		err = rows.Scan(&quotation.ID, &quotation.Value, &quotation.Date)
		if err != nil {
			return nil, err
		}
		quotations = append(quotations, quotation)
	}

	return quotations, nil
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
