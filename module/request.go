package module

// JenkinsDingTalkRequest 请求参数
type JenkinsDingTalkRequest struct {
	Name   string `json:"name" binding:"required"`
	Url    string `json:"url" binding:"required"`
	Id     string `json:"id" binding:"required"`
	Branch string `json:"branch" binding:"required"`
}
