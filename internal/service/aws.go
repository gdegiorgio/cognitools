package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gdegiorgio/cognitools/internal/pkg"
)

type Service interface {
	ListPools() []CognitoPool
}

type CognitoPool struct {
	PoolId string
	Name   string
}

type CognitoScope struct {
	ScopeId string
	Name    string
}

type CognitoClient struct {
	ClientId string
	Name     string
}

type AWSSevice struct {
	client cognitoidentityprovider.Client
}

func NewAWSService() *AWSSevice {

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))

	if err != nil {
		panic(err)
	}

	client := cognitoidentityprovider.NewFromConfig(config)

	return &AWSSevice{
		client: *client,
	}
}

func (svc *AWSSevice) ListPools() ([]CognitoPool, error) {

	spinner := pkg.NewSpinner()
	spinner.Suffix = "Retrieving Cognito Pools\n"
	spinner.Start()
	defer spinner.Stop()

	maxResult := int32(50)

	var nextToken *string = nil

	for {

		params := cognitoidentityprovider.ListUserPoolsInput{
			MaxResults: &maxResult,
			NextToken:  nextToken,
		}

		res, err := svc.client.ListUserPools(context.TODO(), &params)

		if err != nil {
			return nil, fmt.Errorf("could not list cognito pools: %v", err)
		}

		poolsId := []CognitoPool{}
		for _, p := range res.UserPools {

			pool := &CognitoPool{
				PoolId: *p.Id,
				Name:   *p.Name,
			}
			poolsId = append(poolsId, *pool)
		}

		if res.NextToken == nil {
			return poolsId, nil
		}
		nextToken = res.NextToken
	}
}

func (svc *AWSSevice) ListClients(poolId string) ([]CognitoClient, error) {

	spinner := pkg.NewSpinner()
	spinner.Suffix = "Retrieving Cognito Client\n"
	spinner.Start()
	defer spinner.Stop()

	maxResult := int32(50)

	var nextToken *string = nil

	for {

		params := cognitoidentityprovider.ListUserPoolClientsInput{
			UserPoolId: &poolId,
			MaxResults: &maxResult,
			NextToken:  nextToken,
		}

		res, err := svc.client.ListUserPoolClients(context.TODO(), &params)

		if err != nil {
			return nil, fmt.Errorf("could not list cognito clients: %v", err)
		}

		clientsId := []CognitoClient{}
		for _, c := range res.UserPoolClients {

			client := &CognitoClient{
				ClientId: *c.ClientId,
				Name:     *c.ClientName,
			}
			clientsId = append(clientsId, *client)
		}

		if res.NextToken == nil {
			return clientsId, nil
		}
		nextToken = res.NextToken
	}
}

func (svc *AWSSevice) ListScopes(clientId string) ([]CognitoScope, error) {

	spinner := pkg.NewSpinner()
	spinner.Suffix = "Retrieving Cognito Scopes\n"
	spinner.Start()
	defer spinner.Stop()

	return []CognitoScope{}, nil
}

func (svc AWSSevice) GetCognitoDomain(poolId string) (string, error) {

	spinner := pkg.NewSpinner()
	spinner.Suffix = "Retrieving Cognito Domain\n"
	spinner.Start()
	defer spinner.Stop()

	params := cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: &poolId,
	}

	res, err := svc.client.DescribeUserPool(context.TODO(), &params)

	if err != nil {
		return "", fmt.Errorf("could not describe user pool: %v", err)
	}

	domain := fmt.Sprintf("%s.auth.%s.amazoncognito.com", *res.UserPool.Domain, svc.client.Options().Region)

	return domain, nil
}

func (svc *AWSSevice) GetCognitoClientSecret(userPoolId string, clientId string) (string, error) {

	spinner := pkg.NewSpinner()
	spinner.Suffix = "Retrieving Cognito Client Secret\n"
	spinner.Start()
	defer spinner.Stop()

	params := cognitoidentityprovider.DescribeUserPoolClientInput{
		UserPoolId: &userPoolId,
		ClientId:   &clientId,
	}

	res, err := svc.client.DescribeUserPoolClient(context.TODO(), &params)

	if err != nil {
		return "", fmt.Errorf("could not describe user pool client: %v", err)
	}

	return *res.UserPoolClient.ClientSecret, nil
}
