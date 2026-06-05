package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"subtraction", "sub"},
	Short:   "subtract 2 numbers",
	Long:    "Handles subtract operator on two numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Difference of %s - %s = %s", args[0], args[1], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
