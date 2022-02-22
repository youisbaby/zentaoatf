package commDomain

type SyncSettings struct {
	ProductId int    `json:"productId"`
	ModuleId  int    `json:"moduleId"`
	SuiteId   int    `json:"suiteId"`
	TaskId    int    `json:"taskId"`
	Lang      string `json:"lang"`

	ByModule        bool `json:"byModule"`
	IndependentFile bool `json:"independentFile"`
}
