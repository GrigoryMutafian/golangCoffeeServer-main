package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type 

var coffeeMenu = map[string]string{
	"coffee": "cap",
}

var update string

func main() {
	http.HandleFunc("/getAllCoffeeMenu", getAllCoffeeMenu)
	http.HandleFunc("/addNewCoffee", addNewCoffee)
	http.HandleFunc("/deleteCoffee", deleteCoffee)
	http.HandleFunc("/updateCoffeeMenu", updateCoffeeMenu)
	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("сервер не запущен", err)
	}
}

func getAllCoffeeMenu(w http.ResponseWriter, r *http.Request) { //curl http://localhost:8080/getAllCoffeeMenu
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(coffeeMenu)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}
}

func addNewCoffee(w http.ResponseWriter, r *http.Request) { //-X POST -H "Content-Type: application/json" -d "{\"espresso\":\"strong\"}" http://localhost:8080/addNewCoffee
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var newCoffies map[string]string
	err := json.NewDecoder(r.Body).Decode(&newCoffies)
	if err != nil {
		http.Error(w, "Неверный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close() // прочесть

	for key, value := range newCoffies {
		coffeeMenu[key] = fmt.Sprintf("%v", value)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "кофе добавлен"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
	}

}

func deleteCoffee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Ключ не передан", http.StatusBadRequest)
		return
	}

	if _, ok := coffeeMenu[key]; !ok {
		http.Error(w, "Такого кофе нет", http.StatusNotFound)
		return
	}

	delete(coffeeMenu, key)
	fmt.Printf("Удалено кофе с ключом '%s', текущее меню: %+v\n", key, coffeeMenu)

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK) //
	response := map[string]string{"message": "Кофе удален", "key": key}
	err := json.NewEncoder(w).Encode(response) //
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func updateCoffeeMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var updateCoffee map[string]string
	err := json.NewDecoder(r.Body).Decode(&updateCoffee)
	if err != nil {
		http.Error(w, "Неверный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	for key, newValue := range updateCoffee {
		if _, exists := coffeeMenu[key]; exists {
			coffeeMenu[key] = newValue
			update = key
		} else {
			http.Error(w, "Такой кофе не найден", http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "Кофе " + update + " обновлён"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
	}
}
