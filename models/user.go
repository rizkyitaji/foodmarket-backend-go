package models

type User struct {
	ID           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Email        string `db:"email" json:"email"`
	Password     string `db:"password" json:"password"`
	PhoneNumber  string `db:"phone_number" json:"phone_number"`
	Address      string `db:"address" json:"address"`
	ProfilePhoto string `db:"profile_photo" json:"profile_photo"`
	CreatedAt    string `db:"created_at" json:"created_at"`
	UpdatedAt    string `db:"updated_at" json:"updated_at"`
}
