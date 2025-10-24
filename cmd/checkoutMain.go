/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// checkoutMainCmd represents the checkoutMain command
var checkoutMainCmd = &cobra.Command{
	Use:   "checkoutMain",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get name of remote default branch
		result, err := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD").CombinedOutput()
		if err != nil {
			fmt.Println(string(result))
			fmt.Println("Error:", err)
			return
		}

		branch_extract := regexp.MustCompile(`refs/remotes/origin/(.+)`)
		branch_name := branch_extract.FindStringSubmatch(string(result))[1]

		command := exec.Command("git", "checkout", branch_name)

		fmt.Println(command.String())
		result, err = command.CombinedOutput()
		if err != nil {
			fmt.Println(string(result))
			fmt.Println("Error:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(checkoutMainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutMainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutMainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
