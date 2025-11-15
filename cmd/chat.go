/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mhelgestad/chatctl/common"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Launch a chat session",
	Run: func(cmd *cobra.Command, args []string) {
		client := common.GetOpenAIClient()
		chatMessages := initializeChatHistory()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter your messages below. Type 'exit' to quit.")
		for {
			fmt.Print("You >>> ")
			prompt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			prompt = strings.TrimSpace(prompt)

			if prompt == "exit" {
				break
			}

			if prompt == "save" {
				data, _ := json.Marshal(chatMessages)
				os.WriteFile("chat_history.json", data, 0644)
				fmt.Println("Chat history saved")
				continue
			}

			chatMessages = append(chatMessages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			})

			resp, err := client.CreateChatCompletion(
				cmd.Context(),
				openai.ChatCompletionRequest{
					Model:    common.GetModel(),
					Messages: chatMessages,
				},
			)

			if err != nil {
				fmt.Printf("\nBot >>> %v\n", err)
				continue
			}
			chatMessages = append(chatMessages, resp.Choices[0].Message)
			fmt.Printf("\nBot >>> %s\n\n", resp.Choices[0].Message.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initializeChatHistory() []openai.ChatCompletionMessage {
	var chatMessages []openai.ChatCompletionMessage
	data, err := os.ReadFile("chat_history.json")
	if err != nil {
		// If file doesn't exist, start with system prompt
		chatMessages = []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: common.GetSystemPrompt(),
			},
		}
		return chatMessages
	}
	err = json.Unmarshal(data, &chatMessages)
	if err != nil {
		// If unmarshal fails, start fresh
		chatMessages = []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: common.GetSystemPrompt(),
			},
		}
	}
	return chatMessages
}
