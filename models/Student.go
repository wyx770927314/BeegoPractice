package models

type Student struct {
	Name  string   `json:"username" form:"username"`
	Age   int      `json:"password" form:"password"`
	Hobby []string `json:"hobby" form:"hobby"`
}
