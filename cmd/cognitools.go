package main

import "github.com/gdegiorgio/cognitools/internal/command"

func main() {
	root := command.NewRootCommand()
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}
