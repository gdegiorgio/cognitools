package clients

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// formatAWSError converts AWS errors to user-friendly messages
func formatAWSError(err error) string {
	if err == nil {
		return ""
	}

	errStr := err.Error()
	
	// Check for specific AWS error types
	var resourceNotFound *types.ResourceNotFoundException
	var invalidParam *types.InvalidParameterException
	var notAuthorized *types.NotAuthorizedException
	
	if errors.As(err, &resourceNotFound) {
		if strings.Contains(errStr, "pool") {
			return "User pool not found. Please check the pool ID and try again."
		}
		if strings.Contains(errStr, "client") {
			return "User pool client not found. Please check the client ID and try again."
		}
		return "Resource not found. Please check your parameters and try again."
	}
	
	if errors.As(err, &invalidParam) {
		return "Invalid parameter provided. Please check your input and try again."
	}
	
	if errors.As(err, &notAuthorized) {
		return "Not authorized to perform this operation. Please check your AWS credentials and permissions."
	}
	
	// For other errors, return a generic message with the original error
	return errStr
}