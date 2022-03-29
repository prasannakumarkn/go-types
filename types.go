package types

type Myevent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"What is your age?"`
}

type Response struct {
	Message string `json:"Answer"`
}
