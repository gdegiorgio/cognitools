package main

import "github.com/gdegiorgio/cognitools/internal/command"

func main() {
	cognitoolsCommand := command.NewCognitoolsCommand()
	cognitoolsCommand.Execute()
}
