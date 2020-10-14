package main

import (
	controller "WebServer/cmd/webserver/controllers"
	"fmt"
	"net/http"

	logs "github.com/sirupsen/logrus"
)

func main() {

	fmt.Printf("hello, world\n")
	http.HandleFunc("/", controller.Sayhello) // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil)  // устанавливаем порт веб-сервера
	if err != nil {
		logs.Fatal("ListenAndServe: ", err)
	}
}
