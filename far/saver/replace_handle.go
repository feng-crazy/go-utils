package saver

import (
	"bufio"
	"fmt"
	"icode.baidu.com/baidu/personal-code/far/global"
	"io"
	"os"
	"strings"
)

func ReplaceHandle() {
	content := global.Content
	for file, fars := range content {
		ReplaceFile(file, fars)
		fmt.Println("FINISH file!", file)
	}
}

func ReplaceFile(file string, fars []global.Far) {
	// 先备份一下文件
	err := CopyFile(file, file+".bak")
	if err != nil {
		fmt.Println("CopyFile file fail:", err)
		return
	}

	in, err := os.Open(file + ".bak")
	if err != nil {
		fmt.Println("open file fail:", err)
		return
	}
	defer in.Close()

	out, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		return
	}
	defer out.Close()

	index := 1
	br := bufio.NewReader(in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			return
		}

		newLine := string(line)
		for _, far := range fars {
			if index == far.Line {
				newLine = ReplaceStr(newLine, far)
				fmt.Println("ReplaceStr index == far.Line", index, far.Line)

			}
		}

		_, err = out.WriteString(newLine)
		if err != nil {
			fmt.Println("write to file fail:", err)
			return
		}

		index++
	}
	fmt.Println("FINISH file! ", file)
}

func ReplaceStr(line string, far global.Far) string {
	if strings.Contains(line, far.Src) {
		line = strings.ReplaceAll(line, far.Src, far.Dest)
	}
	return line
}

func CopyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	//buffer := make([]byte, 1024)
	//
	//n, err := srcFile.Read(buffer)
	//if err == io.EOF {
	//	return err
	//}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	//n, err = destFile.Write(buffer[:n])
	//if err != nil {
	//	return err
	//}

	_, err = io.Copy(destFile, srcFile)               //通过io.Copy实现
	buffer := make([]byte, 1024*1204)                 //1M
	_, err = io.CopyBuffer(destFile, srcFile, buffer) //
	_, err = io.CopyN(destFile, srcFile, 1)           //只包括前n个

	return nil
}
