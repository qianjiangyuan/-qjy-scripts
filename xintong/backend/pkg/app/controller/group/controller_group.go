package group

import (
	"fmt"
	"github.com/gin-gonic/gin"
	groupDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/group"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/group"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Params struct {
	Name string `json:"name" `
}
type QueryParams struct {
	mParams Params
	PageNum int64 `json:"pagenum" `
	PageSize int64 `json:"pagesize" `

}

type PageResult struct {
	Total   int64 `json:"total" `
	Datas   []group.Group `json:"datas" `
}
func QueryPageList(ctx *gin.Context)  {
	cors(ctx)
 	query :=QueryParams{}
	err := ctx.Bind(&query)
	if(err != nil){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
			"data": err,
		})
		return
	}
	queryMap:=make(map[string]string)
	queryMap["name"]=query.mParams.Name
	dao :=groupDao.NewGroupDao()
	total,err :=dao.GetSelectPageList(queryMap,query.PageNum,query.PageSize,dao)
	if(err != nil){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
			"data": err,
		})
		return
	}
	pageResut :=PageResult{}
	pageResut.Total=total
	pageResut.Datas=dao.Results;
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": pageResut,
	})
	return
}
func Deletes(ctx *gin.Context) {
	cors(ctx)
	ids :=[]primitive.ObjectID{}
	err := ctx.Bind(&ids)
	if(err != nil){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
			"data": err,
		})
		return
	}
	dao :=groupDao.NewGroupDao()
	err =dao.Deletes(ids,dao)
	if(err != nil){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
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
	group:= group.NewGroup()
	err := ctx.Bind(&group)
	if(err != nil ){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
			"data": err,
		})
		return
	}
	dao :=groupDao.NewGroupDao()
	m ,err:=dao.Insert(group,dao)
	if(err != nil){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
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

func Update(ctx *gin.Context)  {
	cors(ctx)
	group:= group.NewGroup()
	err := ctx.Bind(&group)
	if(err != nil ){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
			"data": err,
		})
		return
	}
	dao :=groupDao.NewGroupDao()
	m ,err:=dao.Update(group,group.ID,dao)
	if(err != nil){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
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
	id :=ctx.Query("id")
	dao :=groupDao.NewGroupDao()
	err:=dao.SelectById(id,dao)
	if(err != nil){
		ctx.JSON(200, &gin.H{
			"code":    -1,
			"msg": err.Error(),
			"data": err,
		})
		return
	}
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": dao.Model,
	})
	return
}
func cors(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Access-Control-Allow-Origin,x-token,Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
}