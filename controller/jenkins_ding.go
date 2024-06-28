package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"ops_tool/common"
	"ops_tool/module"
	"ops_tool/public"
)

// JksDingTalk jenkins构建钉钉通知
// @Summary jenkins构建钉钉通知接口
// @Description 构建环境;构建人,构建的项目名,构建的分支名,构建的URL,构建ID,执行时间,状态
// @Tags jenkins相关接口
// @Accept application/json
// @Produce application/json
// @Param object query module.JenkinsDingTalkRequest false "查询参数"
// @Success 200 {object} ResponseData
// @Router /jenkins/jkDingTalk [post]
func JksDingTalk(c *gin.Context) {
	// 1. 获取args参数并传入到结构体中
	p := new(module.JenkinsDingTalkRequest)
	if err := c.ShouldBindJSON(p); err != nil {
		public.Log.Error("SignUp with invalid param failed, err:", err)

		// 判断错误是否为validator.ValidationErrors类型的
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			public.Log.Error("module.JenkinsDingTalkRequest() ShouldBindJSON failed, err:", err)
			ResponseError(c, http.StatusBadRequest, CodeInvalidParam)
			return
		}
		ResponseWithMsg(c, http.StatusBadRequest, CodeInvalidParam, public.RemoveTopStruct(errs.Translate(public.Trans)))
		return
	}

	// 2. 根据参数获取需要展示的字段
	dingInfo, err := common.GetJenkinsInfo(p)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, CodeServerBusy)
		return
	}

	// 3. 根据字段发送给dingding
	err = common.PostDing(dingInfo)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
