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
		fmt.Printf("%s\n", dir.File.Name)
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
		fmt.Printf("%s %s\n", head+near, item.File.Name)
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

func ReadDir(dir *Dir) *Dir {
	infos, err := ioutil.ReadDir(dir.File.Path)
	if err != nil {
		fmt.Println(err)
	}

	childLen := len(infos)

	childs := make([]Dir, 0, childLen)
	for _, info := range infos {
		item := Dir{}
		item.File.Name = info.Name()
		item.File.Path = dir.File.Path + string(filepath.Separator) + info.Name()
		item.File.IsDir = info.IsDir()
		item.File.Size = ByteSize(info.Size())
		childs = append(childs, item)
	}

	dir.Clilds = childs

	for index, item := range dir.Clilds {
		if item.File.IsDir {
			ReadDir(&dir.Clilds[index])
		}
	}

	return dir
}
