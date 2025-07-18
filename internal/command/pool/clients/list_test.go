package clients

import (
	"bytes"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/spf13/cobra"
)

func TestList(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		mockSetup      func() service.AWS
		expectedOutput string
		expectedError  bool
	}{
		{
			name: "successful list",
			args: []string{"us-east-1_123456789"},
			mockSetup: func() service.AWS {
				return &service.AwsMockService{}
			},
			expectedOutput: "👤 User Pool Clients:",
			expectedError:  false,
		},
		{
			name: "empty client list",
			args: []string{"us-east-1_123456789"},
			mockSetup: func() service.AWS {
				return &mockAWSServiceEmpty{}
			},
			expectedOutput: "❌ No clients found for this user pool.",
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd := &cobra.Command{}
			cmd.SetOut(buf)

			err := list(cmd, tt.args, tt.mockSetup())

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

// Mock service for empty results
type mockAWSServiceEmpty struct{}

func (m *mockAWSServiceEmpty) DescribeUserPool(poolID string) (types.UserPoolType, error) {
	return types.UserPoolType{}, nil
}

func (m *mockAWSServiceEmpty) DescribeUserPoolClient(userPoolID, clientID string) (types.UserPoolClientType, error) {
	return types.UserPoolClientType{}, nil
}

func (m *mockAWSServiceEmpty) ListUsersPools() ([]types.UserPoolDescriptionType, error) {
	return []types.UserPoolDescriptionType{}, nil
}

func (m *mockAWSServiceEmpty) ListUserPoolClients(poolID string) ([]types.UserPoolClientDescription, error) {
	return []types.UserPoolClientDescription{}, nil
}

func (m *mockAWSServiceEmpty) GetCognitoHost(domain string) string {
	return ""
}