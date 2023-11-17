package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() (bool, error) {
	var desktopConfig bool

	var rootCmd = &cobra.Command{
		Use:   "averveil",
		Short: "Averveil is OpenSea for Data",
		Run: func(cmd *cobra.Command, args []string) {
			if desktopConfig {
				fmt.Println("desktopConfig is set to true")
				// Additional logic when desktopConfig is true
			} else {
				fmt.Println("desktopConfig is set to false")
				// Additional logic when desktopConfig is false
			}
		},
	}

	// Define the desktopConfig flag
	rootCmd.Flags().BoolVarP(&desktopConfig, "desktopConfig", "d", false, "Set config for desktop")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		return false, err
	}

	return desktopConfig, nil
}
