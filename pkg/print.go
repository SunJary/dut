package pkg

import (
	"fmt"

	"github.com/fatih/color"
)

const space3 = "   "
const space4 = "    "

// 目录名打印为蓝色文字
var dirP = color.BlueString

/*
	打印目录结构
	dir 目录
	level 要打印多少层
	n 当前是
*/
func Print(dir *Dir, config Config, n int, head string) {

	if n == 0 {
		fmt.Printf("%-8s %s\n", renderSize(dir.File.Size, config.SzieUnit), dir.File.Name)
	}

	len := len(dir.Clilds)
	for index, item := range dir.Clilds {
		isLast := false
		if index+1 == len {
			isLast = true
		}
		var currentLevel string
		if isLast {
			currentLevel = "└"
		} else {
			currentLevel = "├"
		}

		near := currentLevel + "──"

		// 如果是目录，打印为蓝色字
		var fileNameP string
		if item.File.IsDir {
			fileNameP = dirP(item.File.Name)
		} else {
			fileNameP = item.File.Name
		}

		fmt.Printf("%-8s %s %s\n", renderSize(item.File.Size, config.SzieUnit), head+near, fileNameP)
		if (config.PrintLevel < 0 || n < config.PrintLevel) && item.File.IsDir {
			lastLevel := ""
			if !isLast {
				lastLevel = "│" + space3
			} else {
				lastLevel = space4
			}

			Print(&item, config, n+1, head+lastLevel)
		}
	}
}
