package users

import (
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Created_At  time.Time `json:"created_at"`
	Created_By  string    `json:"created_by"`
	Modified_At time.Time `json:"modified_at"`
	Modified_By string    `json:"modified_by"`
}
