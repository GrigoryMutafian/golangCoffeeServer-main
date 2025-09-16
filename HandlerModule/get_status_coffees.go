package HandlerModule

import (
	"encoding/json"
	cm "golangCoffeeServer-main/coffeeModel"
	"golangCoffeeServer-main/db"
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

	response := make(map[cm.Status][]int)

	for _, value := range statusesArray {
		rows, err := db.DB.Query(`select id from coffees
		WHERE status = $1`, value)

		if err != nil {
			http.Error(w, "DB query error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var currentStatusId int
		currentStatusIdArray := []int{}

		for rows.Next() {
			err := rows.Scan(&currentStatusId)

			if err != nil {
				http.Error(w, "String processing error", http.StatusInternalServerError)
				return
			}

			currentStatusIdArray = append(currentStatusIdArray, currentStatusId)
		}

		response[value] = currentStatusIdArray
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError)
	}
}
