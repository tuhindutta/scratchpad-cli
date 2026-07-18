package backend

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/tuhindutta/scratchpad-cli/internals/envSetup"
)

func Setup() {

	// Verify Python installed version
	valid := envSetup.VerifyPythonVersion()
	if !valid {
		log.Fatalf("Python version required above 3.12.12")
	} else {
		slog.Info("Valid Python installation found.")
	}

	// Downloading backend github repo
	owner := "tuhindutta"
	repo := "scratchpad-backend"
	branch := "main"
	zipURL := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/%s.zip", owner, repo, branch)

	backendDir := "./backend"

	fmt.Printf("Downloading %s...\n", zipURL)
	if err := envSetup.DownloadBackend(zipURL, backendDir); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	slog.Info("Success! Repository extracted.")

	// Create Virtual Env
	venvDir := "./backend/projectenv"

	envSetup.CreateVenv(venvDir)

	// Installing backend packages
	venv := "./backend/projectenv"
	req := "./backend/scratchpad-backend-main/requirements.txt"
	err := envSetup.InstallBackendDependencies(venv, req)
	if err == nil {
		fmt.Printf("Error: %v\n", err)
	}
}
