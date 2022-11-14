package parser

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatchRegexp(t *testing.T) {
	//pattern := `iregistry.baidu-int.com/\S+\/\S+`
	pattern := `iregistry.baidu-int.com/\S+((?!").)*`
	//pattern := `iregistry.baidu-int.com/\S+"?`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpStr := "iregistry.baidu-int.com/chartrepo/acg-det-build/charts/sparkoperator-init-0.0.5-14147767.tgz\"ttt"
	// 执行正则匹配
	matchArr := re.FindStringSubmatch(tmpStr)
	fmt.Println(matchArr)
}
