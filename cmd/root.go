/*
Copyright © 2022 Jary <jarysun@outlook.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/SunJary/dut/pkg"
	"github.com/spf13/cobra"
)

var (
	level int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dut",
	Short: "du like command",
	Long:  `du like command in tree view`,
	Run: func(cmd *cobra.Command, args []string) {
		len := len(args)
		var dirList = make([]pkg.Dir, 0, len)
		if len == 0 {
			// 没有参数，默认当前目录
			dir := pkg.Dir{}
			dir.File.Name = "."
			dir.File.Path = "."
			dir.File.IsDir = true

			dirList = append(dirList, dir)
		} else {
			// 有参数，遍历，加入list
			for _, dirName := range args {
				// 判断文件是否存在
				info, err := os.Stat(dirName)
				if err != nil {
					if os.IsPermission(err) {
						fmt.Printf("read %s permission is denied\n", dirName)
					}

					fmt.Printf("%s not exist\n", dirName)
					continue
				}

				dir := pkg.Dir{}
				dir.File.Name = dirName
				dir.File.IsDir = info.IsDir()
				if dir.File.IsDir {
					dir.File.Path = dirName
				} else {
					dir.File.Size = pkg.ByteSize(info.Size())
				}
				dirList = append(dirList, dir)
			}

		}

		for index := range dirList {

			pkg.ReadDir(&dirList[index])

			pkg.Print(&dirList[index], level-1, 0, "")
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&level, "level", "l", 1, "tree level")
}
