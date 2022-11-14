package parser

import (
	"bufio"
	"fmt"
	"icode.baidu.com/baidu/personal-code/far/global"
	"io"
	"os"
	"regexp"
	"strings"
)

func MatchRegexp(pattern string, filepath string) ([]global.Far, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 读取文件每一行
	fileHanle, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer fileHanle.Close()

	reader := bufio.NewReader(fileHanle)

	var results []global.Far
	index := 1
	// 按行处理
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Printf("read line:%s\n", line)

		// 执行正则匹配
		matchArr := re.FindStringSubmatch(string(line))
		if len(matchArr) >= 1 {
			//results = append(results, matchArr...)
			for _, m := range matchArr {
				results = append(results, global.Far{
					Src:  m,
					Dest: "hdf_tmp",
					Line: index,
				})
			}
		} else {
			if strings.Contains(string(line), "iregistry.baidu-int.com") {
				fmt.Println("匹配不正确 ", string(line), matchArr)
			}
		}
		index++
	}

	return results, nil
}
