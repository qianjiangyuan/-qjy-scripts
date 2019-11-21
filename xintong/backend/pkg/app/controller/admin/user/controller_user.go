package user

import (
	"github.com/gin-gonic/gin"
	userDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/admin/user"
	workspaceDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/workspace"
	groupDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/group"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/admin/user"
	"github.com/xxmyjk/xintong/backend/pkg/app/util"
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
	Datas   []user.UserVo `json:"datas" `
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
	uDao := userDao.NewUserDao()
	total,err := uDao.GetSelectPageList(queryMap, query.PageNum, query.PageSize, uDao)

	userVos := []user.UserVo{}
	for _, u := range uDao.Results {
		userVo := copyValueToVo(u)
		wDao := workspaceDao.NewWorkSpaceDao()
		err = wDao.SelectById(u.WorkspaceID.Hex(), wDao)
		userVo.Workspace = wDao.Model
		gDao:= groupDao.NewGroupDao()
		err =gDao.SelectById(u.GroupId.Hex(),gDao)
		userVo.Group=gDao.Model;
		userVos = append(userVos, userVo)
	}
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
	pageResut.Datas=userVos;
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": pageResut,
	})
	return
}

func Save(ctx *gin.Context) {
	cors(ctx)
	m := user.NewUser()
	err := ctx.Bind(&m)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}

	m.Passwd = util.MD5(m.Passwd)
	uDao := userDao.NewUserDao()
	user, err := uDao.FindOne(m.Name)
	if user.Name == m.Name {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  "User already exists",
			"data": err,
		})
		return
	}
	rs, err := uDao.Insert(m, uDao)
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
		"data": rs,
	})
	return
}
func CommitPassWd(ctx *gin.Context){
	cors(ctx)
	m := user.UserVo{}
	err := ctx.Bind(&m)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	uDao := userDao.NewUserDao()
	err = uDao.UpdatePassWd(m)
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
	uDao := userDao.NewUserDao()
	err = uDao.Deletes(ids, uDao)
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

func Update(ctx *gin.Context) {
	cors(ctx)
	m := user.NewUser()
	err := ctx.Bind(&m)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	m.Passwd = util.MD5(m.Passwd)
	uDao := userDao.NewUserDao()
	rs, err := uDao.Update(m, m.ID, uDao)
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
		"data": rs,
	})
	return
}
func Detail(ctx *gin.Context) {
	cors(ctx)
	id := ctx.Query("id")
	uDao := userDao.NewUserDao()
	err := uDao.SelectById(id, uDao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	userVo := copyValueToVo(uDao.Model)
	wDao := workspaceDao.NewWorkSpaceDao()
	err = wDao.SelectById(uDao.Model.WorkspaceID.Hex(), wDao)
	userVo.Workspace = wDao.Model
	gDao:= groupDao.NewGroupDao()
	err =gDao.SelectById(uDao.Model.GroupId.Hex(),gDao)
	userVo.Group=gDao.Model;
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": userVo,
	})
	return
}
func copyValueToVo(u user.User) user.UserVo {
	userVo := user.UserVo{}
	userVo.Name = u.Name
	userVo.ID = u.ID
	userVo.Nickname = u.Nickname
	userVo.Ctime = u.Ctime
	userVo.Email = u.Email
	userVo.Mtime = u.Mtime
	userVo.GroupId=u.GroupId
	userVo.WorkspaceID=u.WorkspaceID
	userVo.Passwd=u.Passwd
	userVo.Phone = u.Phone
	userVo.Administrator=u.Administrator
	return userVo

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
