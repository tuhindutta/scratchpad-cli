package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/tuhindutta/scratchpad-cli/internals/backend"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	be := backend.Backend{
		AgentModel:      os.Getenv("AGENT_MODEL"),
		RagModel:        os.Getenv("RAG_MODEL"),
		OpenaiApiKey:    os.Getenv("OPENAI_API_KEY"),
		GoogleApiKey:    os.Getenv("GOOGLE_API_KEY"),
		FirecrawlApiKey: os.Getenv("FIRECRAWL_API_KEY"),
		WorkDir:         os.Getenv("WORK_DIR"),
	}

	be.Start(8081)

	time.Sleep(30 * time.Second)

	be.Stop()
}
