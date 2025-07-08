package pkg

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectUserPool(pools []string) (int, error) {

	prompt := promptui.Select{
		Label: "🏖️ Select Cognito Pool",
		Items: pools,
	}

	idx, _, err := prompt.Run()

	if err != nil {
		return -1, fmt.Errorf("could not select cognito pool : %v", err)
	}

	return idx, nil
}

func SelectClients(clients []string) (int, error) {

	prompt := promptui.Select{
		Label: "👤 Select Client",
		Items: clients,
	}

	idx, _, err := prompt.Run()

	if err != nil {
		return -1, fmt.Errorf("could not select cognito pool : %v", err)
	}

	return idx, nil
}

func SelectScope(scopes []string) (string, error) {

	prompt := promptui.Prompt{
		Label: "🎯 Please write the Cognito Scope Name to use for request",
	}

	scope, err := prompt.Run()

	if err != nil {
		return "", fmt.Errorf("could not select cognito scope : %v", err)
	}

	return scope, nil
}
