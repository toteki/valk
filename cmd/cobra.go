package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "valk",
	Short: "Short Desc",
	Long:  `Long Desc`,
	Run:   BaseCmd,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func BaseCmd(cmd *cobra.Command, args []string) {
	// This runs when no subcommand is given
	fmt.Println("Valk Base Cmd")
}
