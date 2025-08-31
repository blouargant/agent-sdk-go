package oaic

import (
	"github.com/Ingenimax/agent-sdk-go/pkg/interfaces"
	"github.com/Ingenimax/agent-sdk-go/pkg/llm/openai_base"
	"github.com/Ingenimax/agent-sdk-go/pkg/logging"
	"github.com/Ingenimax/agent-sdk-go/pkg/retry"
)

type OpenAICompatibleClient struct {
	openai_base.OpenAIBaseClient
}

// WithModel sets the model for the OpenAI client
func WithModel(model string) openai_base.Option {
	return openai_base.WithModel(model)
}

// WithLogger sets the logger for the OpenAI client
func WithLogger(logger logging.Logger) openai_base.Option {
	return openai_base.WithLogger(logger)
}

// NewClient creates a new OpenAI client
func NewClient(apiKey string, options ...openai_base.Option) *OpenAICompatibleClient {
	return &OpenAICompatibleClient{
		OpenAIBaseClient: *openai_base.NewClient(apiKey, options...),
	}
}

func WithRetry(opts ...retry.Option) openai_base.Option {
	return openai_base.WithRetry(opts...)
}

func WithTemperature(temperature float64) interfaces.GenerateOption {
	return openai_base.WithTemperature(temperature)
}

func WithTopP(topP float64) interfaces.GenerateOption {
	return openai_base.WithTopP(topP)
}

func WithPresencePenalty(presencePenalty float64) interfaces.GenerateOption {
	return openai_base.WithPresencePenalty(presencePenalty)
}

func WithStopSequences(stopSequences []string) interfaces.GenerateOption {
	return openai_base.WithStopSequences(stopSequences)
}

func WithSystemMessage(systemMessage string) interfaces.GenerateOption {
	return openai_base.WithSystemMessage(systemMessage)
}

func WithResponseFormat(format interfaces.ResponseFormat) interfaces.GenerateOption {
	return openai_base.WithResponseFormat(format)
}

func WithReasoning(reasoning string) interfaces.GenerateOption {
	return openai_base.WithReasoning(reasoning)
}
