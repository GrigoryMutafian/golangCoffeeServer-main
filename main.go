package main

import (
	"fmt"
	hn "golangCoffeeServer-main/HandlerModule"
	"golangCoffeeServer-main/db"
	"net/http"
)

func main() {
	if err := db.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	defer db.DB.Close()

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
