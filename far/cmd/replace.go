package cmd

import (
	"fmt"
	"icode.baidu.com/baidu/personal-code/far/saver"

	"github.com/spf13/cobra"
)

func init() {
	replaceCmd.Flags().StringVarP(&input, "input", "i", "./far.out.json", "json from findCmd")
	rootCmd.AddCommand(replaceCmd)
}

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "替换目录下文件内容中匹配JSON key的内容为JSON value",
	Long:  `查找目录下文件内容中匹配正则的内容，以JSON格式输出`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("替换目录下文件内容中匹配JSON key的内容为JSON value")

		// 读取文件内容
		err := saver.ReadJsonFile(input)
		if err != nil {
			fmt.Println("ReadJsonFile err ", err.Error())
			return
		}
		saver.ReplaceHandle()
	},
}
