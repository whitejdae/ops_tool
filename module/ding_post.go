package module

// ProjectInfo Jenkins信息栏
type ProjectInfo struct {
	User                   string //构建人
	ProjectName            string //构建的项目名
	ProjectBranchName      string //构建的分支名
	ProjectBuildConsoleUrl string //构建的URL
	ProjectEnvironment     string //构建环境
	ProjectBuildID         string //构建ID
	ProjectBuildTime       int    //执行时间
	ProjectBuildState      string //状态
}

// Parameters 钉钉接口参数
type Parameters struct {
	Msgtype  string          `json:"msgtype"`
	Markdown MarkDownContent `json:"markdown"`
	At       AtContent       `json:"at"`
}

type MarkDownContent struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type AtContent struct {
	AtMobiles []string `json:"atMobiles"`
}
