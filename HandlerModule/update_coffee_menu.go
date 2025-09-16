package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"golangCoffeeServer-main/db"
	"net/http"
)

func UpdateCoffeeMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method Not Allowed: ", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var updateCoffee cm.CoffeeInput
	err := json.NewDecoder(r.Body).Decode(&updateCoffee)
	if err != nil {
		http.Error(w, "JSON decoding error: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	result, err := db.DB.Exec(`UPDATE coffees 
	                           SET name = $1, description = $2, price = $3, weight = $4, roast_level = $5, status = $6
	                           WHERE id = $7`,
		updateCoffee.Name, updateCoffee.Description, updateCoffee.Price, updateCoffee.Weight, updateCoffee.RoastLevel, updateCoffee.Status, updateCoffee.ID)

	if err != nil {
		http.Error(w, "Database manipulating error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No rows updated", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Coffee " + updateCoffee.Name + " updated"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
