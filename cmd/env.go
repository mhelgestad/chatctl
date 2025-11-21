/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "See environment information",
	Run: func(cmd *cobra.Command, args []string) {
		model, exists := os.LookupEnv("CHATCTL_OPENAI_MODEL")
		if !exists {
			model = "gpt-5-nano"
		}

		prompt, exists := os.LookupEnv("CHATCTL_OPENAI_SYSTEM_PROMPT")
		if !exists {
			prompt = "You are a helpful assistant. Please respond to this question with a brief but informative answer: "
		}

		agentUrl, exists := os.LookupEnv("CHATCTL_AGENT_BASE_URL")
		if !exists {
			agentUrl = "http://localhost:8000/"
		}

		enhancePrompt, exists := os.LookupEnv("CHATCTL_ENHANCE_PROMPT")
		if !exists {
			enhancePrompt = "I'm an expert prompt engineer. Improve the following prompt to be more effective and precise, and only provide the prompt in the output. Limit the enhanced prompt to three sentences:"
		}

		openaiKey, exists := os.LookupEnv("CHATCTL_OPEN_API_KEY")
		if !exists {
			openaiKey = "NOT_SET"
		} else {
			openaiKey = "******"
		}

		fmt.Printf("CHATCTL_OPENAI_MODEL=\"%s\"\n", model)
		fmt.Printf("CHATCTL_OPENAI_SYSTEM_PROMPT=\"%s\"\n", prompt)
		fmt.Printf("CHATCTL_AGENT_BASE_URL=\"%s\"\n", agentUrl)
		fmt.Printf("CHATCTL_ENHANCE_PROMPT=\"%s\"\n", enhancePrompt)
		fmt.Printf("CHATCTL_OPEN_API_KEY=\"%s\"\n", openaiKey)
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
