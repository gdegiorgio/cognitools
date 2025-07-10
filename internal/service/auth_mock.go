package service

type MockAuthService struct{}

func (m *MockAuthService) GenerateJWT(domain, clientId, secret, scope string) (string, error) {
	if domain == "" || clientId == "" || secret == "" || scope == "" {
		return "", nil
	}
	return "mock-jwt-token", nil
}
