package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func CheckStatusOfCoffee(w http.ResponseWriter, r *http.Request) {
	notEnoughCoffee := make(map[string]cm.Status)

	for key := range cm.CoffeeDatabase {
		if cm.CoffeeDatabase[key].Status == cm.LowStatus {
			notEnoughCoffee[key] = cm.CoffeeDatabase[key].Status
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(notEnoughCoffee)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
	}
}
