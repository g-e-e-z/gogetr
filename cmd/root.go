package cmd

import (
	"fmt"
	"os"

	"github.com/g-e-e-z/gogetr/gui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "gogetr",
    Short: "A simple HTTP client with TUI",
    Run: func(cmd *cobra.Command, args []string) {
        // Initialize and run the TUI
        gui.RunTUI()
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

