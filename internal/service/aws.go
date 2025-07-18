package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gdegiorgio/cognitools/internal/ui"
)

type AWS interface {
	DescribeUserPoolClient(userPoolID, clientID string) (types.UserPoolClientType, error)
	DescribeUserPool(poolID string) (types.UserPoolType, error)
	GetCognitoHost(domain string) string
	ListUsersPools() ([]types.UserPoolDescriptionType, error)
	ListUserPoolClients(poolID string) ([]types.UserPoolClientDescription, error)
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

func (svc *awsservice) DescribeUserPool(poolID string) (types.UserPoolType, error) {
	var result *cognitoidentityprovider.DescribeUserPoolOutput

	err := ui.WithSpinner("Describing Cognito User Pool\n", func() error {
		res, err := svc.client.DescribeUserPool(context.TODO(), &cognitoidentityprovider.DescribeUserPoolInput{
			UserPoolId: &poolID,
		})
		if err != nil {
			return fmt.Errorf("could not describe user pool: %w", err)
		}

		result = res
		return nil
	})

	return *result.UserPool, err
}

func (svc *awsservice) DescribeUserPoolClient(userPoolID, clientID string) (types.UserPoolClientType, error) {
	var result *cognitoidentityprovider.DescribeUserPoolClientOutput

	err := ui.WithSpinner("Describing Cognito User Pool Client\n", func() error {
		res, err := svc.client.DescribeUserPoolClient(context.TODO(), &cognitoidentityprovider.DescribeUserPoolClientInput{
			UserPoolId: &userPoolID,
			ClientId:   &clientID,
		})
		if err != nil {
			return fmt.Errorf("could not describe user pool client: %w", err)
		}

		result = res
		return nil
	})

	return *result.UserPoolClient, err
}

func (svc *awsservice) ListUsersPools() ([]types.UserPoolDescriptionType, error) {

	var pools []types.UserPoolDescriptionType

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

			pools = append(pools, res.UserPools...)

			if res.NextToken == nil {
				break
			}
			nextToken = res.NextToken
		}
		return nil
	})

	return pools, err
}

func (svc *awsservice) ListUserPoolClients(poolID string) ([]types.UserPoolClientDescription, error) {
	var clients []types.UserPoolClientDescription

	err := ui.WithSpinner("Retrieving Cognito Clients\n", func() error {
		maxResult := int32(50)
		var nextToken *string

		for {
			res, err := svc.client.ListUserPoolClients(context.TODO(), &cognitoidentityprovider.ListUserPoolClientsInput{
				UserPoolId: &poolID,
				MaxResults: &maxResult,
				NextToken:  nextToken,
			})
			if err != nil {
				return fmt.Errorf("could not list cognito clients: %w", err)
			}

			clients = append(clients, res.UserPoolClients...)

			if res.NextToken == nil {
				break
			}
			nextToken = res.NextToken
		}
		return nil
	})

	return clients, err
}

func (svc *awsservice) GetCognitoHost(domain string) string {
	return domain + ".auth." + svc.client.Options().Region + ".amazoncognito.com"
}
