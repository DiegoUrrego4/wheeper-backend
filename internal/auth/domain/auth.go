package domain

type Auth struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	HashPassword string
}
