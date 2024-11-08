package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("name") {
		w.Write([]byte("Hello," + r.URL.Query().Get("name") + "!"))
	}
}

func main() {
	http.HandleFunc("/api/user", handler)

	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
