package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.nereides.weborama.com/go/playground/Test/laboratory/pkg/polymere"
)

var reactCmd = &cobra.Command{
	Use:   "react abBC",
	Short: "make polymere reaction",
	Long:  `Print the reaction polymere`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range args {
			r, err := polymere.React(p)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
				os.Exit(1)
			}
			fmt.Printf("polymere: %s, reaction: %s\n", p, r)
		}
	},
}
