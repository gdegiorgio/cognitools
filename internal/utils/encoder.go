package utils

import (
	"encoding/json"
	"fmt"
)

func FormatJSON(input interface{}) (string, error) {
	data, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}
	return string(data), nil
}
