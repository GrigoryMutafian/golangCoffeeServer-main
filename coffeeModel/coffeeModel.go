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

type CoffeeInput struct {
	ID          int     `json: id, omitempty`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	RoastLevel  string  `json:"roast_level"`
	Status      string  `json:"status"`
}

var Update string
