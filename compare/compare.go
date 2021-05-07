package compare

import (
	"reflect"

	"github.com/feng-crazy/go-utils/cast"
)

func boolCompare(a bool, b bool, exp string) bool {
	switch exp {
	case "==", "=":
		if a == b {
			return true
		}
	case "!=":
		if a != b {
			return true
		}
	}
	return false
}

func stringCompare(a string, b string, exp string) bool {
	switch exp {
	case ">":
		if a > b {
			return true
		}
	case "<":
		if a < b {
			return true
		}
	case "==", "=":
		if a == b {
			return true
		}
	case "!=":
		if a != b {
			return true
		}
	case ">=":
		if a >= b {
			return true
		}
	case "<=":
		if a <= b {
			return true
		}
	}
	return false
}

func byteCompare(a byte, b byte, exp string) bool {
	switch exp {
	case ">":
		if a > b {
			return true
		}
	case "<":
		if a < b {
			return true
		}
	case "==", "=":
		if a == b {
			return true
		}
	case "!=":
		if a != b {
			return true
		}
	case ">=":
		if a >= b {
			return true
		}
	case "<=":
		if a <= b {
			return true
		}
	}
	return false
}

func intCompare(a int, b int, exp string) bool {
	switch exp {
	case ">":
		if a > b {
			return true
		}
	case "<":
		if a < b {
			return true
		}
	case "==", "=":
		if a == b {
			return true
		}
	case "!=":
		if a != b {
			return true
		}
	case ">=":
		if a >= b {
			return true
		}
	case "<=":
		if a <= b {
			return true
		}
	}
	return false
}

func int32Compare(a int32, b int32, exp string) bool {
	switch exp {
	case ">":
		if a > b {
			return true
		}
	case "<":
		if a < b {
			return true
		}
	case "==", "=":
		if a == b {
			return true
		}
	case "!=":
		if a != b {
			return true
		}
	case ">=":
		if a >= b {
			return true
		}
	case "<=":
		if a <= b {
			return true
		}
	}
	return false
}

func float32Compare(a float32, b float32, exp string) bool {
	switch exp {
	case ">":
		if a > b {
			return true
		}
	case "<":
		if a < b {
			return true
		}
	case "==", "=":
		if a == b {
			return true
		}
	case "!=":
		if a != b {
			return true
		}
	case ">=":
		if a >= b {
			return true
		}
	case "<=":
		if a <= b {
			return true
		}
	}
	return false
}

func float64Compare(a float64, b float64, exp string) bool {
	switch exp {
	case ">":
		if a > b {
			return true
		}
	case "<":
		if a < b {
			return true
		}
	case "==", "=":
		if a == b {
			return true
		}
	case "!=":
		if a != b {
			return true
		}
	case ">=":
		if a >= b {
			return true
		}
	case "<=":
		if a <= b {
			return true
		}
	}
	return false
}

func Compare(left, right interface{}, op string) bool {
	if left == nil || right == nil {
		switch op {
		case "==", "=", "<=", ">=":
			if left == nil && right == nil {
				return true
			} else {
				return false
			}
		case "!=":
			if left == nil && right == nil {
				return false
			} else {
				return true
			}
		case "<", ">":
			return false
		default:
			return false
		}
	}

	switch left.(type) {
	case bool:
		if reflect.TypeOf(right).Kind() != reflect.Bool {
			right = cast.ToBool(right)
		}
		return boolCompare(left.(bool), right.(bool), op)
	case byte:
		if reflect.TypeOf(right).Kind() != reflect.Uint8 {
			right = cast.ToInt8(right)
		}
		return byteCompare(left.(byte), right.(byte), op)
	case int:
		if reflect.TypeOf(right).Kind() != reflect.Int {
			right = cast.ToInt(right)
		}
		return intCompare(left.(int), right.(int), op)
	case int32:
		if reflect.TypeOf(right).Kind() != reflect.Int32 {
			right = cast.ToInt32(right)
		}
		return int32Compare(left.(int32), right.(int32), op)
	case string:
		if reflect.TypeOf(right).Kind() != reflect.String {
			right = cast.ToString(right)
		}
		return stringCompare(left.(string), right.(string), op)
	case float32:
		if reflect.TypeOf(right).Kind() != reflect.Float32 {
			right = cast.ToFloat32(right)
		}
		return float32Compare(left.(float32), right.(float32), op)
	case float64:
		if reflect.TypeOf(right).Kind() != reflect.Float64 {
			right = cast.ToFloat64(right)
		}
		return float64Compare(left.(float64), right.(float64), op)
	}
	return false
}
