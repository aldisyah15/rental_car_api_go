package model

type RequestLogo struct {
	Name string `json:"name"`
	Logo string `json:"brand"`
}
type ResponseLogo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"brand"`
}
