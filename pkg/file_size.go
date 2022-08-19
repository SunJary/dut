package pkg

import (
	"errors"
	"fmt"
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

func IsAllowUnit(config Config) error {
	if config.ByteSzieUnit == "" {
		return nil
	}
	_, ok := AllowUnit[config.ByteSzieUnit]
	if !ok {
		err := errors.New(`flag -b, --byte Error:
	Byte size unit allow [K|M|G|T|P|E|Z|Y]
	or don't use -b flag to auto choose byte unit`)
		return err
	}
	return nil
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
	unitSz := SizeUnit2Byte(unit)

	return fmt.Sprintf("%.1f%s", fsize/unitSz, unit)
}

// 通过字符串单位获取单位的大小
func SizeUnit2Byte(unit string) float64 {
	value := AllowUnit[unit]

	return value
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
