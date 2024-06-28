package common

import (
	"context"
	"ops_tool/module"
	"ops_tool/public"
	"strconv"
	"strings"
	"time"
)

func GetJenkinsInfo(request *module.JenkinsDingTalkRequest) (projectInfo *module.ProjectInfo, err error) {
	// 初始化jenkins连接
	ctx := context.Background()
	Init(strings.Split(request.Url, "/job/")[0], ctx)

	// 获取需要的jenkins构建信息
	job, err := jk.GetJob(ctx, request.Name)
	if err != nil {
		public.Log.Errorf("Get %s info failed,err:%s", request.Name, err)
		return
	}

	atoi, _ := strconv.Atoi(request.Id)
	build, err := job.GetBuild(ctx, int64(atoi))
	if err != nil {
		public.Log.Errorf("Get %s 的 %v info failed,err:%v", request.Name, request.Id, err)
		return
	}

	var userName string
	for _, pd := range build.Raw.Actions {
		for _, cause := range pd.Causes {
			for key, value := range cause {
				if key == "userId" {
					userName, _ = value.(string) // 类型断言将 interface{} 转换为 string
				}
			}
		}
	}

	// 构建钉钉通知信息字段
	projectInfo = &module.ProjectInfo{
		User:                   userName,
		ProjectName:            request.Name,
		ProjectBranchName:      request.Branch,
		ProjectBuildConsoleUrl: request.Url + "/" + request.Id + "/console",
		ProjectEnvironment:     ifProjectEnvironment(request.Name),
		ProjectBuildID:         request.Id,
		ProjectBuildTime:       buildSubTime(build.Raw.Timestamp),
		ProjectBuildState:      build.Raw.Result,
	}
	return projectInfo, err
}

// ifProjectEnvironment 判断项目环境
func ifProjectEnvironment(projectName string) (projectEnvironment string) {
	index := strings.Split(projectName, "_")
	switch {
	case strings.Contains(index[0], "dev"):
		projectEnvironment = "Dev"
	case strings.Contains(index[0], "test"):
		projectEnvironment = "Test"
	case strings.Contains(index[0], "uat"):
		projectEnvironment = "Uat"
	case strings.Contains(index[0], "pro") || strings.Contains(index[0], "ops"):
		projectEnvironment = "Pro"
	default:
		projectEnvironment = "test"
	}
	return projectEnvironment
}

// buildSubTime 项目总共执行时间
func buildSubTime(TimeStamp int64) int {
	// 将时间戳转换为时间对象
	timeObject := time.Unix(0, TimeStamp*int64(time.Millisecond))
	// 当前时间
	now := time.Now()

	// 计算时间差
	duration := now.Sub(timeObject)

	// 将时间差转换为分钟
	minutes := int(duration.Minutes())
	if minutes <= 0 {
		return 1
	}
	return minutes
}
