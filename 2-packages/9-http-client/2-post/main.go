package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Breno"}`))
	resp, err := client.Post("https://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
