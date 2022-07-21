package pkg

import "fmt"

func renderSize(size ByteSize, unit ByteSize) string {
	if unit == 0 {
		// 自动以合适的单位打印
		if size < KB {
			return fmt.Sprintf("%.0f", size/1)
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
		} else {
			return fmt.Sprintf("%.2fY", size/YB)
		}
	} else {
		// 以指定的单位打印
		return fmt.Sprintf("%.2f%s", size/unit, const2unit(unit))
	}
}

/*
	大小单位还原为字符串
*/
func const2unit(size ByteSize) string {
	switch size {
	case KB:
		return "K"
	case MB:
		return "M"
	case GB:
		return "G"
	case TB:
		return "T"
	case PB:
		return "P"
	case EB:
		return "E"
	case ZB:
		return "Z"
	case YB:
		return "Y"
	default:
		return ""
	}
}
