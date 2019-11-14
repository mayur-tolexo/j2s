package jsonToStruct

//User model
type User struct {
	Age    float64 `json:"age"`
	Social Social  `json:"social"`
	Test1  Test1   `json:"test1"`
}

//Test1 model
type Test1 struct {
	Social Test1Social `json:"social"`
}

//Test1Social model
type Test1Social struct {
	Test2 Test2 `json:"test2"`
}

//Test2 model
type Test2 struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

//Social model
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}
