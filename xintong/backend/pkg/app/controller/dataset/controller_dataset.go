package dataset

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller"
	datasetDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/dataset"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dataset"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Params struct {
	Name string `json:"name" `
	Url  string `json:"url" `
}
type QueryParams struct {
	Mparams  Params `json:"Params"`
	PageNum  int64 `json:"pagenum" `
	PageSize int64 `json:"pagesize" `
}

type PageResult struct {
	Total   int64 `json:"total" `
	Datas   []dataset.DataSetVo `json:"datas" `
}

func QueryPageFiles(ctx *gin.Context) {

	cors(ctx)
	query := QueryParams{}
	err := ctx.Bind(&query)
	if controller.ErrCatch(err, ctx) {
		return
	}
	//files, total, err := upload.GetAllFile(query.Mparams.Url, query.PageNum, query.PageSize)
	//if controller.ErrCatch(err, ctx) {
	//	return
	//}
	datasetVo := dataset.DataSetVo{}
	//datasetVo.Files = files
	//datasetVo.FileTotal = total
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": datasetVo,
	})

}

func QueryPageList(ctx *gin.Context) {
	cors(ctx)
	claims ,exist :=ctx.Get("claims")
	if(exist==false){

	}
	fmt.Println( claims)
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
	queryMap["name"] = query.Mparams.Name
	dsDao := datasetDao.NewDataSetDao()
	total,err := dsDao.GetSelectPageList(queryMap, query.PageNum, query.PageSize, dsDao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	pageResut :=PageResult{}
	pageResut.Total=total
	pageResut.Datas=dsDao.Results;
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
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dsDao := datasetDao.NewDataSetDao()
	err = dsDao.Deletes(ids, dsDao)
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
	dataset := dataset.NewDataSet()
	err := ctx.Bind(&dataset)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dsDao := datasetDao.NewDataSetDao()
	m, err := dsDao.Insert(dataset, dsDao)
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
	dataset := dataset.NewDataSet()
	err := ctx.Bind(&dataset)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dsDao := datasetDao.NewDataSetDao()
	m, err := dsDao.Update(dataset, dataset.ID, dsDao)
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
	dsDao := datasetDao.NewDataSetDao()
	err := dsDao.SelectById(id, dsDao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	//files, err := upload.GetAllFile(dsDao.Model.Url, 1, 100)
	//if err != nil {
	//	ctx.JSON(200, &gin.H{
	//		"code": -1,
	//		"msg":  err.Error(),
	//		"data": err,
	//	})
	//	return
	//}
	datasetVo := dataset.DataSetVo{}
	datasetVo.DataSet = dsDao.Model
	//datasetVo.Files = files.Files
	//datasetVo.FileTotal = files.Total
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": datasetVo,
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
