package pkg

import "fmt"

const space3 = "   "
const space4 = "    "

/*
	打印目录结构
	dir 目录
	level 要打印多少层
	n 当前是
*/
func Print(dir *Dir, level, n int, head string) {

	if n == 0 {
		fmt.Printf("%-8s %s\n", renderSize(dir.File.Size, 0), dir.File.Name)
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
		fmt.Printf("%-8s %s %s\n", renderSize(item.File.Size, 0), head+near, item.File.Name)
		if (level < 0 || n < level) && item.File.IsDir {
			lastLevel := ""
			if !isLast {
				lastLevel = "│" + space3
			} else {
				lastLevel = space4
			}

			Print(&item, level, n+1, head+lastLevel)
		}
	}
}
