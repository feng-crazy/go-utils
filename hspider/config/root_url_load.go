package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// RootUrlRead Load seeds from path.
func RootUrlRead(path string) ([]string, error) {
	var rootUrls []string
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &rootUrls)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %s", err.Error())
	}
	if len(rootUrls) == 0 {
		return nil, fmt.Errorf("no seed in %s", path)
	}
	return rootUrls, nil
}
