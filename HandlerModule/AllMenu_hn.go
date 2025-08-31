package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func GetAllCoffeeMenu(w http.ResponseWriter, r *http.Request) { //curl http://localhost:8080/getAllCoffeeMenu
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cm.CoffeeDatabase)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}
}
