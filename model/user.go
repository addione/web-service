package model

import "time"

const UserCollectionName = "users"

const UserTableName = "users"

const (
	STATUS_NEW      = "new"
	STATUS_ACTIVE   = "active"
	STATUS_INACTIVE = "inactive"
)

type User struct {
	Name        string  `json:"name"`
	ID          int64   `bson:"mysql_id"`
	Email       string  `json:"email"`
	Pass        string  `json:"password"`
	Balance     float64 `json:"balance"`
	PhoneNumber string  `bson:"phone_number"`
	Salutation  string
	Details     Details
	Address     string
	CreatedAt   time.Time `bson:"created_at"`
}

type Details struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}
