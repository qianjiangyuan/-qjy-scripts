package train_task

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xxmyjk/xintong/backend/pkg/app/connect"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller"
	"github.com/xxmyjk/xintong/backend/pkg/app/controller/upload"
	codeHouseDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/code_house"
	datasetDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/dataset"
	train_taskDao "github.com/xxmyjk/xintong/backend/pkg/app/dao/train"
	"github.com/xxmyjk/xintong/backend/pkg/app/k8s"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/dataset"
	"github.com/xxmyjk/xintong/backend/pkg/app/model/train"
	"github.com/xxmyjk/xintong/backend/pkg/app/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/satori/go.uuid"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"os/exec"
	"time"
)

type Params struct {
	Name string `json:"name" `
	Url  string `json:"url" `
}
type QueryParams struct {
	MParams  Params `json:"Params"`
	PageNum  int64 `json:"pagenum" `
	PageSize int64 `json:"pagesize" `
}


func Xterm(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	// del for double set
	w.Header().Del("Access-Control-Allow-Credentials")
	v := ctx.Request.URL.Query()
	index := v.Get("index")
	taskName := v.Get("task_name")
	fmt.Println(index)
	handler := sockjs.NewHandler("/train/xterm", sockjs.DefaultOptions, func(session sockjs.Session) {
		err := k8s.Attach(taskName, index, session)
		fmt.Println(err)
	})
	handler.ServeHTTP(w, r)
}

