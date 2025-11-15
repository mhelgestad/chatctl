package common

import (
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetOpenAIClient() *openai.Client {
	return openai.NewClient(os.Getenv("CHATCTL_OPEN_API_KEY"))
}

func GetSystemPrompt() string {
	defaultPrompt := "You are a helpful assistant. Please respond to this question with a brief but informative answer: "
	prompt, exists := os.LookupEnv("CHATCTL_OPENAI_SYSTEM_PROMPT")
	if !exists {
		return defaultPrompt
	}
	return prompt
}

func GetModel() string {
	defaultModel := "gpt-5-nano"
	model, exists := os.LookupEnv("CHATCTL_OPENAI_MODEL")
	if !exists {
		return defaultModel
	}
	return model
}
