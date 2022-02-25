package codef

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func ConvertByte2String(b []byte) string {
	var str string

	// var decodeBytes, _ = charmap.ISO8859_8.NewDecoder().Bytes(b)
	// str = string(decodeBytes)
	// fmt.Println("ISO8859_8", str)
	//
	dec := charmap.ISO8859_16.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(b), dec)
	dst, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	str = string(dst)
	fmt.Println("ISO8859_8 ", str)

	// reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GB18030.NewDecoder())
	// d, e := ioutil.ReadAll(reader)
	// if e != nil {
	// 	logrus.Error("GBK", e)
	// }
	// str = string(d)
	// fmt.Println("GBK ", str)

	// var decodeBytes, _ =simplifiedchinese.GB18030.NewDecoder().Bytes(b)
	// str= string(decodeBytes)
	// fmt.Println("GB18030 ", str)

	// var decodeBytes, _ =simplifiedchinese.GBK.NewDecoder().Bytes(b)
	// str= string(decodeBytes)
	// fmt.Println("GBK ", str)

	return str
}
