package jsonToStruct

//User model
type User struct {
	Batters Batters	`json:"batters"`
	Id string	`json:"id"`
	Name string	`json:"name"`
	Ppu float64	`json:"ppu"`
	Topping []interface {}	`json:"topping"`
	Type string	`json:"type"`
}