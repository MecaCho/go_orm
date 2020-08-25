package model

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data []byte `json:"data" gorm:"type:blob"`
}
