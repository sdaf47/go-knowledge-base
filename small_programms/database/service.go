package database

type entity interface{}

type DataBase interface {
	Create(entity) (int, error)
	Update(int, entity) error
	Delete(int, entity) error
	GetOne(int, entity) (interface{}, error)
	Get(e entity) (interface{}, error)
	Close() error
}
