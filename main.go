package main

import (
  "github.com/c-bata/go-prompt"
  "github.com/gookit/color"
  "os"
  "os/exec"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "create", Description: "Create new trade log"},
		{Text: "delete", Description: "Delete existing trade log"},
		{Text: "view", Description: "View existing trade log"},
		{Text: "list", Description: "Show all trade logs"},
		{Text: "export", Description: "Export trade logs to file"},
		{Text: "exit", Description: "Exit program"},
		{Text: "help", Description: "Show commands"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(s string) {

}

func main() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()

  color.Println("Welcome to <cyan>TradeLogger</>, type <yellow>\"help\"</> to get started or press <yellow>\"Tab\"</> to see the commands.")
  defer color.Yellow.Println("Bye! :)")

  p := prompt.New(
    executor, 
    completer, 
    prompt.OptionPrefix(">>> "),
    prompt.OptionPrefixTextColor(prompt.Cyan),
    prompt.OptionSuggestionTextColor(prompt.Black),
  )
  p.Run()
  
}
