package jsonToStruct

//Batter model
type Batter struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

//Batters model
type Batters struct {
	Batter []Batter `json:"batter"`
}

//Topping model
type Topping struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

//User model
type User struct {
	Batters Batters   `json:"batters"`
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Ppu     float64   `json:"ppu"`
	Topping []Topping `json:"topping"`
	Type    string    `json:"type"`
}
