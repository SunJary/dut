package pkg

import (
	"errors"
	"fmt"
	"strings"
)

var AllowUnit = map[string]float64{
	"K": KB,
	"M": MB,
	"G": GB,
	"T": TB,
	"P": PB,
	"E": EB,
	"Z": ZB,
	"Y": YB,
}

// 根据规则显示文件的大小
func renderSize(file File, config Config) string {
	var fsize float64
	unit := config.ByteSzieUnit
	usage := config.Usage

	if usage {
		fsize = float64(file.Usage)
	} else {
		fsize = float64(file.Size)
	}

	if unit == "" {
		return autoSize(fsize)
	}
	// 以指定的单位打印
	unitSz, err := SizeUnit2Byte(unit)
	if err != nil {
		return fmt.Sprintf("%.1f", fsize/unitSz)
	}
	return fmt.Sprintf("%.1f%s", fsize/unitSz, unit)
}

// 通过字符串单位获取单位的大小
func SizeUnit2Byte(unit string) (float64, error) {
	unit = strings.ToUpper(unit)

	value, ok := AllowUnit[unit]
	if !ok {
		err := errors.New("文件大小的单位不正确")
		return 0, err
	}

	return value, nil
}

// 自动以合适的单位打印
func autoSize(fsize float64) string {
	if fsize < KB {
		return fmt.Sprintf("%.0f ", fsize)
	} else if fsize < MB {
		return fmt.Sprintf("%.1fK", fsize/KB)
	} else if fsize < GB {
		return fmt.Sprintf("%.1fM", fsize/MB)
	} else if fsize < TB {
		return fmt.Sprintf("%.1fG", fsize/GB)
	} else if fsize < PB {
		return fmt.Sprintf("%.1fT", fsize/TB)
	} else if fsize < EB {
		return fmt.Sprintf("%.1fP", fsize/PB)
	} else if fsize < ZB {
		return fmt.Sprintf("%.1fE", fsize/EB)
	} else if fsize < YB {
		return fmt.Sprintf("%.1fZ", fsize/ZB)
	}
	return fmt.Sprintf("%.1fY", fsize/YB)
}
