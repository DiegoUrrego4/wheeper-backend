package domain

type AuthToken interface {
	GenerateToken()
	ValidateToken()
}
