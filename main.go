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
	http.HandleFunc("/GetStatusCoffees", hn.GetStatusCoffees)
	fmt.Println("server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("the server is not running", err)
	}
}
