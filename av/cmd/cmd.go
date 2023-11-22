package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute() error {
	var envConfig bool

	var rootCmd = &cobra.Command{
		Use:   "averveil",
		Short: "Averveil is OpenSea for Data",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(envConfig)
		},
	}

	rootCmd.Flags().BoolVarP(&envConfig, "envConfig", "e", false, "Set configs from desktop")
	rootCmd.Flags().BoolVarP(&envConfig, "swarmEnabled", "s", false, "Enable Swarm")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
