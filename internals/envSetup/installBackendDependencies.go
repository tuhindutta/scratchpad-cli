package envSetup

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func InstallBackendDependencies(venv_path string, requirements_path string) error {

	pip := filepath.Join(venv_path, "bin", "pip3")

	// fmt.Println(pip)

	cmd := exec.Command(pip, "install", "-r", requirements_path)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	log.Println("Starting pip installation directly inside venv...")
	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return errors.New("pip exited with a non-zero status code")
		}
		return err
	}

	return nil
}
