//go:build linux || openbsd || darwin || netbsd || freebsd
// +build linux openbsd darwin netbsd freebsd

package pkg

import (
	"os"
	"syscall"
)

const devBSize = 512

func (file *File) SetUsage(f os.FileInfo) {
	switch stat := f.Sys().(type) {
	case *syscall.Stat_t:
		file.Usage = ByteSize(stat.Blocks * devBSize)
	default:
		file.Usage = file.Size
	}
}
