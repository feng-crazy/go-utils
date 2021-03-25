package str

import (
	"errors"
	"strconv"
)

func StringTo(s string, t string) (value interface{}, err error) {
	switch t {
	case "int":
		value, err = strconv.ParseInt(s, 10, 64)
		return
	case "uint":
		value, err = strconv.ParseUint(s, 10, 64)
		return
	case "int32":
		value, err = strconv.Atoi(s)
		return
	case "uint32":
		tmp, err := strconv.ParseUint(s, 10, 32)
		value = uint32(tmp)
		return value, err
	case "byte":
		tmp, err := strconv.ParseUint(s, 10, 8)
		value = byte(tmp)
		return value, err

	case "string":
		return s, nil
	case "double":
		value, err = strconv.ParseFloat(s, 64)
		return
	case "float":
		tmp, err := strconv.ParseFloat(s, 32)
		value = float32(tmp)
		return value, err
	case "boolean":
		value, err = strconv.ParseBool(s)
		return
	default:
		return nil, errors.New("不支持的类型")
	}
}
