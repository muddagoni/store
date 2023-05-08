package model

type Service interface {
	Add()
	Update()
	Delete()
	Get()
}

type Repository interface {
	Add()
	Update()
	Delete()
	Get()
}
