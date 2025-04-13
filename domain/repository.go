package domain

type PropertyRepository interface {
	Save(property Property) error
	FindAll() ([]Property, error)
}
