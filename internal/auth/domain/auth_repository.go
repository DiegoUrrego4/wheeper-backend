package domain

type AuthRepository interface {
	Create()
	Update()
	Delete()
	GetByID()
	GetAll()
}
