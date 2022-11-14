package saver

import (
	"encoding/json"
	"fmt"
	"icode.baidu.com/baidu/personal-code/far/global"
	"io/ioutil"
	"os"
)

func ReadJsonFile(savePath string) error {
	// savePath=far.out.json
	jsonFile, err := os.Open(savePath)
	if err != nil {
		fmt.Println("ReadJsonFile error opening json file")
		return nil
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading json file")
		return err
	}

	if len(jsonData) != 0 {
		err = json.Unmarshal(jsonData, &global.Content)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteJsonFile(savePath string) error {
	out, err := os.OpenFile(savePath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		return err
	}
	defer out.Close()

	data, err := json.MarshalIndent(global.Content, "", "    ")
	if err != nil {
		return err
	}

	_, err = out.Write(data)
	if err != nil {
		return err
	}

	return nil
}
