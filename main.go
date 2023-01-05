package main

import (
  "github.com/c-bata/go-prompt"
  "github.com/gookit/color"
  "os"
  "os/exec"
  "fmt"
)

type TradeLog struct {
  EntryTime uint64
  ExitTime uint64
  EntryPrice float32
  ExitPrice float32
  PositionType bool
  EntryReason string
  ExitReason string
  PostAnalysis string
}

var logs []TradeLog

func StartPrompt(executorFunction func(string), completerFunction prompt.Completer, prefix string, prefixTextColor prompt.Color, inputTextColor prompt.Color) {
  p := prompt.New(
    executorFunction, 
    completerFunction, 
    prompt.OptionPrefix(prefix),
    prompt.OptionPrefixTextColor(prefixTextColor),
    prompt.OptionInputTextColor(inputTextColor),
    prompt.OptionSuggestionTextColor(prompt.Black),
  )
  p.Run()
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "create", Description: "Create new trade log"},
		{Text: "delete", Description: "Delete existing trade log"},
		{Text: "view", Description: "View existing trade log"},
		{Text: "edit", Description: "Edit existing trade log"},
		{Text: "list", Description: "Show all trade logs"},
		{Text: "export", Description: "Export trade logs to file"},
		{Text: "import", Description: "Import trade logs from file"},
		{Text: "exit", Description: "Exit program"},
		{Text: "help", Description: "Show commands"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(s string) {

  switch s {
  case "create":
    // Create new Trade Log
    var new_log = TradeLog{
      EntryTime:  289012890490,
      ExitTime:   349208902840,
      EntryPrice: 23121.23,
      ExitPrice:  25921.97,
      EntryReason: "Some entry reason",
      ExitReason: "Some exit reason",
    }

    logs = append(logs, new_log)

    fmt.Println(logs)
  case "delete":
  case "view":
  case "edit":
  case "list":
  case "export":
  case "import":
  case "exit":
    // Prompt user to exit
    var input = prompt.Input("Are you sure you want to exit? [Y/n] ", func(d prompt.Document) []prompt.Suggest {
      return prompt.FilterHasPrefix([]prompt.Suggest{
        {Text: "Y", Description: "Exit"},
        {Text: "n", Description: "Cancel"},
      }, d.GetWordBeforeCursor(), true)
    })

    // Exit program
    if input == "Y" {
      color.Yellow.Println("Bye! :)")
      os.Exit(0)
    }
  case "help":
  default:
    color.Println("Unknown command. Try \"help\" for more information.")
  }

}
func main() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()

  color.Println("Welcome to <cyan>TradeLogger</>, type <yellow>\"help\"</> to get started or press <yellow>\"Tab\"</> to see the commands.")

  StartPrompt(executor, completer, ">>> ", prompt.Cyan, prompt.White)
}
