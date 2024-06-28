package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ops_tool/conf"
	"ops_tool/dao/mysql"
	"ops_tool/module"
	"ops_tool/public"
)

func PostDing(info *module.ProjectInfo) (err error) {
	dataTitle := info.ProjectEnvironment + "服务发布"
	phone := mysql.GetUserPhone(info.User)
	dataContent := fmt.Sprintf(
		"# %v\n - **构建者**：%v\n - **构建项目**：%v\n - **构建分支**：%v\n - **构建URL**：[点击查看](%v)\n - **构建ID**：%v\n - **构建时间**：%v分\n - **构建状态**：%v\n - @%v",
		dataTitle,
		info.User,
		info.ProjectName,
		info.ProjectBranchName,
		info.ProjectBuildConsoleUrl,
		info.ProjectBuildID,
		info.ProjectBuildTime,
		info.ProjectBuildState,
		phone,
	)
	public.Log.Info("发送钉钉webhook：\n", dataContent)

	// 将参数转换为JSON格式
	var tmp = make([]string, 0, 10)
	tmp = append(tmp, phone)
	params := module.Parameters{
		Msgtype: "markdown",
		Markdown: module.MarkDownContent{
			Title: dataTitle,
			Text:  dataContent,
		},
		At: module.AtContent{
			AtMobiles: tmp,
		},
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		public.Log.Error("Marshal Json failed;err: ", err)
		return
	}

	// 上传参数
	response, err := http.Post(conf.Conf.JenkinsDingTalk.DingDing.Token, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		public.Log.Error("post dingingTalk failed;err: ", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		public.Log.Error("reader response body flied; err: ", err)
	}

	public.Log.Info("发送成功，钉钉webhook返回响应:", string(responseBody))
	return
}
