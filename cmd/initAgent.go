/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mhelgestad/chatctl/common"
	"github.com/spf13/cobra"
)

// initAgentCmd represents the initAgent command
var initAgentCmd = &cobra.Command{
	Use:   "initAgent",
	Short: "Initialize the error explain agent",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		model := common.GetModel()
		r, err := common.InitAgent(model)
		if err != nil {
			return fmt.Errorf("error initializing agent: %v", err)
		}
		fmt.Println(r.Message)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initAgentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initAgentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initAgentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
