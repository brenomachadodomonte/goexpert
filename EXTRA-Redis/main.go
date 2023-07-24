package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

var ErrEmpty string = "redis: nil"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "secret",
		DB:       0,
	})

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gomysql")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client.Set(ctx, "nome", "Breno", 50*time.Second)

	value, err := client.Get(ctx, "nome").Result()
	if err != nil {
		if err.Error() == ErrEmpty {
			println("Variável vazia")
		}
		panic(err)
	}

	client.Del(ctx, "nome")
	fmt.Println(value)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			people, err := selectAllPeople(db)
			if err != nil {
				writer.WriteHeader(http.StatusNotFound)
				return
			}
			json.NewEncoder(writer).Encode(people)
		case http.MethodPost:
			var createPersonInput CreatePersonInput
			err := json.NewDecoder(request.Body).Decode(&createPersonInput)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}

			person := NewPerson(createPersonInput.Name, createPersonInput.Age)
			err = insertPeople(db, person)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
			}

			writer.WriteHeader(http.StatusCreated)
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err)
	}
}

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPerson(name string, age int) *Person {
	return &Person{
		ID:   uuid.New().String(),
		Name: name,
		Age:  age,
	}
}

type CreatePersonInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func selectAllPeople(db *sql.DB) ([]Person, error) {
	rows, err := db.Query("select id, name, age from people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	people := make([]Person, 0)
	for rows.Next() {
		var person Person
		err = rows.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return people, nil
}

func insertPeople(db *sql.DB, person *Person) error {
	stmt, err := db.Prepare("insert into people (id, name, age) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(person.ID, person.Name, person.Age)
	if err != nil {
		return err
	}
	return nil
}
