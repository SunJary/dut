package pkg

import (
	"os"
	"syscall"
)

type ByteSize float64

const (
	_          = iota             // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota) // 1 << (10*1)
	MB                            // 1 << (10*2)
	GB                            // 1 << (10*3)
	TB                            // 1 << (10*4)
	PB                            // 1 << (10*5)
	EB                            // 1 << (10*6)
	ZB                            // 1 << (10*7)
	YB                            // 1 << (10*8)
)

type File struct {
	Name  string
	Path  string
	Size  ByteSize
	Usage ByteSize
	IsDir bool
}

type Dir struct {
	File   File
	Clilds []Dir
}

const devBSize = 512

func (file *File) SetUsage(f os.FileInfo) {
	switch stat := f.Sys().(type) {
	case *syscall.Stat_t:
		file.Usage = ByteSize(stat.Blocks * devBSize)
	default:
		file.Usage = file.Size
	}
}
