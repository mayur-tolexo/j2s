package jsonToStruct

//User model
type User struct {
	Age    float64                `json:"age"`
	Name   string                 `json:"name"`
	Social map[string]interface{} `json:"social"`
	Type   string                 `json:"type"`
}
