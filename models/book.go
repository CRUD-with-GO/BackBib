package models

type Book struct {
    ID    int   `gorm:"autoIncrement;primaryKey;unique;not null" json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    ISBN   string `json:"isbn"`
}