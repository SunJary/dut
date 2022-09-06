package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	Root         Dir
	goos         = runtime.GOOS
	linuxSkipDir = map[string]struct{}{"/proc": {}, "/proc/": {}}

	// 将Goroutine数量大致限制在此值之下
	limitOfGoroutine = 20 * runtime.NumCPU()

	// 当前goroutine数量（大概的），不必每次都获取真实的goroutine数量
	// 不需要考虑并发安全
	numGoroutine = 0
)

func ReadDir(dir *Dir) {
	if !dir.File.IsDir {
		return
	}
	// 获取当前目录下所有文件 或 目录
	dirEntrys, err := os.ReadDir(dir.File.Path)
	if err != nil {
		fmt.Println(err)
	}

	childLen := len(dirEntrys)

	if childLen == 0 {
		return
	}

	childs := make([]Dir, 0, childLen)
	// 遍历所有文件 或 目录
	for _, dirEntry := range dirEntrys {
		item := Dir{}
		item.File.Name = dirEntry.Name()
		item.File.IsDir = dirEntry.IsDir()

		// 获取文件大小等信息
		info, err := dirEntry.Info()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		item.File.Size = ByteSize(info.Size())
		item.File.SetUsage(info)

		// 如果是目录的话，记录路径，文件的话不需要记录
		if item.File.IsDir {
			item.File.Path = filepath.Join(dir.File.Path, item.File.Name)
		}
		childs = append(childs, item)
	}

	dir.Clilds = childs

	var wg sync.WaitGroup

	for index, item := range dir.Clilds {
		if item.File.IsDir && !isSkipDir(item) {
			if isTooManyGoroutine() {
				// 如果Goroutine数量太多，则不使用Goroutine
				ReadDir(&dir.Clilds[index])
			} else {
				wg.Add(1)

				// 每次开启一个 goroutine 时，记录 goroutine 的数量，不需要太精准
				// 只增加，goroutine结束时不必减少
				// 如果超出后，自动获取真实的goroutine数量并更新
				numGoroutine++

				go func(i int) {
					defer wg.Done()
					ReadDir(&dir.Clilds[i])
				}(index)
			}
		}
	}

	// 等待子目录统计结束后
	wg.Wait()

	// 计算当前目录的大小
	var dirSize, dirUsage ByteSize
	for index := range dir.Clilds {
		dirSize += dir.Clilds[index].File.Size
		dirUsage += dir.Clilds[index].File.Usage
	}

	dir.File.Size, dir.File.Usage = dirSize, dirUsage

}

func isSkipDir(dir Dir) bool {
	if goos == "linux" {
		if _, ok := linuxSkipDir[dir.File.Path]; ok {
			return true
		}
	}
	return false
}

// 判断是否有太多的goroutine
func isTooManyGoroutine() bool {
	// 读取缓存中的goroutine数量，如果太多
	if numGoroutine > limitOfGoroutine {
		// 则刷新缓存，在进行比较。返回最终结果
		return refreshNumGoroutine() > limitOfGoroutine
	}
	return false
}

// 获取真实的goroutine数量
func refreshNumGoroutine() int {
	numGoroutine = runtime.NumGoroutine()
	return numGoroutine
}
