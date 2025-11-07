package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"

	"GO/app/updates"
)

func main() {
	godotenv.Load()
	http.HandleFunc("/webhook", httpHandler)
	http.ListenAndServe(":8000", nil)
}

func httpHandler(resp http.ResponseWriter, req *http.Request) {
	var update updates.Update

	// Читаем Body правильно
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	// Декодируем JSON
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, value := range update.Result {
		spew.Dump(value)
		// fmt.Println(value.Message)
	}
}
