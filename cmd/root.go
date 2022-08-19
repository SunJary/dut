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
	Run:   run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&config.PrintLevel, "level", "l", 1, "tree level")

	rootCmd.Flags().BoolVarP(&config.Interact, "interact", "i", false, "interact")

	rootCmd.Flags().BoolVarP(&config.Usage, "usage", "u", false, "usage")

	rootCmd.Flags().StringVarP(&config.ByteSzieUnit, "byte", "b", "", "Byte size unit,allow [K|M|G|T|P|E|Z|Y] , default auto")
}

func run(cmd *cobra.Command, args []string) {
	err := config.Init()
	if err != nil {
		cobra.CheckErr(err)
		return
	}

	pkg.RunDut(args, config)
}
