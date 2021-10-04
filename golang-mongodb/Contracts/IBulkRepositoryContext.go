package contracts

type IBulkRepositoryContext interface {
	GetAll() (interface{}, error)
	ReadById(string) (interface{}, error)
	DeleteById(string) error
	Insert(interface{}) error
	Save(interface{}, string) error
}
