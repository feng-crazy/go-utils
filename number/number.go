package number

func ConvertNum(para interface{}) interface{} {
	if IsInt(para) {
		para = ToInt64(para)
	} else if IsFloat(para) {
		para = ToFloat64(para)
	}
	return para
}

func IsInt(para interface{}) bool {
	switch para.(type) {
	case int:
		return true
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	}
	return false
}

func ToInt64(para interface{}) int64 {
	if v, ok := para.(int); ok {
		return int64(v)
	} else if v, ok := para.(int8); ok {
		return int64(v)
	} else if v, ok := para.(int16); ok {
		return int64(v)
	} else if v, ok := para.(int32); ok {
		return int64(v)
	} else if v, ok := para.(int64); ok {
		return v
	}
	return 0
}

func IsFloat(para interface{}) bool {
	switch para.(type) {
	case float32:
		return true
	case float64:
		return true
	}
	return false
}

func ToFloat64(para interface{}) float64 {
	if v, ok := para.(float32); ok {
		return float64(v)
	} else if v, ok := para.(float64); ok {
		return v
	}
	return 0
}
