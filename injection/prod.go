package injection

type ProdStore struct{}

func (m *ProdStore) GetSomething() (string, error) {
	return "PRODUCTION CONTENT!!", nil
}
