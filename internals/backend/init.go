package backend

import (
	"log"
	"log/slog"
	"os"
	"os/exec"
)

type Backend struct {
	AgentModel      string
	RagModel        string
	OpenaiApiKey    string
	GoogleApiKey    string
	FirecrawlApiKey string
	WorkDir         string

	port *int
	cmd  *exec.Cmd
}

func (b Backend) injectEnvVar() {
	os.Setenv("AGENT_MODEL", b.AgentModel)
	os.Setenv("RAG_MODEL", b.RagModel)
	os.Setenv("OPENAI_API_KEY", b.OpenaiApiKey)
	os.Setenv("GOOGLE_API_KEY", b.GoogleApiKey)
	os.Setenv("FIRECRAWL_API_KEY", b.FirecrawlApiKey)
	os.Setenv("WORK_DIR", b.WorkDir)
}

func (b *Backend) init(port int) {
	_, err := os.Stat("backend")
	if os.IsNotExist(err) {
		Setup()
	}

	venv_path := "./../projectenv"
	backend_dir := "./backend/scratchpad-backend-main"

	b.port = &port
	// StartServer(port, venv_path, backend_dir)
	cmd, err := StartServer(port, venv_path, backend_dir)
	if err != nil {
		log.Fatalf("Failed to start backend: %v", err)
	}

	b.cmd = cmd

	go func() {
		if err := cmd.Wait(); err != nil {
			log.Fatalf("Backend exited: %v\n", err)
		}
	}()
}

func (b *Backend) Start(port int) {
	b.injectEnvVar()
	b.init(port)
}

func (b *Backend) Stop() {

	if b.cmd == nil || b.cmd.Process == nil {
		log.Println("Backend not running.")
		return
	} else {

		slog.Info("Stopping backend...")

		if err := b.cmd.Process.Kill(); err != nil {
			log.Printf("Failed to stop backend: %s", err)
			return
		}

		slog.Info("Backend stopped.")

	}

}
