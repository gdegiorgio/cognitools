package ui

type PromptMock struct{}

func (p *PromptMock) SelectFromList(label string, items []string) (int, error) {
	if len(items) == 0 {
		return -1, nil
	}
	return 0, nil
}

func (p *PromptMock) PromptInput(label string) (string, error) {
	if label == "" {
		return "", nil
	}
	return "test-input", nil
}
