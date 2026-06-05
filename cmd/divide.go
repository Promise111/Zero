package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var divideCmd = &cobra.Command{
	Use:     "divide",
	Aliases: []string{"part", "split"},
	Short:   "Divide 2 numbers",
	Long:    "Divide second args from first args",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("The quotient of %s and %s = %s", args[0], args[1], Divide(args[0], args[1], shouldRoundUp))
	},
}

func init() {
	divideCmd.Flags().BoolVarP(&shouldRoundUp, "divide", "d", false, "Divides two number")
	rootCmd.AddCommand(divideCmd)
}
