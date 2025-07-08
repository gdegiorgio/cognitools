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

	params := cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: &maxResult,
	}

	res, err := svc.client.ListUserPools(context.TODO(), &params)

	if err != nil {
		return nil, fmt.Errorf("Could not list cognito pools: %v", err)
	}

	poolsId := []CognitoPool{}

	for {
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
	}
}
