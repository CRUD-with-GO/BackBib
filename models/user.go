package models

type User struct {
    ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
	Loans []Loan `gorm:"foreignKey:UserID"`
}

