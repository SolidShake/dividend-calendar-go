package models

//User is struct of basic
type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
