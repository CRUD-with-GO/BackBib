package models




type Loan struct {
    ID         int      	`gorm:"primaryKey;autoIncrement" json:"id"`
    BookID     int      	`gorm:"column:book_id;foreignKey:BookID" json:"book_id"`
    UserID     int      	`gorm:"column:user_id;foreignKey:UserID" json:"user_id"`
    LoanDate   string		`gorm:"column:loan_date" json:"loan_date"`
    ReturnDate string		`gorm:"column:return_date"  json:"return_date"`
    Book       Book 		`gorm:"foreignKey:BookID;references:ID"`
    User       User 		`gorm:"foreignKey:UserID;references:ID"`
}