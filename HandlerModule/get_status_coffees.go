package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func GetStatusCoffees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed: ", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	statusesArray := []cm.Status{}
	err := json.NewDecoder(r.Body).Decode(&statusesArray)
	if err != nil {
		http.Error(w, "JSON decoding error:: "+err.Error(), http.StatusInternalServerError)
	}

	response := make(map[string]cm.Status)

	for _, value := range statusesArray {
		for key := range cm.CoffeeDatabase {
			if cm.CoffeeDatabase[key].Status == value {
				response[key] = cm.CoffeeDatabase[key].Status
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError)
	}
}
