package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tuhindutta/scratchpad-cli/cmd/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute(os.Getenv("URL"), 8081)
}

// prompts, errs := apirequests.Assistant("http://localhost:8081", "u1", "t1", "Hello!")

// for prompt := range prompts {
// 	fmt.Print(prompt)
// }

// if err := <-errs; err != nil {
// 	log.Fatalf("Assistant stream failed: %v", err)
// }
