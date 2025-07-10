package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gdegiorgio/cognitools/internal/ui"
)

type AWS interface {
	ListPools() ([]CognitoPool, error)
	ListClients(poolId string) ([]CognitoClient, error)
	GetCognitoDomain(poolId string) (string, error)
	GetCognitoClientSecret(userPoolId, clientId string) (string, error)
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

type awsservice struct {
	client cognitoidentityprovider.Client
}

func NewAWSService() *awsservice {

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))

	if err != nil {
		panic(err)
	}

	client := cognitoidentityprovider.NewFromConfig(config)

	return &awsservice{
		client: *client,
	}
}
func (svc *awsservice) ListPools() ([]CognitoPool, error) {

	var pools []CognitoPool

	err := ui.WithSpinner("Retrieving Cognito Pools\n", func() error {
		maxResult := int32(50)
		var nextToken *string

		for {
			res, err := svc.client.ListUserPools(context.TODO(), &cognitoidentityprovider.ListUserPoolsInput{
				MaxResults: &maxResult,
				NextToken:  nextToken,
			})
			if err != nil {
				return fmt.Errorf("could not list cognito pools: %w", err)
			}

			for _, p := range res.UserPools {
				pools = append(pools, CognitoPool{
					PoolId: *p.Id,
					Name:   *p.Name,
				})
			}

			if res.NextToken == nil {
				break
			}
			nextToken = res.NextToken
		}
		return nil
	})

	return pools, err
}
func (svc *awsservice) ListClients(poolId string) ([]CognitoClient, error) {
	var clients []CognitoClient

	err := ui.WithSpinner("Retrieving Cognito Clients\n", func() error {
		maxResult := int32(50)
		var nextToken *string

		for {
			res, err := svc.client.ListUserPoolClients(context.TODO(), &cognitoidentityprovider.ListUserPoolClientsInput{
				UserPoolId: &poolId,
				MaxResults: &maxResult,
				NextToken:  nextToken,
			})
			if err != nil {
				return fmt.Errorf("could not list cognito clients: %w", err)
			}

			for _, c := range res.UserPoolClients {
				clients = append(clients, CognitoClient{
					ClientId: *c.ClientId,
					Name:     *c.ClientName,
				})
			}

			if res.NextToken == nil {
				break
			}
			nextToken = res.NextToken
		}
		return nil
	})

	return clients, err
}
func (svc *awsservice) GetCognitoDomain(poolId string) (string, error) {
	var domain string
	err := ui.WithSpinner("Retrieving Cognito Domain\n", func() error {
		res, err := svc.client.DescribeUserPool(context.TODO(), &cognitoidentityprovider.DescribeUserPoolInput{
			UserPoolId: &poolId,
		})
		if err != nil {
			return fmt.Errorf("could not describe user pool: %w", err)
		}

		domain = fmt.Sprintf("%s.auth.%s.amazoncognito.com", *res.UserPool.Domain, svc.client.Options().Region)
		return nil
	})

	return domain, err
}
func (svc *awsservice) GetCognitoClientSecret(userPoolId, clientId string) (string, error) {
	var secret string
	err := ui.WithSpinner("Retrieving Cognito Client Secret\n", func() error {
		res, err := svc.client.DescribeUserPoolClient(context.TODO(), &cognitoidentityprovider.DescribeUserPoolClientInput{
			UserPoolId: &userPoolId,
			ClientId:   &clientId,
		})
		if err != nil {
			return fmt.Errorf("could not describe user pool client: %w", err)
		}
		secret = *res.UserPoolClient.ClientSecret
		return nil
	})
	return secret, err
}
