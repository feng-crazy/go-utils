package cmd

import (
	"fmt"
	"icode.baidu.com/baidu/personal-code/far/crawler"
	"icode.baidu.com/baidu/personal-code/far/global"
	"icode.baidu.com/baidu/personal-code/far/parser"
	"icode.baidu.com/baidu/personal-code/far/saver"

	"github.com/spf13/cobra"
)

func init() {
	findCmd.Flags().StringVarP(&dirPath, "dir", "d", "", "d or f, replace dir")
	findCmd.Flags().StringVarP(&filePath, "file", "f", "", "d or f, replace file")
	findCmd.Flags().StringVarP(&regexp, "regexp", "e", "", "regular expression")
	findCmd.Flags().StringVarP(&output, "output", "o", "./far.out.json", "json for replaceCmd")

	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "查找目录下文件内容中匹配正则的内容",
	Long:  `查找目录下文件内容中匹配正则的内容，以JSON格式输出`,
	Run: func(cmd *cobra.Command, args []string) {
		// 读取输出文件
		err := saver.ReadJsonFile(output)
		if err != nil {
			fmt.Println("saver.ReadJsonFile err ", err.Error())
			return
		}

		files := make([]string, 0)
		if dirPath != "" {
			//获取所有value.yaml文件
			files, err = crawler.GetAllValueFile(dirPath)
			if err != nil {
				fmt.Println("GetAllValueFile err ", err.Error())
				return
			}
		}

		if filePath != "" {
			// 只是读取每一个文件
			files = append(files, filePath)
		}

		// 解析所有文件,并写入Content
		for _, file := range files {
			keys, err := parser.MatchRegexp(regexp, file)
			if err != nil {
				fmt.Println("parser.MatchRegexp err ", err.Error())
				return
			}

			global.WriteContentSrc(file, keys)
		}

		// 写入json文件
		err = saver.WriteJsonFile(output)
		if err != nil {
			fmt.Println("saver.WriteJsonFile err ", err.Error())
			return
		}
	},
}
