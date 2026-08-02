package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func StartShell(root *cobra.Command) error {

	fmt.Println("Scratchpad CLI")
	fmt.Println("Type 'exit' to quit.")

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
