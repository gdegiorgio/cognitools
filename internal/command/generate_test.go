package command

import (
	"bytes"
	"testing"

	"github.com/gdegiorgio/cognitools/internal/service"
	"github.com/gdegiorgio/cognitools/internal/ui"
)

func TestGenerate(t *testing.T) {

	buf := new(bytes.Buffer)

	generateCmd := NewGenerateCommand()
	generateCmd.SetOut(buf)

	generate(generateCmd, []string{}, &service.AwsMockService{}, &ui.PromptMock{}, &service.MockAuthService{})

	if buf.Len() == 0 {
		t.Error("Expected output, got none")
	}

	if !bytes.Contains(buf.Bytes(), []byte("✅ JWT generated successfully")) {
		t.Error("Expected output to contain '✅ JWT generated successfully'")
	}
}
