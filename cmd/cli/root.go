package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	apirequests "github.com/tuhindutta/scratchpad-cli/internals/apiRequests"
)

var rootCmd = &cobra.Command{
	Use:   "scp",
	Short: "Personal Assistant",
	Long:  `An AI personal assistant for day to day life help and research.`,
}

func SetUserThreadIDsCmd(userId string, threadID string, credentialPath string) *cobra.Command {
	var command = &cobra.Command{
		Use:   "cred arg[1] arg[1]",
		Short: "Ser user and thread IDs",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			SetUserThreadIDs(args[0], args[1], credentialPath)
		},
	}

	return command
}

func InitCmd(port int) *cobra.Command {
	var command = &cobra.Command{
		Use:   "init",
		Short: "Initialize and start the service.",
		Run: func(cmd *cobra.Command, args []string) {
			srv := Server{Port: port}
			fmt.Printf("Starting assistant server")
			srv.Start()
		},
	}

	return command
}

func StopCmd(url string, port int) *cobra.Command {
	var command = &cobra.Command{
		Use:   "stop",
		Short: "Stop the running service.",
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", url, port)
			apirequests.Shutdown(api_url)
		},
	}

	return command
}

func IngestCmd(url string, port int, userId string, threadId string) *cobra.Command {
	var command = &cobra.Command{
		Use:   "ingest",
		Short: "Ingest external provided knowledge in .pdf and .txt formats.",
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", url, port)
			apirequests.Ingest(api_url, userId, threadId)
		},
	}

	return command
}

func ListThreadsCmd(url string, port int) *cobra.Command {
	var command = &cobra.Command{
		Use:   "lt",
		Short: "List conversation threads.",
		// Run: func(cmd *cobra.Command, args []string) {

		// 	api_url := fmt.Sprintf("%s:%d", url, port)
		// 	apirequests.ListThreads(api_url)
		// },
		Run: func(cmd *cobra.Command, args []string) {
			apiURL := fmt.Sprintf("%s:%d", url, port)

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
		},
	}

	return command
}

func DeleteThreadCmd(url string, port int, threadId string) *cobra.Command {
	var command = &cobra.Command{
		Use:   "dt",
		Short: "Delete conversation thread.",
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", url, port)
			apirequests.DeleteChatThread(api_url, threadId)
		},
	}

	return command
}

func DeleteFullChatCmd(url string, port int) *cobra.Command {
	var command = &cobra.Command{
		Use:   "dc",
		Short: "Delete all conversation threads.",
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", url, port)
			apirequests.DeleteFullChatHistory(api_url)
		},
	}

	return command
}

func AssistantCmd(url string, port int, userId string, threadId string) *cobra.Command {
	var command = &cobra.Command{
		Use:   "a <userMessage>",
		Short: "Assistant chat.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			apiURL := fmt.Sprintf("%s:%d", url, port)
			prompt := strings.Join(args, " ")

			fmt.Println(prompt)

			out, errChan := apirequests.Assistant(
				apiURL,
				userId,
				threadId,
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
		},
	}

	return command
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scratchpad v1.0.0")
	},
}

type App struct {
	Url            string
	Port           int
	UserId         string
	ThreadId       string
	CredentialPath string
}

func (a App) Execute() {

	rootCmd.AddCommand(SetUserThreadIDsCmd(a.UserId, a.ThreadId, a.CredentialPath))
	rootCmd.AddCommand(InitCmd(a.Port))
	rootCmd.AddCommand(StopCmd(a.Url, a.Port))
	rootCmd.AddCommand(IngestCmd(a.Url, a.Port, a.UserId, a.ThreadId))
	rootCmd.AddCommand(ListThreadsCmd(a.Url, a.Port))
	rootCmd.AddCommand(DeleteThreadCmd(a.Url, a.Port, a.ThreadId))
	rootCmd.AddCommand(DeleteFullChatCmd(a.Url, a.Port))
	rootCmd.AddCommand(AssistantCmd(a.Url, a.Port, a.UserId, a.ThreadId))
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
