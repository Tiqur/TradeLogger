package main

import (
  "github.com/c-bata/go-prompt"
  "github.com/gookit/color"
  "os"
  "os/exec"
  "fmt"
  "strconv"
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

func ShortPrompt(prefix string) string {
    return prompt.Input(prefix, func(d prompt.Document) []prompt.Suggest {return prompt.FilterHasPrefix([]prompt.Suggest{}, d.GetWordBeforeCursor(), true)})
}

func executor(s string) {

  switch s {
  case "create":
    var string = ""

    PositionType:
    string = prompt.Input("Position Type: ", func(d prompt.Document) []prompt.Suggest {return prompt.FilterHasPrefix([]prompt.Suggest{
      {Text: "long", Description: "Long position"},
      {Text: "short", Description: "Short position"},
    }, d.GetWordBeforeCursor(), true)})
    if string != "long" && string != "short" {
      fmt.Println("Invalid position type ( please type \"long\" or \"short\" )")
      goto PositionType
    }
    var position_type = string == "long"

    EntryTime:
    string = ShortPrompt("Entry Time: ")
    entry_time, err := strconv.ParseUint(string, 10, 64)
    if err != nil {
      fmt.Println("Invalid entry time. Please input an unsigned integer")
      goto EntryTime
    }

    ExitTime:
    string = ShortPrompt("Exit Time: ")
    exit_time, err := strconv.ParseUint(string, 10, 64)
    if err != nil {
      fmt.Println("Invalid exit time. Please input an unsigned integer")
      goto ExitTime
    }

    EntryPrice:
    string = ShortPrompt("Entry Price: ")
    entry_price, err := strconv.ParseFloat(string, 32)
    if err != nil {
      fmt.Println("Invalid entry price. Please input a float")
      goto EntryPrice
    }

    ExitPrice:
    string = ShortPrompt("Exit Price: ")
    exit_price, err := strconv.ParseFloat(string, 32)
    if err != nil {
      fmt.Println("Invalid exit price. Please input a float")
      goto ExitPrice
    }

    // Entry Reason
    var entry_reason = ShortPrompt("Entry Reason: ")

    // Exit Reason
    var exit_reason = ShortPrompt("Exit Reason: ")

    // Post Analysis
    var post_analysis = ShortPrompt("Post Analysis: ")

    // Create new Trade Log
    var new_log = TradeLog{
      PositionType:  position_type,
      EntryTime:  entry_time,
      ExitTime:   exit_time,
      EntryPrice: float32(entry_price),
      ExitPrice:  float32(exit_price),
      EntryReason: entry_reason,
      ExitReason: exit_reason,
      PostAnalysis: post_analysis,
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
