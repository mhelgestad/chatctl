/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mhelgestad/chatctl/common"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

// improvePromptCmd represents the improvePrompt command
var improvePromptCmd = &cobra.Command{
	Use:   "improvePrompt",
	Short: "helps improve a prompt",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		p := args[0]

		fmt.Println("Original Prompt\n----------------")
		fmt.Println(p)

		client := common.GetOpenAIClient()
		chatMessages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Im an expert prompt engineer. Improve the following prompt to be more effective and precise, and only provide the prompt in the output:",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: p,
			},
		}

		resp, err := client.CreateChatCompletion(
			cmd.Context(),
			openai.ChatCompletionRequest{
				Model:    common.GetModel(),
				Messages: chatMessages,
			},
		)

		if err != nil {
			return fmt.Errorf("completion error: %v", err)
		}
		fmt.Println("\nImproved Prompt\n----------------")
		fmt.Println(resp.Choices[0].Message.Content)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(improvePromptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// improvePromptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// improvePromptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
