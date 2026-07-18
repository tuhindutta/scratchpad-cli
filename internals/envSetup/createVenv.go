package envSetup

import (
	"log"
	"log/slog"
	"os/exec"
)

func CreateVenv(venv_name string) {

	// env_dir := filepath.Join(directory, "projectenv")
	cmd := exec.Command("python3", "-m", "venv", venv_name)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("python3 command not found: %v", err)
	}

	slog.Info("Virtual environment created.")
}
