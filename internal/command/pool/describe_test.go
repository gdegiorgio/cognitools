package pool

import (
	"bytes"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/spf13/cobra"
)

func TestDescribe(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		mockSetup      func() service.AWS
		expectedOutput string
		expectedError  bool
	}{
		{
			name: "successful describe",
			args: []string{"us-east-1_123456789"},
			mockSetup: func() service.AWS {
				return &service.AwsMockService{}
			},
			expectedOutput: "🏖️ User Pool Details:",
			expectedError:  false,
		},
		{
			name: "successful describe with JSON output",
			args: []string{"us-east-1_123456789"},
			mockSetup: func() service.AWS {
				return &service.AwsMockService{}
			},
			expectedOutput: "\"Name\": \"Test User Pool\"",
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd := &cobra.Command{}
			cmd.SetOut(buf)

			// Set JSON flag for JSON test
			if tt.name == "successful describe with JSON output" {
				outputJSON = true
			} else {
				outputJSON = false
			}

			err := describe(cmd, tt.args, tt.mockSetup())

			if tt.expectedError && err == nil {
				t.Error("Expected error, got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}

			output := buf.String()
			if !bytes.Contains(buf.Bytes(), []byte(tt.expectedOutput)) {
				t.Errorf("Expected output to contain '%s', got: %s", tt.expectedOutput, output)
			}
		})
	}
}

// Mock service for error testing
type mockAWSServiceWithError struct{}

func (m *mockAWSServiceWithError) DescribeUserPool(poolID string) (types.UserPoolType, error) {
	return types.UserPoolType{}, &types.ResourceNotFoundException{
		Message: stringPtr("User pool not found"),
	}
}

func (m *mockAWSServiceWithError) DescribeUserPoolClient(userPoolID, clientID string) (types.UserPoolClientType, error) {
	return types.UserPoolClientType{}, nil
}

func (m *mockAWSServiceWithError) ListUsersPools() ([]types.UserPoolDescriptionType, error) {
	return nil, nil
}

func (m *mockAWSServiceWithError) ListUserPoolClients(poolID string) ([]types.UserPoolClientDescription, error) {
	return nil, nil
}

func (m *mockAWSServiceWithError) GetCognitoHost(domain string) string {
	return ""
}

func TestDescribeError(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := &cobra.Command{}
	cmd.SetOut(buf)

	err := describe(cmd, []string{"invalid-pool"}, &mockAWSServiceWithError{})

	if err == nil {
		t.Error("Expected error, got none")
	}

	output := buf.String()
	if !bytes.Contains(buf.Bytes(), []byte("❌")) {
		t.Errorf("Expected error output with ❌, got: %s", output)
	}
}

func stringPtr(s string) *string {
	return &s
}