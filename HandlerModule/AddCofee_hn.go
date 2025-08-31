package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func AddNewCoffee(w http.ResponseWriter, r *http.Request) { //-X POST -H "Content-Type: application/json" -d "{\"espresso\":\"strong\"}" http://localhost:8080/addNewCoffee
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var newCoffies map[string]cm.Coffee
	err := json.NewDecoder(r.Body).Decode(&newCoffies)
	if err != nil {
		http.Error(w, "Неверный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close() // прочесть

	for key, value := range newCoffies {
		cm.CoffeeDatabase[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "кофе добавлен"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
	}

}
