package main

import (
	"fmt"
	"net/http"
)

func main() {
	startMyServer()
}

func startMyServer() {
	// Делаем переменную хранящую адрес
	address := "localhost:3000"

	// Создаем объект, отвечающий за маршрутизацию api
	myMux := http.NewServeMux()

	// Прописываем маршруты
	myMux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("это тестовый сервер Даши Лагутиной"))
	})

	myMux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is healthy"))
	})

	myMux.HandleFunc("GET /createOrder", CreateOrderHandler)
	myMux.HandleFunc("GET /completeOrder", CompleteOrderHandler)
	myMux.HandleFunc("GET /cancelOrder", CancelOrderHandler)

	// Запускаем сервер под адресу переменной adress с функцией обработчиком handler
	err := http.ListenAndServe(address, myMux)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateOrderHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("order created"))
}

func CompleteOrderHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("order completed"))
}

func CancelOrderHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("order canceled"))
}
