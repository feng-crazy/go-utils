package str

import (
	"errors"
	"strconv"

	json "github.com/json-iterator/go"
)

func StringTo(s string, t string) (value interface{}, err error) {
	switch t {
	case "int", "int64":
		value, err = strconv.ParseInt(s, 10, 64)
		return
	case "uint", "uint64":
		value, err = strconv.ParseUint(s, 10, 64)
		return
	case "int32":
		value, err = strconv.Atoi(s)
		return
	case "uint32":
		tmp, err := strconv.ParseUint(s, 10, 32)
		value = uint32(tmp)
		return value, err
	case "byte", "int8":
		tmp, err := strconv.ParseUint(s, 10, 8)
		value = byte(tmp)
		return value, err
	case "string", "String":
		return s, nil
	case "double", "float64":
		value, err = strconv.ParseFloat(s, 64)
		return
	case "float", "float32":
		tmp, err := strconv.ParseFloat(s, 32)
		value = float32(tmp)
		return value, err
	case "boolean", "bool":
		value, err = strconv.ParseBool(s)
		return
	default:
		return nil, errors.New("不支持的类型")
	}
}

func StringToMap(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	_ = json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}
