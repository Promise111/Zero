package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var multipyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"mult", "times"},
	Short:   "Multiply 2 numbers",
	Long:    "Perform multiplication on two numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Product of %s and %s = %s.\n\n", args[0], args[1], Multiply(args[0], args[1], shouldRoundUp))
	},
}

func init() {
	multipyCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
	rootCmd.AddCommand(multipyCmd)
}
