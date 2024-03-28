package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
)

type Config struct {
	DBHost string `json:"db_host"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
	DBName string `json:"db_name"`
}

var config Config

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println(event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
					fmt.Println("modified file", event.Name)
					fmt.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error: ", err)
			}
		}
	}()
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done
}
