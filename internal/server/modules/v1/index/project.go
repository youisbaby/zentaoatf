package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type ProjectModule struct {
	ProjectCtrl *controller.ProjectCtrl `inject:""`
}

func NewProjectModule() *ProjectModule {
	return &ProjectModule{}
}

// Party 项目
func (m *ProjectModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck())

		index.Post("/", m.ProjectCtrl.Create).Name = "创建项目"
		index.Delete("/", m.ProjectCtrl.Delete).Name = "删除项目"

		index.Get("/getByUser", m.ProjectCtrl.GetByUser).Name = "获取用户参与的项目"
	}
	return module.NewModule("/projects", handler)
}
