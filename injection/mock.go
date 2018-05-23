package injection

type MockStore struct{}

func (m *MockStore) GetSomething() (string, error) {
	return "Hello World!!", nil
}
