package pkg

import (
	"time"

	"github.com/briandowns/spinner"
)

func WithSpinner(suffix string, fn func() error) error {
	spinner := spinner.New(spinner.CharSets[39], time.Minute*1)
	spinner.Suffix = suffix
	spinner.Start()
	defer spinner.Stop()
	return fn()
}
