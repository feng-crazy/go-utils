package str

import (
	"errors"
	"sort"

	"github.com/feng-crazy/go-utils/hmap"
)

func KeysOfMap(m map[string]interface{}) []string {
	keys := make(sort.StringSlice, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	keys.Sort()
	return []string(keys)
}

func SingleMapKeyToStrArray(ms map[string]interface{}) []string {
	keys := make([]string, len(ms))
	for key := range ms {
		keys = append(keys, key)
	}

	return keys
}

func SingleStructJsonTagToStrArray(st interface{}) ([]string, error) {
	if st == nil {
		return nil, errors.New("is nil")
	}

	ms, err := hmap.Struct2MapWithJson(st)
	if err != nil {
		return nil, errors.New("is nil")
	}

	keys := make([]string, len(ms))
	for key := range ms {
		keys = append(keys, key)
	}

	return keys, nil
}
