package backend

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func StartServer(port int, venv_path string, backend_dir string) (*exec.Cmd, error) {

	python := filepath.Join(venv_path, "bin", "python")

	cmd := exec.Command(python, "main.py", "--port", strconv.Itoa(port))

	cmd.Dir = backend_dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	log.Println("Starting server...")
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
		return nil, err
	}

	return cmd, nil
}
