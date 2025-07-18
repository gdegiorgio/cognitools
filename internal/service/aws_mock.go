package service

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type AwsMockService struct{}

func (a *AwsMockService) DescribeUserPoolClient(userPoolID, clientID string) (types.UserPoolClientType, error) {
	clientName := "Test Client"
	return types.UserPoolClientType{
		ClientId:   &clientID,
		UserPoolId: &userPoolID,
		ClientName: &clientName,
	}, nil
}

func (a *AwsMockService) DescribeUserPool(poolID string) (types.UserPoolType, error) {
	name := "Test User Pool"
	status := types.StatusTypeEnabled
	return types.UserPoolType{
		Id:               &poolID,
		Name:             &name,
		Status:           status,
		CreationDate:     &time.Time{},
		LastModifiedDate: &time.Time{},
	}, nil
}

func (a *AwsMockService) ListUsersPools() ([]types.UserPoolDescriptionType, error) {
	name := "Test Pool"
	id := "us-east-1_123456789"
	return []types.UserPoolDescriptionType{
		{
			Name: &name,
			Id:   &id,
		},
	}, nil
}

func (a *AwsMockService) ListUserPoolClients(poolID string) ([]types.UserPoolClientDescription, error) {
	name := "Test Client"
	clientID := "1234567890abcdef"
	return []types.UserPoolClientDescription{
		{
			ClientId:   &clientID,
			ClientName: &name,
		},
	}, nil
}

func (a *AwsMockService) GetCognitoHost(domain string) string {
	return "https://" + domain + ".auth.us-east-1.amazoncognito.com"
}
