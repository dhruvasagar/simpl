package main

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/dhruvasagar/pay-later-go/cmd"
)

func main() {
	var completer = readline.NewPrefixCompleter(
		readline.PcItem("new",
			readline.PcItem("user"),
			readline.PcItem("merchant"),
			readline.PcItem("txn"),
		),
		readline.PcItem("update",
			readline.PcItem("user"),
			readline.PcItem("merchant"),
		),
		readline.PcItem("report",
			readline.PcItem("discount"),
			readline.PcItem("dues"),
			readline.PcItem("users-at-credit-limit"),
			readline.PcItem("total-dues"),
		),
		readline.PcItem("payback"),
	)
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "> ",
		EOFPrompt:         "exit",
		AutoComplete:      completer,
		InterruptPrompt:   "^D",
		HistorySearchFold: true,
	})

	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		cmdName, _ := l.Readline()
		cmdName = strings.TrimSpace(cmdName)
		if cmdName == "" || cmdName == "exit" {
			return
		}
		cmd.Execute(cmdName)
		fmt.Println()
	}
}
