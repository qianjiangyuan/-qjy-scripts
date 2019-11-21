package app

import (
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/admin/user"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/auth"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/code_house"
	userGroup "github.com/xxmyjk/xintong/backend/pkg/app/controller/group"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/dataset"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/dlmodel"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/dlservice"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/image"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/train"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/upload"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/workspace"
)

func RegisterRouter(app *gin.Engine) {
	rUpload := app.Group("/upload")
	{
		rUpload.POST("/upload", upload.HandleUploadMutiFile)
		rUpload.POST("/list", upload.List)

	}
	rDataset := app.Group("/dataset")
	{
		rDataset.POST("/queryPageList", dataset.QueryPageList)
		rDataset.POST("/save", dataset.Save)
		rDataset.POST("/update", dataset.Update)
		rDataset.GET("/detail", dataset.Detail)
		rDataset.POST("/deletes", dataset.Deletes)
		rDataset.POST("/queryPageFiles", dataset.QueryPageFiles)
	}
	rImage := app.Group("/image")
	{
		rImage.POST("/save", image.Save)
		rImage.POST("/update", image.Update)
		rImage.GET("/detail", image.Detail)
		rImage.POST("/queryPageList", image.QueryPageList)
		rImage.POST("/deletes", image.Deletes)
	}
	rWorkSpace := app.Group("/workspace")
	{
		rWorkSpace.POST("/queryPageList", workspace.QueryPageList)
		rWorkSpace.POST("/save", workspace.Save)
		rWorkSpace.POST("/update", workspace.Update)
		rWorkSpace.GET("/detail", workspace.Detail)
		rWorkSpace.POST("/deletes", workspace.Deletes)
	}
	rUser := app.Group("/user")
	{
		rUser.POST("/queryPageList", user.QueryPageList)
		rUser.POST("/save", user.Save)
		rUser.POST("/update", user.Update)
		rUser.GET("/detail", user.Detail)
		rUser.POST("/deletes", user.Deletes)
		rUser.POST("/commitPassWd", user.CommitPassWd)
	}
	rTrain := app.Group("/train")
	{
		rTrain.GET("/release", train_task.Release)
		rTrain.POST("/queryPageList", train_task.QueryPageList)
		rTrain.POST("/save", train_task.Save)
		rTrain.POST("/update", train_task.Update)
		rTrain.GET("/detail", train_task.Detail)
		rTrain.POST("/deletes", train_task.Deletes)
		rTrain.GET("/xterm/*w", train_task.Xterm)
		rTrain.POST("/queryPageFiles", train_task.QueryPageFiles)
		rTrain.POST("/queryModelFileUrl", train_task.QueryModelFileUrl)
		rTrain.GET("/stopTrainTask", train_task.StopTrainTask)
	}
	rDlmodel := app.Group("/dlmodel")
	{
		rDlmodel.POST("/queryPageList", dlmodel.QueryPageList)
		rDlmodel.POST("/save", dlmodel.Save)
		rDlmodel.POST("/update", dlmodel.Update)
		rDlmodel.GET("/detail", dlmodel.Detail)
		rDlmodel.POST("/deletes", dlmodel.Deletes)
		rDlmodel.POST("/queryPageFiles", dlmodel.QueryPageFiles)
	}
	rDlservice := app.Group("/dlservice")
	{
		rDlservice.POST("/queryPageList", dlservice.QueryPageList)
		rDlservice.POST("/save", dlservice.Save)
		rDlservice.POST("/update", dlservice.Update)
		rDlservice.GET("/detail", dlservice.Detail)
		rDlservice.POST("/deletes", dlservice.Deletes)
	}
	codeHouse := app.Group("/code-house")
	{
		codeHouse.POST("/queryPageList", code.QueryPageList)
		codeHouse.POST("/save", code.Save)
		codeHouse.POST("/update", code.Update)
		codeHouse.GET("/detail", code.Detail)
		codeHouse.POST("/deletes", code.Deletes)
	}
	rGroup := app.Group("/group")
	{
		rGroup.POST("/queryPageList", userGroup.QueryPageList)
		rGroup.POST("/save", userGroup.Save)
		rGroup.POST("/update", userGroup.Update)
		rGroup.GET("/detail", userGroup.Detail)
		rGroup.POST("/deletes", userGroup.Deletes)
	}
	rAuth := app.Group("/auth")
	{
		rAuth.POST("/login", auth.Login)
		rAuth.POST("/github-login", auth.GitHubLogin)
	}
}
