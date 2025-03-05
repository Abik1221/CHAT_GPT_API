package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("missing api key")
	}

	ctx := context.Background()

	client := gpt3.NewClient(apiKey)
	response, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"this is the first thing you have to know about golang"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Choices[0].Text)
}