func Release(ctx *gin.Context) {
	k8s.Release()

	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": nil,
	})
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
	dao := train_taskDao.NewTrain_taskDao()
	_,err = dao.GetSelectPageList(queryMap, query.PageNum, query.PageSize, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	train_taskVos := []train_task.Train_taskVo{}
	for _, t := range dao.Results {
		taskVo := copyValueToVo(t)
		dsDao := datasetDao.NewDataSetDao()
		err = dsDao.SelectById(taskVo.DatasetId.Hex(), dsDao)
		taskVo.DataSet = dsDao.Model
		codeDao := codeHouseDao.NewCodeHouseDao()
		err = codeDao.SelectById(taskVo.CodeHouseId.Hex(), codeDao)
		taskVo.CodeHouse=codeDao.Model
		train_taskVos = append(train_taskVos, taskVo)

	}

	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": train_taskVos,
	})
	return

	fmt.Println(query.PageSize, query.PageNum)
	/*
		col := connect.Mongo().Collection("train_task")
		pipeline := bson.A{
			bson.D{
				{"$lookup", bson.D{
					{"from", "image"},
					{"localField", "image_id"},
					{"foreignField", "_id"},
					{"as", "image"},
				}},
			},
			bson.D{
				{"$lookup", bson.D{
					{"from", "dataset"},
					{"localField", "dataset_id"},
					{"foreignField", "_id"},
					{"as", "dataset"},
				}},
			},
			bson.D{
				{"$unwind", "$image"},
			},
			bson.D{
				{"$unwind", "$dataset"},
			},
			bson.D{
				{"$addFields", bson.D{
					{"image_name", "$image.name"},
					{"dataset_name", "$dataset.name"},
				}},
			},
			bson.D{
				{"$project", bson.D{
					{"image", 0},
					{"dataset", 0},
				}},
			},
			bson.D{
				{"$skip", (query.PageNum - 1) * query.PageSize},
			},
			bson.D{
				{"$limit", query.PageSize},
			},
		}
		opt := options.Aggregate()
		cur, err := col.Aggregate(context.Background(), pipeline, opt)

		if controller.ErrCatch(err, ctx) {
			return
		}


		if cur == nil {
			ctx.JSON(200, &gin.H{
				"code": 0,
				"msg":  "OK",
				"data": bson.A{},
			})
			return
		}

		rs := bson.A{}
		for cur.Next(context.Background()) {
			one := bson.M{}
			err := cur.Decode(&one)

			fmt.Println("11111")
			fmt.Println(one)
			fmt.Println("11111")
			if controller.ErrCatch(err, ctx) {
				return
			}

			rs = append(rs, one)
		}

		ctx.JSON(200, &gin.H{
			"code": 0,
			"msg":  "OK",
			"data": rs,
		})*/
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
	dao := train_taskDao.NewTrain_taskDao()
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
	trainTask := train_task.Train_task{}
	err := ctx.Bind(&trainTask)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dao := train_taskDao.NewTrain_taskDao()
	m, err := dao.Insert(trainTask, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	dataSetDao := datasetDao.NewDataSetDao()
	err = dataSetDao.SelectById(trainTask.DatasetId.Hex(), dataSetDao)

	//todo

	trainTaskVo := train_task.Train_taskVo{}
	trainTaskVo.Name = trainTask.Name
	trainTaskVo.PodNum = trainTask.PodNum
	trainTaskVo.GpuNum = trainTask.GpuNum
	trainTaskVo.CpuNum = trainTask.CpuNum
	trainTaskVo.Cmd = trainTask.Cmd
	trainTaskVo.MemoryNum = trainTask.MemoryNum
	trainTaskVo.Rdma = trainTask.Rdma
	trainTaskVo.ImageUrl=trainTask.ImageUrl
	trainTaskVo.DataSet = dataSetDao.Model
	err =startK8sJob(trainTaskVo,ctx)
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
func startK8sJob(trainTaskVo train_task.Train_taskVo,ctx *gin.Context) error {

       	fmt.Printf("begin type=%s ",trainTaskVo.Type)
         trainTaskVo.Type = "code" 
	if trainTaskVo.Type == "code"{
        	fmt.Printf("begin code")
		u2 := uuid.NewV4()
		trainTaskVo.CodePath = u2.String()
		util.Clone(trainTaskVo.CodeHouse.Git, trainTaskVo.CodePath)
		imageUrl, err := buildImage(trainTaskVo)
		if controller.ErrCatch(err, ctx) {
			return err
		}
		trainTaskVo.ImageUrl = imageUrl
	} else {
		trainTaskVo.ImageUrl =trainTaskVo.ImageUrl
	}
	err := k8s.ExecJob(trainTaskVo)
	if controller.ErrCatch(err, ctx) {
		return err
	}
	return nil
}
func Update(ctx *gin.Context) {
	cors(ctx)
	trainTask := train_task.NewTrain_task()
	err := ctx.Bind(&trainTask)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
        fmt.Printf("begin train1")
	dao := train_taskDao.NewTrain_taskDao()
	m, err := dao.Update(trainTask, trainTask.ID, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}


	dataSetDao := datasetDao.NewDataSetDao()
	err = dataSetDao.SelectById(trainTask.DatasetId.Hex(), dataSetDao)

	codeDao := codeHouseDao.NewCodeHouseDao()
	err = codeDao.SelectById(trainTask.CodeHouseId.Hex(), codeDao)




	trainTaskVo := train_task.Train_taskVo{}
	trainTaskVo.Name = trainTask.Name
	trainTaskVo.PodNum = trainTask.PodNum
	trainTaskVo.GpuNum = trainTask.GpuNum
	trainTaskVo.CpuNum = trainTask.CpuNum
	trainTaskVo.MemoryNum = trainTask.MemoryNum
	trainTaskVo.Rdma = trainTask.Rdma
	trainTaskVo.Cmd = trainTask.Cmd
	trainTaskVo.DataSet = dataSetDao.Model
	trainTaskVo.CodeHouse=codeDao.Model
//	trainTaskVo.ImageUrl=trainTask.ImageUrl
	trainTaskVo.ImageUrl="test:3.0"
	trainTaskVo.Cmd = "python test.py"
        fmt.Printf("begin train")
	err =startK8sJob(trainTaskVo,ctx)
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
func Detail(ctx *gin.Context) {
	cors(ctx)
	id := ctx.Query("id")
	dao := train_taskDao.NewTrain_taskDao()
	err := dao.SelectById(id, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}
	trainVo := copyValueToVo(dao.Model)
	dataSetDao := datasetDao.NewDataSetDao()
	err = dataSetDao.SelectById(dao.Model.DatasetId.Hex(), dataSetDao)

	codeDao := codeHouseDao.NewCodeHouseDao()
	err = codeDao.SelectById(dao.Model.CodeHouseId.Hex(), codeDao)
	trainVo.CodeHouse=codeDao.Model
	trainVo.DataSet=dataSetDao.Model
	_, err=k8s.DetailJob(trainVo.Name+"-job")
	if err == nil {
	//	trainVo.Status=status
	}
	fmt.Println("trainVo"+trainVo.Status)
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": trainVo,
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
func copyValueToVo(t train_task.Train_task) train_task.Train_taskVo {
	train_taskVo := train_task.Train_taskVo{}
	train_taskVo.Name = t.Name
	train_taskVo.GpuNum = t.GpuNum
	train_taskVo.Ctime = t.Ctime
	train_taskVo.Mtime = t.Mtime
	train_taskVo.ID = t.ID
	train_taskVo.Cmd = t.Cmd
	train_taskVo.DatasetId = t.DatasetId
	train_taskVo.Rdma = t.Rdma
	train_taskVo.MemoryNum = t.MemoryNum
	train_taskVo.CpuNum = t.CpuNum
	train_taskVo.CuserID = t.CuserID
	train_taskVo.PodNum = t.PodNum
	train_taskVo.ImageUrl=t.ImageUrl
	train_taskVo.CodeHouseId=t.CodeHouseId
	return train_taskVo

}


func buildImage(vo train_task.Train_taskVo) (string,error) {
	//cmdStr := "/Users/like/WorkSpace/xintong/backend/deploy/sh/load.sh " + upload.GetPath(image.Url) + "/" + image.FileName + " " + "core.harbor.domain/modelzoo" + " " + image.Name + " " + image.Version
	conf := connect.Conf
	script := conf.GetString("registry.buildimagecript")
	storePath := upload.GetPath(vo.CodePath)
	imageName := "core.harbor.domain/modelzoo/"+vo.Name+ time.Now().Format("20060102150405")

	cmdStr := fmt.Sprintf(
		"%s %s %s ",
		script,
		storePath,
		imageName,
	)
	fmt.Println(cmdStr)
	fmt.Printf("error %s",cmdStr)
	cmd := exec.Command("sh", "-c", cmdStr)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error %s %s ", err.Error(), stderr.String())
		return "",err
	}
	fmt.Printf("out= %s", out.String())
	return imageName,nil

}
func publish(ctx *gin.Context){
	id := ctx.Query("id")
	dao := train_taskDao.NewTrain_taskDao()
	err := dao.SelectById(id, dao)
	if err != nil {
		ctx.JSON(200, &gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": err,
		})
		return
	}

}
func QueryModelFileUrl(ctx *gin.Context){
	cors(ctx)
	query := QueryParams{}
	err := ctx.Bind(&query)
	modelUrl,err :=k8s.GetPV(query.MParams.Name)
	if controller.ErrCatch(err, ctx) {
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
	})
		return
	}
	modelUrl ="../../../"+modelUrl
	ctx.JSON(200, &gin.H{
		"code": 0,
		"msg":  "OK",
		"data": modelUrl,
	})
}
func QueryPageFiles(ctx *gin.Context){
	cors(ctx)
	query := QueryParams{}
	err := ctx.Bind(&query)
	if controller.ErrCatch(err, ctx) {
		return
	}
	files, err := upload.GetAllFile(query.MParams.Url, query.PageNum, query.PageSize)
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

func StopTrainTask(ctx *gin.Context){
	cors(ctx)
	id := ctx.Query("id")
	dao := train_taskDao.NewTrain_taskDao()
	err := dao.SelectById(id, dao)
	if controller.ErrCatch(err, ctx) {
		return
	}

	err = k8s.StopJob(dao.Model.Name+"-job")
	if controller.ErrCatch(err, ctx) {
		return
	}
	Detail(ctx)
}
