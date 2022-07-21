package pkg

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var Root Dir

func Start() {
	Root.File.Path = "."
	Root.File.Name = "."

	ReadDir(&Root)

	dump(&Root, 0, "")
}

const space3 = "   "
const space4 = "    "

func dump(dir *Dir, i int, head string) {

	if i == 0 {
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
		if item.File.IsDir {
			lastLevel := ""
			if !isLast {
				lastLevel = "│" + space3
			} else {
				lastLevel = space4
			}
			dump(&item, i+1, head+lastLevel)
		}
	}
}

func ReadDir(dir *Dir) ByteSize {
	// 获取当前目录下所有文件 或 目录
	infos, err := ioutil.ReadDir(dir.File.Path)
	if err != nil {
		fmt.Println(err)
	}

	childLen := len(infos)

	if childLen == 0 {
		return 0
	}

	childs := make([]Dir, 0, childLen)
	// 遍历所有文件 或 目录
	for _, info := range infos {
		item := Dir{}
		item.File.Name = info.Name()
		item.File.IsDir = info.IsDir()
		item.File.Size = ByteSize(info.Size())
		// 如果是目录的话，记录路径，文件的话不需要记录
		if item.File.IsDir {
			item.File.Path = dir.File.Path + string(filepath.Separator) + info.Name()
		}
		childs = append(childs, item)
	}

	dir.Clilds = childs

	var dirSize ByteSize
	for index, item := range dir.Clilds {
		if item.File.IsDir {
			dir.Clilds[index].File.Size = ReadDir(&dir.Clilds[index])
		}
		dirSize = dirSize + item.File.Size
	}

	dir.File.Size = dirSize

	return dirSize
}
