package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var dirList []Dir

func RunDut(args []string, config Config) {
	startTime := time.Now().UnixMilli()

	dirList = getTarget(args)

	for index := range dirList {

		ReadDir(&dirList[index])

		Print(&dirList[index], config, 0, "")
	}

	duration := time.Now().UnixMilli() - startTime
	fmt.Printf("耗时：%.3fs\n", float64(duration)/1000)

	if config.Interact {
		Interact()
	}
}

// 根据参数获取目标文件夹
func getTarget(args []string) []Dir {
	len := len(args)
	var dirList = make([]Dir, 0, len)
	if len == 0 {
		// 没有参数，默认当前目录
		currDirName, err := os.Getwd()
		if err != nil {
			currDirName = "."
		}
		dir := Dir{}
		dir.File.Name = currDirName
		dir.File.Path = currDirName
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

			absDirName, err := filepath.Abs(dirName)
			if err != nil {
				absDirName = "."
			}

			dir := Dir{}
			dir.File.Name = absDirName
			dir.File.IsDir = info.IsDir()
			if dir.File.IsDir {
				dir.File.Path = absDirName
			} else {
				dir.File.Size = ByteSize(info.Size())
			}
			dirList = append(dirList, dir)
		}
	}
	return dirList
}
