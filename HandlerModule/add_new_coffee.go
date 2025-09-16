package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"golangCoffeeServer-main/db"
	"net/http"
)

func AddNewCoffee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed: ", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var newCoffee cm.CoffeeInput

	err := json.NewDecoder(r.Body).Decode(&newCoffee)

	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close() //

	result, err := db.DB.Exec(
		`INSERT INTO coffees (name, description, price, weight, roast_level, status)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		newCoffee.Name, newCoffee.Description, newCoffee.Price, newCoffee.Weight, newCoffee.RoastLevel, newCoffee.Status,
	)

	if err != nil {
		http.Error(w, "Database insertion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No rows inserted", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "coffee added", "coffeeName": newCoffee.Name}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError) //
	}

}
