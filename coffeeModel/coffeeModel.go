package coffeeModel

type Coffee struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Weight      float64
	RoastLevel  string
	Status      Status
}

type Status string

const (
	LowStatus    = "LowStatus"
	MediumStatus = "MediumStatus"
	HighStatus   = "HighStatus"
)

var CoffeeDatabase = map[string]Coffee{
	"coffee_id_1": {
		ID:          1,
		Name:        "Supremo",
		Description: "Best coffee from Colombia",
		Price:       120.0,
		Weight:      250.0,
		RoastLevel:  "Medium",
		Status:      LowStatus,
	},
	"coffee_id_2": {
		ID:          2,
		Name:        "kek",
		Description: "Best coffee from Russia",
		Price:       120.0,
		Weight:      250.0,
		RoastLevel:  "Medium",
		Status:      MediumStatus,
	},
}

var Update string
