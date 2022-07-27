package pkg

import (
	"errors"
	"fmt"
	"strings"
)

var AllowUnit = map[string]ByteSize{
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
func renderSize(size ByteSize, unit string) string {
	if unit == "" {
		return autoSize(size)
	}
	// 以指定的单位打印
	unitSz, err := SizeUnit2Byte(unit)
	if err != nil {
		return fmt.Sprintf("%.2f", size/unitSz)
	}
	return fmt.Sprintf("%.2f%s", size/unitSz, unit)
}

// 通过字符串单位获取单位的大小
func SizeUnit2Byte(unit string) (ByteSize, error) {
	unit = strings.ToUpper(unit)

	value, ok := AllowUnit[unit]
	if !ok {
		err := errors.New("文件大小的单位不正确")
		return 0, err
	}

	return value, nil
}

// 自动以合适的单位打印
func autoSize(size ByteSize) string {
	if size < KB {
		return fmt.Sprintf("%.0f", size)
	} else if size < MB {
		return fmt.Sprintf("%.0fK", size/KB)
	} else if size < GB {
		return fmt.Sprintf("%.0fM", size/MB)
	} else if size < TB {
		return fmt.Sprintf("%.2fG", size/GB)
	} else if size < PB {
		return fmt.Sprintf("%.2fT", size/TB)
	} else if size < EB {
		return fmt.Sprintf("%.2fP", size/PB)
	} else if size < ZB {
		return fmt.Sprintf("%.2fE", size/EB)
	} else if size < YB {
		return fmt.Sprintf("%.2fZ", size/ZB)
	}
	return fmt.Sprintf("%.2fY", size/YB)
}
