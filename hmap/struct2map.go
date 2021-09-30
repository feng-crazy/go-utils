package hmap

import (
	"reflect"
	"strings"

	json "github.com/json-iterator/go"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Struct2MapWithJson(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Struct2MapWithGorm(obj interface{}) (map[string]interface{}, error) {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		key := t.Field(i).Name
		tagValue := f.Tag.Get("gorm")
		tags := strings.Split(tagValue, ";")
		for _, tag := range tags {
			dd := strings.Split(tag, ":")
			if dd[0] == "column" {
				key = dd[1]
				break
			}
		}

		data[key] = v.Field(i).Interface()
	}

	return data, nil
}

// obj 必须要结构体指针
func Map2Struct(m map[string]interface{}, obj interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}

func MapArray2Struct(m interface{}, obj interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}
