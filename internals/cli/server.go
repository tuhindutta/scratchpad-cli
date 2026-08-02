package cli

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

	be := backend.Backend{}

	s.Backend = &be

	s.Backend.Start(s.Port)
	time.Sleep(15 * time.Second)
}
