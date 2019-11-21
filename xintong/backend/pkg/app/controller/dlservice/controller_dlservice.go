package dlservice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller"
	dlmodelDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/dlmodel"
	dlserviceDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/dlservice"
	imageDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/image"
	"github.com/xxmyjk/xintong/backend/pkg/app/k8s"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dlservice"
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
	queryMap["name"] = query.mParams.Name
	dao := dlserviceDao.NewDlserviceDao()
	_,err = dao.GetSelectPageList(queryMap, query.PageNum, query.PageSize, dao)
	dlserviceVos := []dlservice.DlserviceVo{}
	for _, service := range dao.Results {
		dlserviceVo := copyValueToVo(service)
		igDao := imageDao.NewImageDao()
		err = igDao.SelectById(service.ImageID.Hex(), igDao)
		dlserviceVo.Image = igDao.Model

		dmDao := dlmodelDao.NewDlmodelDao()
		err = dmDao.SelectById(service.DlmodelID.Hex(), dmDao)
		dlserviceVo.Dlmodel = dmDao.Model
		dlserviceVos = append(dlserviceVos, dlserviceVo)
	}
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
		"data": dlserviceVos,
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
	dao := dlserviceDao.NewDlserviceDao()
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
	dlservice := dlservice.NewDlservice()
	err := ctx.Bind(&dlservice)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dao := dlserviceDao.NewDlserviceDao()
	m, err := dao.Insert(dlservice, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dlserviceVo := copyValueToVo(dlservice)
	dmDao := dlmodelDao.NewDlmodelDao()
	err = dmDao.SelectById(dlserviceVo.DlmodelID.Hex(), dmDao)
	dlserviceVo.Dlmodel = dmDao.Model
	err = k8s.ExecService(dlserviceVo)
	if controller.ErrCatch(err, ctx) {
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
	dlservice := dlservice.NewDlservice()
	err := ctx.Bind(&dlservice)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dao := dlserviceDao.NewDlserviceDao()
	m, err := dao.Update(dlservice, dlservice.ID, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}

	dlserviceVo := copyValueToVo(dlservice)
	dmDao := dlmodelDao.NewDlmodelDao()
	err = dmDao.SelectById(dlserviceVo.DlmodelID.Hex(), dmDao)
	dlserviceVo.Dlmodel = dmDao.Model

	err = k8s.ExecService(dlserviceVo)

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
	dao := dlserviceDao.NewDlserviceDao()
	err := dao.SelectById(id, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dlserviceVo := copyValueToVo(dao.Model)
	igDao := imageDao.NewImageDao()
	err = igDao.SelectById(dlserviceVo.ImageID.Hex(), igDao)
	dlserviceVo.Image = igDao.Model

	dmDao := dlmodelDao.NewDlmodelDao()
	err = dmDao.SelectById(dlserviceVo.DlmodelID.Hex(), dmDao)
	dlserviceVo.Dlmodel = dmDao.Model
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": dlserviceVo,
	})
	return
}
func copyValueToVo(service dlservice.Dlservice) dlservice.DlserviceVo {
	serviceVo := dlservice.DlserviceVo{}
	serviceVo.Name = service.Name
	serviceVo.ID = service.ID
	serviceVo.Ctime = service.Ctime
	serviceVo.Mtime = service.Mtime
	serviceVo.Version = service.Version
	serviceVo.ImageID = service.ImageID
	serviceVo.DlmodelID = service.DlmodelID
	serviceVo.Label = service.Label
	serviceVo.Platform = service.Platform
	serviceVo.Description = service.Description
	serviceVo.Path = service.Path
	return serviceVo
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
