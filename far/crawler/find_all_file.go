package crawler

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetAllValueFile(folder string) ([]string, error) {
	allFileIncludeSubFolder, err := GetAllFileIncludeSubFolder(folder)
	if err != nil {
		return nil, err
	}
	result := FilterDeployFile(allFileIncludeSubFolder)
	return result, nil
}

//GetAllFileIncludeSubFolder 递归获取某个目录下的所有文件
func GetAllFileIncludeSubFolder(folder string) ([]string, error) {
	var result []string

	filepath.Walk(folder, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}

		if !fi.IsDir() {
			//如果想要忽略这个目录，请返回filepath.SkipDir，即：
			//return filepath.SkipDir
			result = append(result, path)
		}

		return nil
	})

	return result, nil
}

// FilterDeployFile 过滤出来部署文件
func FilterDeployFile(files []string) []string {
	var result []string

	for _, filename := range files {
		if JudgeFile(filename) {
			result = append(result, filename)
		}
	}

	return result
}

// JudgeFile 判断文件是否包含values.yaml
func JudgeFile(filename string) bool {
	if strings.Contains(filename, "values.yaml") {
		return true
	}
	return false
}
