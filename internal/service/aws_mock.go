package service

type AwsMockService struct{}

func (a *AwsMockService) ListPools() ([]CognitoPool, error) {
	return []CognitoPool{
		{Name: "Test Pool", PoolId: "us-east-1_123456789"},
	}, nil
}

func (a *AwsMockService) GetCognitoClientSecret(poolId, clientId string) (string, error) {
	return "test-secret", nil
}

func (a *AwsMockService) GetCognitoDomain(poolId string) (string, error) {
	return "test-domain", nil
}

func (a *AwsMockService) ListClients(poolId string) ([]CognitoClient, error) {
	return []CognitoClient{
		{Name: "Test Client", ClientId: "1234567890abcdef"}}, nil
}
