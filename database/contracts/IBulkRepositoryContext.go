package contracts

type IBulkRepositoryContext interface{
	GetAll() (interface{}, error)
	GetId(string) (interface{}, error)
	DeleteById(string) error
	Insert(interface{}) error
	Save(interface{},string) error
}