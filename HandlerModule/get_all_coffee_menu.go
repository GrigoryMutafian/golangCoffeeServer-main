package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func GetAllCoffeeMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cm.CoffeeDatabase)
	if err != nil {
		http.Error(w, "JSON encoding error: ", http.StatusInternalServerError)
		return
	}
}
