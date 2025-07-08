package pkg

import (
	"time"

	"github.com/briandowns/spinner"
)

func NewSpinner() *spinner.Spinner {
	return spinner.New(spinner.CharSets[39], time.Minute * 1)
}
