package main

import (
	"fmt"
	hn "golangCoffeeServer-main/HandlerModule"
	"net/http"
)

func main() {
	http.HandleFunc("/getAllCoffeeMenu", hn.GetAllCoffeeMenu)
	http.HandleFunc("/addNewCoffee", hn.AddNewCoffee)
	http.HandleFunc("/deleteCoffee", hn.DeleteCoffee)
	http.HandleFunc("/updateCoffeeMenu", hn.UpdateCoffeeMenu)
	http.HandleFunc("/checkStatusOfCoffee", hn.CheckStatusOfCoffee)
	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("сервер не запущен", err)
	}
}
