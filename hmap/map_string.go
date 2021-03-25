package hmap

import (
	"encoding/json"
	"errors"
)

func GetMapValue(mapStr, key string) (error, interface{}) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(mapStr), &m)
	if err != nil {
		return err, nil
	}

	v, ok := m[key]
	if ok {
		return nil, v
	} else {
		return errors.New("not found this key "), nil
	}
}

func MapToStrForJson(m map[string]interface{}) (string, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func MapsToStrForJson(ms []map[string]interface{}) (string, error) {
	data, err := json.Marshal(ms)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetMapValueWithString(mapStr, key string) (error, string) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(mapStr), &m)
	if err != nil {
		return err, ""
	}

	v, ok := m[key].(string)
	if ok {
		return nil, v
	} else {
		return errors.New("not found this key "), ""
	}
}

func GetMapValueWithInt64(mapStr, key string) (error, int64) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(mapStr), &m)
	if err != nil {
		return err, 0
	}

	v, ok := m[key].(int64)
	if ok {
		return nil, v
	} else {
		return errors.New("not found this key "), 0
	}
}

func GetMapValueWithInt(mapStr, key string) (error, int) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(mapStr), &m)
	if err != nil {
		return err, 0
	}

	v, ok := m[key].(int)
	if ok {
		return nil, v
	} else {
		return errors.New("not found this key "), 0
	}
}

func GetMapValueWithFloat64(mapStr, key string) (error, float64) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(mapStr), &m)
	if err != nil {
		return err, 0
	}

	v, ok := m[key].(float64)
	if ok {
		return nil, v
	} else {
		return errors.New("not found this key "), 0
	}
}

func GetMapValueWithFloatBool(mapStr, key string) (error, bool) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(mapStr), &m)
	if err != nil {
		return err, false
	}

	v, ok := m[key].(bool)
	if ok {
		return nil, v
	} else {
		return errors.New("not found this key "), false
	}
}
