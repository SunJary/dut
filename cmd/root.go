/*
Copyright © 2022 Jary <jarysun@outlook.com>

*/
package cmd

import (
	"os"

	"github.com/SunJary/dut/pkg"
	"github.com/spf13/cobra"
)

var config = pkg.Config{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dut",
	Short: "du like command",
	Long:  `du like command in tree view`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.RunDut(args, config)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var level int
	rootCmd.Flags().IntVarP(&level, "level", "l", 1, "tree level")
	config.PrintLevel = level - 1

	rootCmd.Flags().BoolVarP(&config.Interact, "it", "i", false, "interact")

	rootCmd.Flags().StringVarP(&config.SzieUnit, "unit", "u", "", "sizeUnit K/M/G/T/P/E/Z/Y")

}
