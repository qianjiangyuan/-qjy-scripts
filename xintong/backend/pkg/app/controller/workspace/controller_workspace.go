package workspace

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller"
	workspaceDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/workspace"
	"github.com/xxmyjk/xintong/backend/pkg/app/k8s"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/workspace"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Params struct {
	Name string `json:"name" `
}
type QueryParams struct {
	mParams  Params
	PageNum  int64 `json:"pagenum" `
	PageSize int64 `json:"pagesize" `
}
type PageResult struct {
	Total   int64 `json:"total" `
	Datas   []workspace.WorkSpaceVo `json:"datas" `
}


func QueryPageList(ctx *gin.Context) {
	cors(ctx)
	query := QueryParams{}
	err := ctx.Bind(&query)
	if controller.ErrCatch(err, ctx) {
		return
	}
	queryMap := make(map[string]string)
	queryMap["name"] = query.mParams.Name
	wsDao := workspaceDao.NewWorkSpaceDao()
	total,err := wsDao.GetSelectPageList(queryMap, query.PageNum, query.PageSize, wsDao)
	if controller.ErrCatch(err, ctx) {
		return
	}
	pageResut :=PageResult{}
	pageResut.Total=total
	pageResut.Datas=wsDao.Results;
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": pageResut,
	})
	return
}

func Deletes(ctx *gin.Context) {
	cors(ctx)
	ids := []primitive.ObjectID{}
	err := ctx.Bind(&ids)
	if controller.ErrCatch(err, ctx) {
		return
	}
	wsDao := workspaceDao.NewWorkSpaceDao()
	err = wsDao.Deletes(ids, wsDao)
	if controller.ErrCatch(err, ctx) {
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": "",
	})

}

func Save(ctx *gin.Context) {
	cors(ctx)
	fmt.Printf("begin save\n")
	workspace := workspace.NewWorkspace()
	err := ctx.Bind(&workspace)
	if controller.ErrCatch(err, ctx) {
		return
	}
	wsDao := workspaceDao.NewWorkSpaceDao()
	rs, err := wsDao.Insert(workspace, wsDao)
	if controller.ErrCatch(err, ctx) {
		return
	}
	err = k8s.ExecNameSpace(workspace)
	if controller.ErrCatch(err, ctx) {
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": rs,
	})

	return
}

func Update(ctx *gin.Context) {
	cors(ctx)
	workspace := workspace.NewWorkspace()
	err := ctx.Bind(&workspace)
	if controller.ErrCatch(err, ctx) {
		return
	}
	wsDao := workspaceDao.NewWorkSpaceDao()
	rs, err := wsDao.Update(workspace, workspace.ID, wsDao)
	if controller.ErrCatch(err, ctx) {
		return
	}
	err = k8s.ExecNameSpace(workspace)
	if controller.ErrCatch(err, ctx) {
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": rs,
	})
	return
}
func Detail(ctx *gin.Context) {
	cors(ctx)
	id := ctx.Query("id")
	wsDao := workspaceDao.NewWorkSpaceDao()
	err := wsDao.SelectById(id, wsDao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"cod:e": -1,
			"msg":   err.Error(),
			"data":  err,
		})
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": wsDao.Model,
	})
	return
}
func cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Access-Control-Allow-Origin,x-token,Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
}
