package pkg

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectFromList(label string, items []string) (int, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	idx, _, err := prompt.Run()
	if err != nil {
		return -1, fmt.Errorf("failed to select %s: %w", label, err)
	}
	return idx, nil
}

func PromptInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}
	val, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to input %s: %w", label, err)
	}
	return val, nil
}
