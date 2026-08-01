package cmd

import (
	"os"
	"path/filepath"
	"time"

	"github.com/tuhindutta/scratchpad-cli/internals/backend"
)

type Server struct {
	Port    int
	Backend *backend.Backend
}

func (s *Server) Start() {

	workDir := filepath.Join("./../../", os.Getenv("WORK_DIR"))
	os.Setenv("WORK_DIR", workDir)

	// be := backend.Backend{
	// 	AgentModel:      os.Getenv("AGENT_MODEL"),
	// 	RagModel:        os.Getenv("RAG_MODEL"),
	// 	OpenaiApiKey:    os.Getenv("OPENAI_API_KEY"),
	// 	GoogleApiKey:    os.Getenv("GOOGLE_API_KEY"),
	// 	FirecrawlApiKey: os.Getenv("FIRECRAWL_API_KEY"),
	// 	WorkDir:         filepath.Join("./../../", os.Getenv("WORK_DIR")),
	// }

	be := backend.Backend{}

	s.Backend = &be

	s.Backend.Start(s.Port)
	time.Sleep(15 * time.Second)
}

// func (s *Server) Stop() {
// 	// s.Backend.Stop()
// 	url := fmt.Sprintf("%s:%d", s.Url, s.Port)
// 	apirequests.Shutdown(url)
// }
