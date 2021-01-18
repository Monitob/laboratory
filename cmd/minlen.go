package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.nereides.weborama.com/go/playground/Test/laboratory/pkg/polymere"
)

func init() {
	RootCmd.AddCommand(minLenCmd)
}

var minLenCmd = &cobra.Command{
	Use:   "minlen abBC",
	Short: "minimal length reaction",
	Long:  `Print the lenght of the shortest reaction removing one element`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range args {
			l, err := polymere.LenMinimalReaction(p)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
				os.Exit(1)
			}
			fmt.Printf("polymere: %s, minimal length: %d\n", p, l)
		}
	},
}
