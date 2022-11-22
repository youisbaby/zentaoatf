package serverDomain

import "time"

type ZentaoHeartbeatReq struct {
	Secret string `json:"secret"`
	Token  string `json:"token"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
}
type ZentaoHeartbeatResp struct {
	Token           string    `json:"token" yaml:"token"`
	ExpiredTimeUnix int64     `json:"expiredTimeUnix" yaml:"expiredTimeUnix"`
	ExpiredDate     time.Time `json:"expiredDate" yaml:"expiredDate"`
}

type ZentaoResp struct {
	Status string
	Data   string
}
type ZentaoRespData struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type ZentaoResultSubmitReq struct {
	Name        string `json:"name"`
	Seq         string `json:"seq"`
	WorkspaceId int    `json:"workspaceId"`
	ProductId   int    `json:"productId"`
	TaskId      int    `json:"taskId"`

	Task int `json:"task"`
}

type ZentaoLang struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ZentaoSite struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`

	Checked bool `json:"checked"`
}
type ZentaoProduct struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

type ZentaoModule struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ZentaoSuite struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ZentaoTask struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
