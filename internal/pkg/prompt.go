package pkg

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectUserPool(pools []string) (int, error) {

	prompt := promptui.Select{
		Label: "Select Cognito Pool",
		Items: pools,
	}

	idx, _, err := prompt.Run()

	if err != nil {
		return -1, fmt.Errorf("Could not select cognito pool : %v", err)
	}

	return idx, nil
}
