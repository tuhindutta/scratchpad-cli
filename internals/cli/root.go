package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	apirequests "github.com/tuhindutta/scratchpad-cli/internals/apiRequests"
	cliCreds "github.com/tuhindutta/scratchpad-cli/internals/cli/creds"
)

var rootCmd = &cobra.Command{
	Use:   "scratchpad",
	Short: "Personal Assistant",
	Long:  `An AI personal assistant for day to day life help and research.`,
}

var credCmd = &cobra.Command{
	Use:     "credentials",
	Aliases: []string{"cred"},
	Short:   "Manage credentials",
}

func SetUserThreadIDsPortCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "set arg[1] arg[2]",
		Short: "Set user and thread IDs, port",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			if *app.ServerRunning != true {
				userId := args[0]
				threadId := args[1]
				cliCreds.SetUserThreadIDsPort(userId, threadId, app.CredentialPath)
				_, _, port := cliCreds.ReadUserThreadIDsPort(app.CredentialPath)
				app.UserId = userId
				app.ThreadId = threadId
				app.Port = port
			} else {
				fmt.Printf("Cannot change port while server is running.")
			}
			fmt.Println()
		},
	}

	return command
}

func SetUserThreadIDsCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "setUserThread arg[1] arg[2]",
		Short: "Set user and thread IDs",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			userId := args[0]
			threadId := args[1]
			cliCreds.SetUserThreadIDsPort(userId, threadId, app.CredentialPath)
			app.UserId = userId
			app.ThreadId = threadId

			fmt.Println()
		},
	}

	return command
}

func GetUserThreadIDsPortCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "get",
		Short: "Get user and thread IDs, port",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`
user_id:   %s
thread_id: %s
port:      %d

`, app.UserId, app.ThreadId, app.Port)
		},
	}

	return command
}

func StartCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "start",
		Short: "Initialize and start the service.",
		Run: func(cmd *cobra.Command, args []string) {
			srv := Server{Port: app.Port}
			fmt.Printf("Starting assistant server")
			srv.Start()
			val := true
			app.ServerRunning = &val

			fmt.Println()
		},
	}

	return command
}

func StopCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "stop",
		Short: "Stop the running service.",
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", app.Url, app.Port)
			apirequests.Shutdown(api_url)
			val := false
			app.ServerRunning = &val

			fmt.Println()
		},
	}

	return command
}

func IngestCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:     "knowledge",
		Aliases: []string{"ingest"},
		Short:   "Ingest external provided knowledge in .pdf and .txt formats.",
		Run: func(cmd *cobra.Command, args []string) {
			apiURL := fmt.Sprintf("%s:%d", app.Url, app.Port)

			out, errChan := apirequests.Ingest(apiURL, app.UserId, app.ThreadId)

			for out != nil || errChan != nil {
				select {
				case msg, ok := <-out:
					if !ok {
						out = nil
						continue
					}
					fmt.Print(msg)

				case err, ok := <-errChan:
					if !ok {
						errChan = nil
						continue
					}
					if err != nil {
						fmt.Println("Error:", err)
					}
				}
			}
			fmt.Println()
		},
	}

	return command
}

var threadCmd = &cobra.Command{
	Use:   "thread",
	Short: "Manage conversation threads",
}

func ListThreadsCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "list",
		Short: "List conversation threads.",
		Run: func(cmd *cobra.Command, args []string) {
			apiURL := fmt.Sprintf("%s:%d", app.Url, app.Port)

			out, errChan := apirequests.ListThreads(apiURL)

			for out != nil || errChan != nil {
				select {
				case msg, ok := <-out:
					if !ok {
						out = nil
						continue
					}
					fmt.Print(msg)

				case err, ok := <-errChan:
					if !ok {
						errChan = nil
						continue
					}
					if err != nil {
						fmt.Println("Error:", err)
					}
				}
			}
			fmt.Println()
		},
	}

	return command
}

func DeleteThreadCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "delete [arg1]",
		Short: "Delete conversation thread.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", app.Url, app.Port)
			apirequests.DeleteChatThread(api_url, args[0])

			fmt.Println()
		},
	}

	return command
}

func DeleteFullChatCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:   "clear",
		Short: "Delete all conversation threads.",
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", app.Url, app.Port)
			apirequests.DeleteFullChatHistory(api_url)

			fmt.Println()
		},
	}

	return command
}

func AssistantCmd(app *App) *cobra.Command {
	var command = &cobra.Command{
		Use:     "chat <userMessage>",
		Aliases: []string{"c"},
		Short:   "Assistant chat.",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			apiURL := fmt.Sprintf("%s:%d", app.Url, app.Port)
			prompt := strings.Join(args, " ")

			fmt.Println(prompt)

			out, errChan := apirequests.Assistant(
				apiURL,
				app.UserId,
				app.ThreadId,
				prompt,
			)

			for out != nil || errChan != nil {
				select {
				case msg, ok := <-out:
					if !ok {
						out = nil
						continue
					}
					fmt.Print(msg)

				case err, ok := <-errChan:
					if !ok {
						errChan = nil
						continue
					}
					if err != nil {
						fmt.Println("Error:", err)
					}
				}
			}

			fmt.Println()
		},
	}

	return command
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scratchpad v1.0.0")
		fmt.Println()
	},
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start interactive shell",
	RunE: func(cmd *cobra.Command, args []string) error {
		return StartShell(cmd.Root())
	},
}

type App struct {
	Url            string
	Port           int
	UserId         string
	ThreadId       string
	CredentialPath string
	ServerRunning  *bool
}

func (a *App) Execute() {

	threadCmd.AddCommand(ListThreadsCmd(a))
	threadCmd.AddCommand(DeleteThreadCmd(a))
	threadCmd.AddCommand(DeleteFullChatCmd(a))

	credCmd.AddCommand(SetUserThreadIDsPortCmd(a))
	credCmd.AddCommand(SetUserThreadIDsCmd(a))
	credCmd.AddCommand(GetUserThreadIDsPortCmd(a))

	rootCmd.AddCommand(credCmd)
	rootCmd.AddCommand(StartCmd(a))
	rootCmd.AddCommand(StopCmd(a))
	rootCmd.AddCommand(IngestCmd(a))
	rootCmd.AddCommand(threadCmd)
	rootCmd.AddCommand(AssistantCmd(a))
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(shellCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
