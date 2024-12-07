package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Добавляем заголовки CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Разрешить доступ всем источникам (можно указать конкретный домен вместо "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешённые методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Разрешённые заголовки

	// Обрабатываем OPTIONS-запросы
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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
