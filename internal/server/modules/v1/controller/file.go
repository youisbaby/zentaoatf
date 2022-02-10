package controller

import (
	commConsts "github.com/aaronchen2k/deeptest/internal/comm/consts"
	fileUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/file"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"github.com/kataras/iris/v12"
)

type FileCtrl struct {
	FileService *service.FileService `inject:""`
	BaseCtrl
}

func NewFileCtrl() *FileCtrl {
	return &FileCtrl{}
}

// Upload 上传文件
func (c *FileCtrl) Upload(ctx iris.Context) {
	f, fh, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(c.ErrResp(commConsts.Failure, err.Error()))
		return
	}
	defer f.Close()

	data, err := c.FileService.UploadFile(ctx, fh)
	if err != nil {
		ctx.JSON(c.ErrResp(commConsts.Failure, err.Error()))
		return
	}
	ctx.JSON(c.SuccessResp(data))
}

// ListDir 列出目录
func (c *FileCtrl) ListDir(ctx iris.Context) {
	parentDir := ctx.URLParam("parentDir")

	if parentDir == "" {
		var err error
		parentDir, err = fileUtils.GetUserHome()
		if err != nil {
			c.ErrResp(commConsts.Failure, err.Error())
			return
		}
	}

	data, err := c.FileService.LoadDirs(parentDir)

	if err != nil {
		ctx.JSON(c.ErrResp(commConsts.Failure, err.Error()))
		return
	}
	ctx.JSON(c.SuccessResp(data))
}
