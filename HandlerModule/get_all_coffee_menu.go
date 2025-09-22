package HandlerModule

import (
	"database/sql"
	"encoding/json"
	"golangCoffeeServer-main/db"
	"net/http"
)

func GetAllCoffeeMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query("SELECT id, name, description, price, weight, roast_level, status FROM coffees")
	if err != nil {
		http.Error(w, "DB query error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var coffees []struct {
		ID          int
		Name        string
		Description sql.NullString
		Price       sql.NullFloat64
		Weight      sql.NullFloat64
		RoastLevel  string
		Status      string
	}

	for rows.Next() {
		var c struct {
			ID          int
			Name        string
			Description sql.NullString
			Price       sql.NullFloat64
			Weight      sql.NullFloat64
			RoastLevel  string
			Status      string
		}

		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.Price, &c.Weight, &c.RoastLevel, &c.Status)
		if err != nil {
			http.Error(w, "String processing error", http.StatusInternalServerError)
			return
		}
		coffees = append(coffees, c)
	}

	err = rows.Err()
	if err != nil {
		http.Error(w, "Error as a result of the request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(coffees)
	if err != nil {
		http.Error(w, "JSON encoding error: ", http.StatusInternalServerError)
		return
	}
}
