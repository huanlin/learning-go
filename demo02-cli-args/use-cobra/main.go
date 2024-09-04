package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCommand *cobra.Command

func init() {
	helloCommand = &cobra.Command{
		Use:   "cli-cobra",
		Short: "Print hello world",
		Run:   sayHello,
	}
	helloCommand.Flags().StringP("name", "n", "World", "要跟誰說 hello。")
	helloCommand.MarkFlagRequired("name")
	helloCommand.Flags().StringP("language", "l", "en", "用哪一種語言說 hello。")
}

func sayHello(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	greeting := "Hello"
	language, _ := cmd.Flags().GetString("language")
	switch language {
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "fr":
		greeting = "Bonjour"
	case "zh":
		greeting = "哈囉"
	}
	fmt.Printf("%s %s!\n", greeting, name)
}

func main() {
	helloCommand.Execute()
}
