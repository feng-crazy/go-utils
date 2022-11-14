package global

type Far struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Line int    `json:"line"`
}

// Content 第一个map的key为文件路径
var Content map[string][]Far

func init() {
	Content = make(map[string][]Far)
}

func WriteContentSrc(filepath string, src []Far) {
	if Content == nil {
		Content = make(map[string][]Far)
	}
	var fars []Far
	fars, ok := Content[filepath]
	if !ok {
		fars = make([]Far, 0)
	}
	//判断是否存在,不存在添加
	for _, f := range src {
		if !judgeExist(f, fars) {
			fars = append(fars, f)
		}
	}

	Content[filepath] = fars
}

func judgeExist(f Far, fars []Far) bool {
	for _, far := range fars {
		if far.Src == f.Src || far.Line == f.Line {
			return true
		}
	}
	return false
}
