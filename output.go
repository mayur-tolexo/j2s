package jsonToStruct

//Social model
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

//User model
type User struct {
	Age    float64 `json:"age"`
	Name   string  `json:"name"`
	Social Social  `json:"social"`
	Type   string  `json:"type"`
}
