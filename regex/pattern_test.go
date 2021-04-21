package regex

import (
	"fmt"
	"testing"
)

func TestPatternRegex(t *testing.T) {
	fmt.Println(NumericRegexp.Match("11"))
	fmt.Println(AlphaNumericRegexp.Match("2113daf5SDD"))
	fmt.Println(AlphaRegexp.Match("DSSD"))
	fmt.Println(AlphaCapsOnlyRegexp.Match("SDAF"))
	fmt.Println(AlphaNumericCapsOnlyRegexp.Match("11DADA"))
	fmt.Println(UrlRegexp.Match("http://www.baidu.com	"))
	fmt.Println(EmailRegexp.Match("894220128@qq.com"))
	fmt.Println(HashtagHexRegexp.Match("afafda"))
	fmt.Println(ZeroXHexRegexp.Match("131daf"))
	fmt.Println(IPv4Regexp.Match("192.168.0.1"))
	fmt.Println(IPv6Regexp.Match(""))
}
