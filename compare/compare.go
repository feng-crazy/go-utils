package compare

import "reflect"

func boolCompare(a bool, b bool, exp string) bool {
	switch exp {
	case "==":
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
	case "==":
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
	case "==":
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
	case "==":
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
	case "==":
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
	case "==":
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
	case "==":
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
		case "=", "<=", ">=":
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

	if reflect.TypeOf(left).Kind() != reflect.TypeOf(right).Kind() {
		return false
	}
	switch left.(type) {
	case bool:
		return boolCompare(left.(bool), right.(bool), op)
	case byte:
		return byteCompare(left.(byte), right.(byte), op)
	case int:
		return intCompare(left.(int), right.(int), op)
	case int32:
		return int32Compare(left.(int32), right.(int32), op)
	case string:
		return stringCompare(left.(string), right.(string), op)
	case float32:
		return float32Compare(left.(float32), right.(float32), op)
	case float64:
		return float64Compare(left.(float64), right.(float64), op)
	}
	return false
}
