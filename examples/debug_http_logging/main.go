package main

import (
	"context"
	"log"
	"os"

	"github.com/Ingenimax/agent-sdk-go/pkg/interfaces"
	"github.com/Ingenimax/agent-sdk-go/pkg/llm/openai"
	"github.com/Ingenimax/agent-sdk-go/pkg/logging"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	// Create a logger that will show debug messages
	logger := logging.New()

	// Create OpenAI client with debug logging enabled
	client := openai.NewClient(apiKey,
		openai.WithModel("gpt-4o-mini"),
		openai.WithLogger(logger),
	)

	ctx := context.Background()

	// Example 1: Simple generate request with debug logging
	log.Println("=== Example 1: Simple Generate Request ===")
	response, err := client.Generate(ctx, "What is the capital of France?",
		openai.WithTemperature(0.7),
		openai.WithSystemMessage("You are a helpful geography assistant."),
	)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Response: %s", response)
	}

	// Example 2: Request with response format (JSON schema)
	log.Println("\n=== Example 2: Request with JSON Response Format ===")
	responseFormat := interfaces.ResponseFormat{
		Name: "country_info",
		Schema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"country": map[string]interface{}{
					"type":        "string",
					"description": "The country name",
				},
				"capital": map[string]interface{}{
					"type":        "string",
					"description": "The capital city",
				},
				"population": map[string]interface{}{
					"type":        "integer",
					"description": "The approximate population",
				},
			},
			"required": []string{"country", "capital"},
		},
	}

	response2, err := client.Generate(ctx, "Tell me about Japan",
		openai.WithResponseFormat(responseFormat),
		openai.WithSystemMessage("Provide information about the requested country in the specified JSON format."),
	)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("JSON Response: %s", response2)
	}

	// Example 3: Request with reasoning mode
	log.Println("\n=== Example 3: Request with Comprehensive Reasoning ===")
	response3, err := client.Generate(ctx, "Solve this math problem: If a train travels 120 km in 2 hours, what is its average speed?",
		openai.WithReasoning("comprehensive"),
		openai.WithSystemMessage("You are a math tutor."),
	)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Reasoning Response: %s", response3)
	}
}
