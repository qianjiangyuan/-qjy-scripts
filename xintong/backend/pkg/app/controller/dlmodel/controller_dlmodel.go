package dlmodel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/upload"
	dlmodelDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/dlmodel"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dataset"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dlmodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Params struct {
	Name string `json:"name" `
	Url  string `json:"url" `
}
type QueryParams struct {
	MParams  Params
	PageNum  int64 `json:"pagenum" `
	PageSize int64 `json:"pagesize" `
}

func QueryPageList(ctx *gin.Context) {
	cors(ctx)
	query := QueryParams{}
	err := ctx.Bind(&query)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	queryMap := make(map[string]string)
	queryMap["name"] = query.MParams.Name
	dao := dlmodelDao.NewDlmodelDao()
	_,err = dao.GetSelectPageList(queryMap, query.PageNum, query.PageSize, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": dao.Results,
	})
	return
}
func Deletes(ctx *gin.Context) {
	cors(ctx)
	ids := []primitive.ObjectID{}
	err := ctx.Bind(&ids)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dao := dlmodelDao.NewDlmodelDao()
	err = dao.Deletes(ids, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
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
	dlmodel := dlmodel.NewDlmodel()
	err := ctx.Bind(&dlmodel)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dao := dlmodelDao.NewDlmodelDao()
	m, err := dao.Insert(dlmodel, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": m,
	})
	return
}

func Update(ctx *gin.Context) {
	cors(ctx)
	dlmodel := dlmodel.NewDlmodel()
	err := ctx.Bind(&dlmodel)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dao := dlmodelDao.NewDlmodelDao()
	m, err := dao.Update(dlmodel, dlmodel.ID, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": m,
	})
	return
}
func Detail(ctx *gin.Context) {
	cors(ctx)
	id := ctx.Query("id")
	dao := dlmodelDao.NewDlmodelDao()
	err := dao.SelectById(id, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	files,  err := upload.GetAllFile(dao.Model.Url, 1, 100)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dlmodelVo := dlmodel.DlmodelVo{}
	dlmodelVo.Dlmodel = dao.Model
	dlmodelVo.Files = files.Files
	dlmodelVo.FileTotal = files.Total
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": dlmodelVo,
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

func QueryPageFiles(ctx *gin.Context) {

	cors(ctx)
	query := QueryParams{}
	err := ctx.Bind(&query)
	if controller.ErrCatch(err, ctx) {
		return
	}
	files,  err := upload.GetAllFile(query.MParams.Url, query.PageNum, query.PageSize)
	if controller.ErrCatch(err, ctx) {
		return
	}
	datasetVo := dataset.DataSetVo{}
	datasetVo.Files = files.Files
	datasetVo.FileTotal = files.Total
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": datasetVo,
	})

}