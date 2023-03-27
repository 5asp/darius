package account

import "database/sql"

type Account struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    int64  `json:"created_at"`
}

func NewAccount(db *sql.DB) {

}

func Find() {

}

func Create() {

}

func Delete() {

}

func Edit() {

}
