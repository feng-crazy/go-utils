package slice

import "reflect"

func Contains(sl []interface{}, v interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func Contain(sl interface{}, v interface{}) bool {
	switch reflect.TypeOf(sl).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(sl)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(v, s.Index(i).Interface()) {
					return true
				}
			}
		}
	}
	return false
}

func ContainsInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func ContainsInt64(sl []int64, v int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func ContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func ContainsSlice(smallSlice, bigSlice []string) bool {
	for i := 0; i < len(smallSlice); i++ {
		if !ContainsString(bigSlice, smallSlice[i]) {
			return false
		}
	}

	return true
}
