package config

import (
	"fmt"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	// 测试正常情况
	result, err := ConfigLoad("../conf/spider.conf")
	if err != nil {
		t.Errorf("../conf/spider.conf is valid but there is an error")
		return
	}

	t.Log(result)
}

func TestConfig_Check(t *testing.T) {
	var testCases = []struct {
		input *Config
		want  error
	}{
		{
			&Config{
				Spider{
					UrlListFile:     "../data/url.data",
					OutputDirectory: "../output",
					MaxDepth:        1,
					CrawlInterval:   1,
					CrawlTimeout:    1,
					TargetUrl:       ".*.(htm|html)$",
					ThreadCount:     8,
				},
			},
			nil,
		},
		{
			&Config{
				Spider{
					UrlListFile:     "",
					OutputDirectory: "../output",
					MaxDepth:        1,
					CrawlInterval:   1,
					CrawlTimeout:    1,
					TargetUrl:       ".*.(htm|html)$",
					ThreadCount:     8,
				},
			},
			fmt.Errorf("there must be an error"),
		},
		{
			&Config{
				Spider{
					UrlListFile:     "../data/url.data",
					OutputDirectory: "",
					MaxDepth:        1,
					CrawlInterval:   1,
					CrawlTimeout:    1,
					TargetUrl:       ".*.(htm|html)$",
					ThreadCount:     8,
				},
			},
			fmt.Errorf("there must be an error"),
		},
		{
			&Config{
				Spider{
					UrlListFile:     "../data/url.data",
					OutputDirectory: "../output",
					MaxDepth:        0,
					CrawlInterval:   1,
					CrawlTimeout:    1,
					TargetUrl:       ".*.(htm|html)$",
					ThreadCount:     8,
				},
			},
			fmt.Errorf("there must be an error"),
		},
		{
			&Config{
				Spider{
					UrlListFile:     "../data/url.data",
					OutputDirectory: "../output",
					MaxDepth:        1,
					CrawlInterval:   -1,
					CrawlTimeout:    1,
					TargetUrl:       ".*.(htm|html)$",
					ThreadCount:     8,
				},
			},
			fmt.Errorf("there must be an error"),
		},
		{
			&Config{
				Spider{
					UrlListFile:     "../data/url.data",
					OutputDirectory: "../output",
					MaxDepth:        1,
					CrawlInterval:   1,
					CrawlTimeout:    0,
					TargetUrl:       ".*.(htm|html)$",
					ThreadCount:     8,
				},
			},
			fmt.Errorf("there must be an error"),
		},
		{
			&Config{
				Spider{
					UrlListFile:     "../data/url.data",
					OutputDirectory: "../output",
					MaxDepth:        1,
					CrawlInterval:   1,
					CrawlTimeout:    1,
					TargetUrl:       ".*.(((htm|html)$",
					ThreadCount:     8,
				},
			},
			fmt.Errorf("there must be an error"),
		},
		{
			&Config{
				Spider{
					UrlListFile:     "../data/url.data",
					OutputDirectory: "../output",
					MaxDepth:        1,
					CrawlInterval:   1,
					CrawlTimeout:    1,
					TargetUrl:       ".*.(((htm|html)$",
					ThreadCount:     0,
				},
			},
			fmt.Errorf("there must be an error"),
		},
	}

	for index, testCase := range testCases {
		err := testCase.input.Check()
		if testCase.want != nil {
			if err == nil {
				t.Errorf("testCases[%d] should cause an error but not", index)
				return
			}
		} else {
			if err != nil {
				t.Errorf("testCases[%d] should not cause an error but there is an error", index)
				return
			}
		}
	}
}
