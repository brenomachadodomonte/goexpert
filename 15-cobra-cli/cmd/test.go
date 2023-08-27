/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		command, _ := cmd.Flags().GetString("command")
		fmt.Println(command)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PRE RUN")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("POST RUN")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringP("command", "c", "", "Choose Ping or Pong")
	testCmd.MarkFlagRequired("command")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
