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

type Status int

const (
	LowStatus Status = iota
	MediumStatus
	HighStatus
)

var CoffeeDatabase = map[string]Coffee{
	"1": {
		ID:          1,
		Name:        "Supremo",
		Description: "Best coffee from Colombia",
		Price:       120.0,
		Weight:      250.0,
		RoastLevel:  "Medium",
		Status:      LowStatus,
	},
}

var Update string
