package str

import "strings"

func TrimStringSlice(raw []string) []string {
	if raw == nil {
		return []string{}
	}

	cnt := len(raw)
	arr := make([]string, 0, cnt)
	for i := 0; i < cnt; i++ {
		item := strings.TrimSpace(raw[i])
		if item == "" {
			continue
		}

		arr = append(arr, item)
	}

	return arr
}

func ExtractCN(s string) string {
	r := []rune(s)

	cnstr := ""
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])
		}

	}
	return cnstr
}
