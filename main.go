package main

import (
	"fmt"
	"os"

	"gitlab.nereides.weborama.com/go/playground/Test/laboratory/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}
