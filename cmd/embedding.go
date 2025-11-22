/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mhelgestad/chatctl/common"
	"github.com/spf13/cobra"
)

// embeddingCmd represents the embedding command
var embeddingCmd = &cobra.Command{
	Use:   "embedding",
	Short: "generate an embedding for some text",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var text string
		if len(args) == 0 {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			text = strings.TrimSpace(string(data))
		} else {
			text = args[0]
		}
		model, err := cmd.Flags().GetString("model")
		if model != "all-MiniLM-L6-v2" && model != "all-MiniLM-L12-v2" && model != "all-mpnet-base-v2" {
			fmt.Println("Model must be one of ['all-MiniLM-L6-v2', 'all-MiniLM-L12-v2', 'all-mpnet-base-v2']")
			return nil
		}
		if err != nil {
			panic(err)
		}
		embeddingResponse, err := common.GenerateEmbedding(model, text)
		if err != nil {
			return fmt.Errorf("embedding error: %v", err)
		}
		fmt.Println(embeddingResponse.Embedding)
		return nil
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// embeddingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	embeddingCmd.Flags().StringP("model", "m", "all-MiniLM-L6-v2", "sentence transformer model")
	rootCmd.AddCommand(embeddingCmd)
}
