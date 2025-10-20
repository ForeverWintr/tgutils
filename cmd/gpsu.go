/*
Copyright © 2025 foreverwintr@gmail.com
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// gpsuCmd represents the gpsu command
var gpsuCmd = &cobra.Command{
	Use:   "gpsu",
	Short: "git push --set-upstream",
	Long:  `Push the current branch to the remote repository and set the upstream tracking reference.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		command := exec.Command("git", "push", "--set-upstream", "origin", string(result))
		fmt.Println("Executing command:", command.String())
		result, err = command.Output()
		if err != nil {
			fmt.Println(string(result))
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Current branch:", string(result))
	},
}

func init() {
	rootCmd.AddCommand(gpsuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gpsuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gpsuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
