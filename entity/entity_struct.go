package entity

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Age      uint8  `json:"age"`
}
