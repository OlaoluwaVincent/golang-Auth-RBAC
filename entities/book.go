package entities

type Book struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
