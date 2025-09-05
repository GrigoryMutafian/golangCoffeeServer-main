package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func UpdateCoffeeMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method Not Allowed: ", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var updateCoffee map[string]cm.Coffee
	err := json.NewDecoder(r.Body).Decode(&updateCoffee)
	if err != nil {
		http.Error(w, "JSON decoding error: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	for key, newValue := range updateCoffee {
		if _, exists := cm.CoffeeDatabase[key]; exists {
			cm.CoffeeDatabase[key] = newValue
			cm.Update = key
		} else {
			http.Error(w, "Coffee not found", http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "Coffee " + cm.Update + " updated"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError)
	}
}
