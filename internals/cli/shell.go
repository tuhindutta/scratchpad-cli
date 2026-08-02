package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func StartShell(root *cobra.Command) error {

	fmt.Print(`
                                                                  
 ▄▄▄▄▄▄▄                               ▄▄    ▄▄▄▄▄▄▄           ▄▄ 
█████▀▀▀                    ██         ██    ███▀▀███▄         ██ 
 ▀████▄  ▄████ ████▄  ▀▀█▄ ▀██▀▀ ▄████ ████▄ ███▄▄███▀ ▀▀█▄ ▄████ 
   ▀████ ██    ██ ▀▀ ▄█▀██  ██   ██    ██ ██ ███▀▀▀▀  ▄█▀██ ██ ██ 
███████▀ ▀████ ██    ▀█▄██  ██   ▀████ ██ ██ ███      ▀█▄██ ▀████ 
                                                                  
                                                                  
╭──────────────────────────────────────────────────────────────╮
│             🚀 Your AI Research & Life Assistant             │
╰──────────────────────────────────────────────────────────────╯

Available Commands
──────────────────────────────────────────────────────────────
  chat               Chat with your AI assistant
  knowledge          Import PDF and text documents
  thread             Manage conversation threads
  credentials        Manage user and thread IDs
  start              Start the backend service
  stop               Stop the backend service
  shell              Launch the interactive shell
  version            Show application version
  help               Show command help

Thread Commands
──────────────────────────────────────────────────────────────
  thread list        List all conversation threads
  thread delete      Delete a conversation thread
  thread clear       Delete all conversation threads

Tips
──────────────────────────────────────────────────────────────
  • Type 'help' to see all available commands.
  • Type '<command> --help' for detailed help.
  • Type 'exit' to leave the interactive shell.

`)

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("scratchpad> ")

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if line == "exit" {
			break
		}

		args := strings.Fields(line)

		root.SetArgs(args)

		err := root.Execute()

		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
