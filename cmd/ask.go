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

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask [question]",
	Short: "Ask your question",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		q := args[0]

		client := common.GetOpenAIClient()
		chatMessages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: common.GetSystemPrompt(),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: q,
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
		fmt.Println(resp.Choices[0].Message.Content)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(askCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// add flag for agent url
}
