//go:build windows || plan9
// +build windows plan9

package pkg

import (
	"os"
)

const devBSize = 512

func (file *File) SetUsage(f os.FileInfo) {
	file.Usage = file.Size
}
