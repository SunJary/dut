package pkg

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	Root         Dir
	goos         = runtime.GOOS
	dirSp        = string(filepath.Separator)
	linuxSkipDir = map[string]struct{}{"/proc": {}, "/proc/": {}}
)

func ReadDir(dir *Dir) {
	if !dir.File.IsDir {
		return
	}
	// 获取当前目录下所有文件 或 目录
	infos, err := ioutil.ReadDir(dir.File.Path)
	if err != nil {
		fmt.Println(err)
	}

	childLen := len(infos)

	if childLen == 0 {
		return
	}

	var wg sync.WaitGroup

	childs := make([]Dir, 0, childLen)
	// 遍历所有文件 或 目录
	for _, info := range infos {
		item := Dir{}
		item.File.Name = info.Name()
		item.File.IsDir = info.IsDir()
		item.File.Size = ByteSize(info.Size())
		// 如果是目录的话，记录路径，文件的话不需要记录
		if item.File.IsDir {
			item.File.Path = filepath.Join(dir.File.Path, item.File.Name)
		}
		childs = append(childs, item)
	}

	dir.Clilds = childs

	for index, item := range dir.Clilds {
		if item.File.IsDir && !isSkipDir(item) {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				ReadDir(&dir.Clilds[i])
			}(index)
		}
	}

	// 等待子目录统计结束后
	wg.Wait()

	// 计算当前目录的大小
	var dirSize ByteSize
	for index := range dir.Clilds {
		dirSize = dirSize + dir.Clilds[index].File.Size
	}

	dir.File.Size = dirSize

	return
}

func isSkipDir(dir Dir) bool {
	if goos == "linux" {
		if _, ok := linuxSkipDir[dir.File.Path]; ok {
			return true
		}
	}
	return false
}
