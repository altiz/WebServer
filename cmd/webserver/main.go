package main

import (
	controller "WebServer/cmd/webserver/controllers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/google/glog"
)

func main() {

	fmt.Printf("hello, world\n")
	http.HandleFunc("/", controller.Sayhello) // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil)  // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
