package account

type Account struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    int64  `json:"created_at"`
}
