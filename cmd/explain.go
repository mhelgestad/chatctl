/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mhelgestad/chatctl/common"
	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain code or command errors",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var q string
		if len(args) == 0 {
			var sb strings.Builder
			reader := bufio.NewReader(os.Stdin)
			for {
				line, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				if err != nil {
					return fmt.Errorf("failed to read stdin: %w", err)
				}
				sb.WriteString(line)
			}
			q = strings.TrimSpace(sb.String())
		} else {
			q = args[0]
		}
		fmt.Println(q + "\n")
		agentResponse, err := common.CallAgent(q)
		if err != nil {
			return fmt.Errorf("explain error: %s", err)
		}
		if agentResponse.Error != nil {
			return fmt.Errorf("explain error: %s", *agentResponse.Error)
		}
		fmt.Println("Topic\n-------")
		fmt.Println(agentResponse.Topic)
		fmt.Println("\nSummary\n-------")
		fmt.Println(agentResponse.Summary)
		fmt.Println("\nSources\n-------")
		for _, source := range agentResponse.Sources {
			fmt.Println("-", source)
		}
		fmt.Println("\nTools Used\n-------")
		for _, tool := range agentResponse.ToolsUsed {
			fmt.Println("-", tool)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// explainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// explainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
