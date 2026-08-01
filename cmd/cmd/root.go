package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	apirequests "github.com/tuhindutta/scratchpad-cli/internals/apiRequests"
)

// const pidFile = "/tmp/scratchpad-cli.pid"
// const defaultPort = 8081

var rootCmd = &cobra.Command{
	Use:   "scp",
	Short: "Personal Assistant",
	Long:  `An AI personal assistant for day to day life help and research.`,
}

func InitCmd(port int) *cobra.Command {
	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize and start the service.",
		Run: func(cmd *cobra.Command, args []string) {
			// 1. Save the current process ID so 'stop' can find it later
			// pid := os.Getpid()
			// _ = os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
			srv := Server{Port: port}
			fmt.Printf("Starting assistant server")
			srv.Start()
		},
	}

	return initCmd
}

// func StopServer() {
// 	// s.Backend.Stop()
// 	url := fmt.Sprintf("%s:%d", s.Url, s.Port)
// 	apirequests.Shutdown(url)
// }

func StopCmd(url string, port int) *cobra.Command {
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop the running service.",
		Run: func(cmd *cobra.Command, args []string) {

			api_url := fmt.Sprintf("%s:%d", url, port)
			apirequests.Shutdown(api_url)
			// 1. Read the saved process ID from disk
			// data, err := os.ReadFile(pidFile)
			// if err != nil {
			// 	fmt.Println("Error: Service does not appear to be running.")
			// 	return
			// }

			// pid, _ := strconv.Atoi(string(data))

			// // 2. Find the operating system process and kill it
			// process, err := os.FindProcess(pid)
			// if err != nil {
			// 	fmt.Printf("Failed to find process with PID %d\n", pid)
			// 	return
			// }

			// fmt.Printf("Stopping assistant server (PID: %d)...\n", pid)
			// err = process.Signal(syscall.SIGTERM) // Clean shutdown signal
			// if err != nil {
			// 	fmt.Println("Failed to kill process:", err)
			// 	return
			// }

			// // 3. Clean up the PID file
			// _ = os.Remove(pidFile)
			// fmt.Println("Server stopped successfully.")
		},
	}

	return stopCmd
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mycliapp v1.0.0")
	},
}

func Execute(url string, port int) {

	rootCmd.AddCommand(InitCmd(port))
	rootCmd.AddCommand(StopCmd(url, port))
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
