package command

import "testing"

func TestRootCommandWithNoFlags(t *testing.T) {
	cmd := NewRootCommand()
	// It should not return an error when executed without flags but print help instead
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Expected command to execute without error, got: %v", err)
	}
}
