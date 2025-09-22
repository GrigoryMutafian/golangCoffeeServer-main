package HandlerModule

import (
	"encoding/json"
	"fmt"
	"golangCoffeeServer-main/db"
	"net/http"
)

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed: ", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var coffeeRemover int

	err := json.NewDecoder(r.Body).Decode(&coffeeRemover)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("DELETE FROM coffees WHERE id = $1", coffeeRemover)

	if err != nil {
		http.Error(w, "Database manipulating error: "+err.Error(), http.StatusInternalServerError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error checking rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No rows inserted", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	removedCoffee := fmt.Sprintf("Removed Coffee id: %d", coffeeRemover)
	err = json.NewEncoder(w).Encode(removedCoffee)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
