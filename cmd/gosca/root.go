package gosca

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gosca",
	Short: "cli for gosca any project scaffolding by a config file",
}

func Execute() {
	if cerr := rootCmd.Execute(); cerr != nil {
		fmt.Println(cerr)
		os.Exit(1)
	}
}
