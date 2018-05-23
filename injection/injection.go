package injection

type StoreMode string

const (
	MOCK StoreMode = "MOCK"
	PROD StoreMode = "PROD"
)

type Injection struct {
	Store Store
}

func GetStore(store StoreMode) Store {
	switch store {
	case MOCK:
		return &MockStore{}
	case PROD:
		return &ProdStore{}
	}
	return nil
}

type Store interface {
	GetSomething() (string, error)
}
